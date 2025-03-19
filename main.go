package main

import (
	"flag"
	"log"

	"github.com/austinvalle/terraform-provider-sandbox/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs --provider-name=examplecloud

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	debugEnvFilePath := flag.String("debug-env-file", "", "Path to the debug environment file to which reattach config gets written.")
	flag.Parse()

	// plugin.Serve(&plugin.ServeOpts{
	// 	Debug:        debug,
	// 	ProviderAddr: addr,
	// 	ProviderFunc: sdkprovider.New(),
	// })

	var serveOpts []tf5server.ServeOpt

	if *debugFlag {
		serveOpts = append(
			serveOpts,
			tf5server.WithManagedDebug(),
		)
	}

	if *debugEnvFilePath != "" {
		if !*debugFlag {
			log.Fatalf("debug environment file path provided without debug flag, please also set -debug")
		}

		serveOpts = append(
			serveOpts,
			tf5server.WithManagedDebugEnvFilePath(*debugEnvFilePath),
		)
	}

	err := tf5server.Serve(
		"registry.terraform.io/austinvalle/sandbox",
		providerserver.NewProtocol5(provider.New()()),
		serveOpts...,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
