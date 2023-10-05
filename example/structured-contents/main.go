package main

import (
	"context"
	"log"
	"os"

	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/api"
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

	acl := api.NewClient()
	acl.SetAccessToken(lgn.AccessToken)

	req := &api.Request{
		Fields: []string{"name", "abstract", "description"},
		Filters: []*api.Filter{
			{
				Field: "in_language.identifier",
				Value: "en",
			},
		},
	}
	scs, err := acl.GetStructuredContents(ctx, "Squirrel", req)

	if err != nil {
		log.Fatalln(err)
	}

	for _, sct := range scs {
		log.Println(sct.Name)
		log.Println(sct.Abstract)
		log.Println(sct.Description)
	}
}
