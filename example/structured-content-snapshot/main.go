package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/api"
	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/auth"
)

func main() {
	ctx := context.Background()
	ath := auth.NewClient()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	//To get metadata on all available SC snapshots
	sct, _ := acl.GetStructuredSnapshots(ctx, &api.Request{})

	for _, sc := range sct {
		log.Println(sc.DateModified)
		log.Println(sc.Identifier)
		log.Println(sc.Size)
	}

	//To get metadata on an single SC snapshot using request parameters
	req := &api.Request{
		Filters: []*api.Filter{
			{
				Field: "in_language.identifier",
				Value: "en",
			},
		},
	}

	scs, err := acl.GetStructuredSnapshot(ctx, "enwiki_namespace_0", req) //GetStructuredContents(ctx, "Squirrel", req)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(scs.DateModified)
	log.Println(scs.Identifier)
	log.Println(scs.Size)
}
