package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/oauth1/oauth"
	"gopkg.in/urfave/cli.v1"
)

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

func getToken(consumerKey, consumerSecret string) {
	oauthClient.Credentials.Token = consumerKey
	oauthClient.Credentials.Secret = consumerSecret

	tempCred, err := oauthClient.RequestTemporaryCredentials(nil, "oob", nil)
	if err != nil {
		log.Fatal("RequestTemporaryCredentials:", err)
	}

	u := oauthClient.AuthorizationURL(tempCred, nil)
	fmt.Printf("1. Go to %s\n2. Authorize the application\n3. Enter the supplied PIN here (and press Enter):\n", u)

	var code string
	fmt.Scanln(&code)

	tokenCred, _, err := oauthClient.RequestToken(nil, tempCred, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Access Token: %s\nToken Secret: %s\n", tokenCred.Token, tokenCred.Secret)
}

func main() {
	app := cli.NewApp()
	app.Usage = "Twitter PIN-based OAuth CLI helper"
	app.Version = "1.0.2"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "consumer-key, key",
			Usage: "App's consumer key",
		},
		cli.StringFlag{
			Name:  "consumer-secret, secret",
			Usage: "App's consumer secret",
		},
	}

	app.Action = func(c *cli.Context) error {
		consumerKey := c.String("key")
		consumerSecret := c.String("secret")

		if consumerKey == "" || consumerSecret == "" {
			log.Fatal("please specify both -key and -secret params (-h for help)")
		}

		getToken(consumerKey, consumerSecret)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

