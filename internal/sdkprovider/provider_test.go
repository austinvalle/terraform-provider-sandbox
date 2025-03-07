package sdkprovider_test

import (
	"context"
	"testing"

	"github.com/austinvalle/terraform-provider-sandbox/internal/sdkprovider"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

var providerFactories = map[string]func() (tfprotov5.ProviderServer, error){
	"examplecloud": func() (tfprotov5.ProviderServer, error) { //nolint
		return sdkprovider.New()().GRPCProvider(), nil
	},
}

func TestProvider_ExplictlySetToTrue(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				provider "examplecloud" {
					set_namespace_from_token = true
				}
					
				resource "examplecloud_sdkv2_thing" "test" {}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"examplecloud_sdkv2_thing.test",
						tfjsonpath.New("computed_attr"),
						knownvalue.StringExact("The value passed into the provider config was true"),
					),
				},
			},
		},
	})
}
func TestProvider_ConfigureStuff(t *testing.T) {
	t.Parallel()

	rd := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"set_namespace_from_token": {
			Type:     schema.TypeBool,
			Optional: true,
			// Default: true,
		},
	},
		map[string]interface{}{})

	sdkprovider.ConfigureStuff(context.Background(), rd)
}

func TestProvider_Unset(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				provider "examplecloud" {}
					
				resource "examplecloud_sdkv2_thing" "test" {}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"examplecloud_sdkv2_thing.test",
						tfjsonpath.New("computed_attr"),
						knownvalue.StringExact("The value passed into the provider config was true"),
					),
				},
			},
		},
	})
}

func TestProvider_NoProviderConfig(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "examplecloud_sdkv2_thing" "test" {}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"examplecloud_sdkv2_thing.test",
						tfjsonpath.New("computed_attr"),
						knownvalue.StringExact("The value passed into the provider config was true"),
					),
				},
			},
		},
	})
}

func TestProvider_ExplictlySetToFalse(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				provider "examplecloud" {
					set_namespace_from_token = false
				}
					
				resource "examplecloud_sdkv2_thing" "test" {}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"examplecloud_sdkv2_thing.test",
						tfjsonpath.New("computed_attr"),
						knownvalue.StringExact("The value passed into the provider config was false"),
					),
				},
			},
		},
	})
}
