package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
)

func getToken() string {
	vaultAddr := os.Getenv("vaultAddr")
	vaultToken := os.Getenv("vaultToken")
	secretLocation := "api"
	config := &api.Config{
		Address: vaultAddr,
	}

	// Starts the client and logs in
	client, errClient := api.NewClient(config)
	handleError(errClient, "fatal")
	client.SetToken(vaultToken)

	// Reads the secret data
	secret, errSecret := client.Logical().Read("secret/data/" + secretLocation)
	handleError(errSecret, "fatal")

	// Maps the secret data so we can access the parts we need (the decrypted key value)
	mapping, errMap := secret.Data["data"].(map[string]interface{})
	if !errMap {
		log.Fatalln(errMap)
	}
	// We convert the key value to string and finally we return it
	token := fmt.Sprintf("%v", mapping[secretLocation])

	return token
}
