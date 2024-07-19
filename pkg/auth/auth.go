// Package auth provides simple client for WME authentication.
package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// LoginRequest parameters required for login request.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse payload for the login response.
type LoginResponse struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// RevokeTokenRequest parameters required for revoke token request.
type RevokeTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshTokenRequest parameters required for refresh token request.
type RefreshTokenRequest struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshTokenResponse payload for the refresh token response.
type RefreshTokenResponse struct {
	IDToken     string `json:"id_token"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Loginer interface for login endpoint.
type Loginer interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

// TokenRefresher interface for token refresh endpoint.
type TokenRefresher interface {
	RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error)
}

// TokenRevoker interface for token revoke endpoint.
type TokenRevoker interface {
	RevokeToken(ctx context.Context, req *RevokeTokenRequest) error
}

// API interface for WME authentication.
type API interface {
	Loginer
	TokenRefresher
	TokenRevoker
}

var tokenStoreFile = "tokenstore.json"

// NewClient create new http client for WME authentication.
func NewClient() API {
	return &Client{
		BaseURL:    "https://auth.enterprise.wikimedia.com/v1",
		HTTPClient: &http.Client{},
	}
}

// Client http client to simplify work with WME authentication.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func (c *Client) post(ctx context.Context, url string, v interface{}) (*http.Response, error) {
	body := bytes.NewBuffer([]byte{})

	if err := json.NewEncoder(body).Encode(v); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", c.BaseURL, url), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		data, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
		return nil, fmt.Errorf("%s: %s", res.Status, string(data))
	}

	return res, nil
}

// Login triggers login endpoint and returns fresh set of tokens.
func (c *Client) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	res, err := c.post(ctx, "/login", req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	rsp := new(LoginResponse)

	if err := json.NewDecoder(res.Body).Decode(rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// RefreshToken gets new set of tokens using refresh token.
func (c *Client) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	res, err := c.post(ctx, "/token-refresh", req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	rsp := new(RefreshTokenResponse)

	if err := json.NewDecoder(res.Body).Decode(rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// RevokeToken invalidates refresh token and all related access tokens.
func (c *Client) RevokeToken(ctx context.Context, req *RevokeTokenRequest) error {
	_, err := c.post(ctx, "/token-revoke", req)
	return err
}

type Tokenstore struct {
	AccessToken             string    `json:"access_token"`
	AccessTokenGeneratedAt  time.Time `json:"access_token_generated_at"`
	RefreshToken            string    `json:"refresh_token"`
	RefreshTokenGeneratedAt time.Time `json:"refresh_token_generated_at"`
}

// Helper struct holds the username, password and token data
type Helper struct {
	Username  string
	Password  string
	TokenData *Tokenstore
	mutex     sync.Mutex
	API       API
}

// NewHelper creates a new Helper instance
func NewHelper(api API) (*Helper, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	username := os.Getenv("WME_USERNAME")
	password := os.Getenv("WME_PASSWORD")
	if username == "" || password == "" {
		return nil, errors.New("username or password not set in .env file")
	}

	return &Helper{
		Username: username,
		Password: password,
		API:      api,
	}, nil
}

// GetAccessToken manages the token state and returns a valid access token
func (h *Helper) GetAccessToken() (string, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Check if the tokenstore file exists
	if _, err := os.Stat(tokenStoreFile); os.IsNotExist(err) {
		// File does not exist, create and login
		return h.loginAndStoreTokens(tokenStoreFile)
	}

	// File exists, read and unmarshal token data
	data, err := os.ReadFile(tokenStoreFile)
	if err != nil {
		return "", err
	}

	var tokenStore Tokenstore
	if err := json.Unmarshal(data, &tokenStore); err != nil {
		return "", err
	}

	// Check if the access token is still valid
	if time.Since(tokenStore.AccessTokenGeneratedAt) < 24*time.Hour {
		return tokenStore.AccessToken, nil
	}

	// Check if the refresh token is still valid
	if time.Since(tokenStore.RefreshTokenGeneratedAt) < 30*24*time.Hour {
		return h.refreshAndStoreTokens(tokenStoreFile, &tokenStore)
	}

	// Both tokens are expired, login again
	return h.loginAndStoreTokens(tokenStoreFile)
}

// loginAndStoreTokens logs in and stores the tokens
func (h *Helper) loginAndStoreTokens(tokenStoreFile string) (string, error) {
	loginRequest := &LoginRequest{
		Username: h.Username,
		Password: h.Password,
	}

	loginResponse, err := h.API.Login(context.Background(), loginRequest)
	if err != nil {
		return "", err
	}

	tokenStore := &Tokenstore{
		AccessToken:             loginResponse.AccessToken,
		AccessTokenGeneratedAt:  time.Now(),
		RefreshToken:            loginResponse.RefreshToken,
		RefreshTokenGeneratedAt: time.Now(),
	}

	if err := h.storeTokens(tokenStoreFile, tokenStore); err != nil {
		return "", err
	}

	return tokenStore.AccessToken, nil
}

// refreshAndStoreTokens refreshes the access token and stores the new tokens
func (h *Helper) refreshAndStoreTokens(tokenStoreFile string, tokenStore *Tokenstore) (string, error) {
	refreshRequest := &RefreshTokenRequest{
		Username:     h.Username,
		RefreshToken: tokenStore.RefreshToken,
	}

	var refreshResponse *RefreshTokenResponse
	var err error

	// Check if the refresh token is less than 30 days old
	if time.Since(tokenStore.RefreshTokenGeneratedAt) < 30*24*time.Hour {
		refreshResponse, err = h.API.RefreshToken(context.Background(), refreshRequest)
	}

	if err != nil {
		return "", err
	}

	tokenStore.AccessToken = refreshResponse.AccessToken
	tokenStore.AccessTokenGeneratedAt = time.Now()

	if err := h.storeTokens(tokenStoreFile, tokenStore); err != nil {
		return "", err
	}

	return tokenStore.AccessToken, nil
}

// storeTokens writes the token data to the tokenstore file in json
func (h *Helper) storeTokens(tokenStoreFile string, tokenStore *Tokenstore) error {
	data, err := json.Marshal(tokenStore)
	if err != nil {
		return err
	}

	//nolint:errcheck
	return os.WriteFile(tokenStoreFile, data, 0600)
}

// ClearState clears the token state by revoking the refresh token and deleting the tokenstore file
func (h *Helper) ClearState() error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Check if the file exists
	if _, err := os.Stat(tokenStoreFile); err != nil {
		if os.IsNotExist(err) {
			// File does not exist, nothing to clear
			return nil
		}
		// Some other error occurred when checking the file
		return err
	}

	// Read the file
	data, err := os.ReadFile(tokenStoreFile)
	if err != nil {
		return err
	}

	// Unmarshal the token store
	var tokenStore Tokenstore
	if err := json.Unmarshal(data, &tokenStore); err != nil {
		return err
	}

	// Revoke the refresh token if it exists
	if tokenStore.RefreshToken != "" {
		revokeRequest := &RevokeTokenRequest{
			RefreshToken: tokenStore.RefreshToken,
		}
		if err := h.API.RevokeToken(context.Background(), revokeRequest); err != nil {
			return err
		}
	}

	// Remove the token store file
	if err := os.Remove(tokenStoreFile); err != nil {
		return err
	}

	return nil
}
