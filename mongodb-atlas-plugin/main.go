package main

import (
	"log"
	"os"

	atlas "github.com/dreamteam-gg/vault-atlas-plugin"
	"github.com/hashicorp/vault/helper/pluginutil"
)

func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	err := atlas.Run(apiClientMeta.GetTLSConfig())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
