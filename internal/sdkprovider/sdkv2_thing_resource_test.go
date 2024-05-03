package sdkprovider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestSDKv2ThingResource_DeferredAction(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		// Deferred action support is only available in 1.9 alpha binaries (or dev builds)
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_9_0),
			tfversion.SkipIfNotPrerelease(),
		},
		// Add the -allow-deferral flag
		AdditionalCLIOptions: &r.AdditionalCLIOptions{
			Apply: r.ApplyOptions{AllowDeferral: true},
			Plan:  r.PlanOptions{AllowDeferral: true},
		},
		ExternalProviders: map[string]r.ExternalProvider{
			"random": {
				Source: "registry.terraform.io/hashicorp/random",
			},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"examplecloud": func() (tfprotov5.ProviderServer, error) { //nolint
				return New().GRPCProvider(), nil
			},
		},
		Steps: []r.TestStep{
			{
				Config: configWithDeferredResource(),
				ConfigPlanChecks: r.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Check the exact reason the resource was deferred, for SDKv2 resources, this will always be provider config unknown
						plancheck.ExpectDeferredChange("examplecloud_sdkv2_thing.test", plancheck.DeferredReasonProviderConfigUnknown),
					},
				},
				// The refresh plan will always have changes because of the deferred action (╯°□°)╯︵ ┻━┻
				ExpectNonEmptyPlan: true,
			},
			{
				Config: configWithDeferredResource(),
				ConfigPlanChecks: r.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Ensure there are no more deferred changes
						plancheck.ExpectNoDeferredChanges(),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("examplecloud_sdkv2_thing.test", tfjsonpath.New("name"), knownvalue.StringExact("austin")),
				},
			},
		},
	})
}

func configWithDeferredResource() string {
	return `
	provider "examplecloud" {
		api_key = random_string.str.result
	}

	resource "random_string" "str" {
		length = 5
	}
	
	resource "examplecloud_sdkv2_thing" "test" {
		name = data.examplecloud_sdkv2_thing.test.name
		size = 99
	}
	
	data "examplecloud_sdkv2_thing" "test" {}
	`
}
