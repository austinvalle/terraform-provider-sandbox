package main

import (
	"context"
	"flag"
	"log"

	"github.com/austinvalle/terraform-provider-sandbox/internal/provider"
	"github.com/austinvalle/terraform-provider-sandbox/internal/sdkprovider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
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

	// err := providerserver.Serve(context.Background(), provider.New(), providerserver.ServeOpts{
	// 	Address: addr,
	// 	Debug:   debug,
	// })
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	providers := []func() tfprotov5.ProviderServer{
		providerserver.NewProtocol5(provider.New()()),
		sdkprovider.New().GRPCProvider,
	}

	muxServer, err := tf5muxserver.NewMuxServer(context.Background(), providers...)
	if err != nil {
		log.Fatalf("unable to create provider: %s", err)
	}

	var serveOpts []tf5server.ServeOpt

	if debug {
		serveOpts = append(serveOpts, tf5server.WithManagedDebug())
	}

	err = tf5server.Serve(addr, muxServer.ProviderServer, serveOpts...)
	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
