// Package auth provides simple client for WME authentication.
package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
