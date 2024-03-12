package provider_test

import (
	"testing"

	"github.com/austinvalle/terraform-provider-sandbox/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func protoV6ProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"examplecloud": providerserver.NewProtocol6WithError(provider.NewProvider()),
	}
}

func TestDoThingFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::examplecloud::check_things("hey", 1, true, {a = 123})
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.StringExact(`param1: basetypes.StringType, varparam: [basetypes.NumberType, basetypes.BoolType, types.ObjectType["a":basetypes.NumberType]]`)),
				},
			},
			{
				Config: `
				output "test" {
					value = provider::examplecloud::check_things("hi")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.StringExact(`param1: basetypes.StringType, varparam: []`)),
				},
			},
			{
				Config: `
				output "test" {
					value = provider::examplecloud::check_things("hello", true)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("test", knownvalue.StringExact(`param1: basetypes.StringType, varparam: [basetypes.BoolType]`)),
				},
			},
		},
	})
}
