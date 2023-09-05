package main

import (
	"context"
	"log"
	"os"

	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/auth"
)

func main() {
	ctx := context.Background()
	ath := auth.NewClient()

	lgn, err := ath.Login(ctx, &auth.LoginRequest{
		Username: os.Getenv("WME_USERNAME"),
		Password: os.Getenv("WME_PASSWORD"),
	})

	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		req := &auth.RevokeTokenRequest{
			RefreshToken: lgn.RefreshToken,
		}

		if err := ath.RevokeToken(ctx, req); err != nil {
			log.Println(err)
		}
	}()

	// ...your code goes here...
}
