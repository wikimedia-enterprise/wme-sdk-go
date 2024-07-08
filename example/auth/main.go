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
