# Auth Helper Example

This repository contains an example Go application that demonstrates how to use an authentication helper to manage access and refresh tokens for making authenticated API calls easier. The `main.go` script logs in, checks the validity of tokens, and refreshes the token periodically using a background goroutine.

## Prerequisites

- Go 1.15+ installed on your machine.
- Environment variables for your username and password stored in a `.env` file.

## Setup

### 1. Clone the Repository

```sh
git clone https://github.com/wikimedia-enterprise/wme-sdk-go.git
cd wme-sdk-go
```

### 2. Environment Variables

Create a .env file with the following contents:
```bash
export WME_USERNAME="...your username...";
export WME_PASSWORD="...your password...";
```

### 3. Install Dependencies

Install the necessary Go dependencies:

```sh
go mod tidy
```

## Running the Application

To run the application, use the following command:

```sh
cd example/auth
go run main.go
```

This will start the application, log in, and display the access token. It will also start a background goroutine that refreshes the token every 23 hours and 59 seconds, or whenever an API call is made.

## Project Structure

- `main.go`: The main application file that logs in and handles token refresh.
- `auth/`: Directory containing the authentication helper code.
  - `auth.go`: Contains the authentication client and helper methods for managing tokens.

## `auth` Package

The `auth` package provides a client and helper for managing authentication tokens. Below are the key components:

### Auth Package Components

- `Client`: HTTP client for making authentication requests.
- `Helper`: Manages the state and validity of tokens.
- `Tokenstore`: Struct for storing token data.
- `LoginRequest`, `LoginResponse`, `RefreshTokenRequest`, `RefreshTokenResponse`, `RevokeTokenRequest`: Structs for making API requests and handling responses.

### Helper Methods

- `GetToken() (string, error)`: Manages the token state and returns a valid access token.
- `loginAndStoreTokens(tokenStoreFile string) (string, error)`: Logs in and stores the tokens.
- `refreshAndStoreTokens(tokenStoreFile string, tokenStore *Tokenstore) (string, error)`: Refreshes the access token and stores the new tokens.
- `storeTokens(tokenStoreFile string, tokenStore *Tokenstore) error`: Writes the token data to the tokenstore file in JSON format.

## Example Usage in `main.go`

```go
package main

import (
	"context"
	"log"
	"sync"

	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/api"
	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/auth"
)

func main() {
	// Initialize the authentication client and helper
	// We rely on authentication helper to return a valid access token for all our requests. The helper APIs will internally refresh, re-login as necessary.
	authClient := auth.NewClient()
	helper, err := auth.NewHelper(authClient)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	ctx := context.Background()
	reqs := 50
	wg.Add(reqs)

	for i := 0; i < reqs; i++ {
		go func(req int) {
			defer wg.Done()
			// Get access token
			tkn, err := helper.GetAccessToken()
			if err != nil {
				log.Printf("failed to get access token for request %d: %v", req, err)
				return
			}

			// Call any WME API
			clt := api.NewClient()
			clt.SetAccessToken(tkn)
			arq := &api.Request{
				Fields: []string{"name", "abstract", "url", "version"},
				Filters: []*api.Filter{
					{
						Field: "in_language.identifier",
						Value: "en",
					},
				},
			}

			res, err := clt.GetArticles(ctx, "Montreal", arq)
			if err != nil {
				log.Printf("failed to get articles for request %d: %v", req, err)
				return
			}

			log.Printf("request %d: %d articles found", req, len(res))

			// Your code here......

		}(i)
	}

	wg.Wait()

	// Revoke token and remove state from storage when we are done using WME APIs
	if err := helper.ClearState(); err != nil {
		log.Printf("failed to clear state: %v", err)
	}
}

```

## Contributing

If you want to contribute to this project, feel free to open issues or submit pull requests.

## License

This project is licensed under the  Apache-2.0 license License. See the LICENSE file for details.

---

Feel free to adjust the repository URL, structure, and any other details to better fit your actual project setup.
