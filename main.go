package main

import (
	"flag"

	"github.com/austinvalle/terraform-provider-sandbox/internal/sdkprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs --provider-name=examplecloud

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	addr := "registry.terraform.io/austinvalle/sandbox"

	plugin.Serve(&plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: addr,
		ProviderFunc: sdkprovider.New(),
	})

	// err := providerserver.Serve(context.Background(), provider.New(), providerserver.ServeOpts{
	// 	Address: addr,
	// 	Debug:   debug,
	// })
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
}
