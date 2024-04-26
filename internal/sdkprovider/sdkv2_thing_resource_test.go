package sdkprovider

import (
	"testing"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestSDKv2ThingResource_DeferredAction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "registry.terraform.io/hashicorp/random",
			},
		},
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"examplecloud": func() (tfprotov5.ProviderServer, error) { //nolint
				return New().GRPCProvider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				AdditionalCLIFlags: allowDeferralFlags(),
				Config:             configWithDeferredResource(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectDeferredReason("examplecloud_sdkv2_thing.test", plancheck.DeferredReasonProviderConfigUnknown),
					},
				},
				// The refresh plan will always have changes because of the deferred action (╯°□°)╯︵ ┻━┻
				ExpectNonEmptyPlan: true,
			},
			{
				AdditionalCLIFlags: allowDeferralFlags(),
				Config:             configWithDeferredResource(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
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

func allowDeferralFlags() resource.AdditionalCLIFlags {
	return resource.AdditionalCLIFlags{
		Apply: []tfexec.ApplyOption{
			tfexec.AllowDeferral(true),
		},
		Plan: []tfexec.PlanOption{
			tfexec.AllowDeferral(true),
		},
	}
}
