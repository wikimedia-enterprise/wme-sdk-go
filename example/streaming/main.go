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

	cbk := func(art *api.Article) error {
		log.Println("----------START-----------")
		log.Printf("name: %s\n", art.Name)
		log.Printf("abstract: %s\n", art.Abstract)
		log.Printf("event.identifiers: %s\n", art.Event.Identifier)
		log.Print("-----------END------------\n\n\n")
		return nil
	}
	req := &api.Request{
		Fields: []string{"name", "abstract", "event.*"},
		Filters: []*api.Filter{
			{
				Field: "in_language.identifier",
				Value: "en",
			},
		},
		Parts:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Offsets: map[int]int64{0: int64(45), 1: int64(48)},
	}

	if err := acl.StreamArticles(ctx, req, cbk); err != nil {
		log.Fatalln(err)
	}
}
