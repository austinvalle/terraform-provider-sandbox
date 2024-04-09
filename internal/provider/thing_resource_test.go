package provider

import (
	"testing"

	"github.com/austinvalle/terraform-provider-sandbox/internal/sandboxtesting"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/zclconf/go-cty/cty"
)

func TestThingResource_list_with_dynamics(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"examplecloud": providerserver.NewProtocol6WithError(New()()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "examplecloud_thing" "test" {
					list_with_dynamics = ["hello", "world", true]
				}
				
				# Used to check the Terraform type
				output "type_test" {
					value = examplecloud_thing.test.list_with_dynamics
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					// Now it's a list of strings!
					sandboxtesting.ExpectOutputType("type_test", cty.List(cty.String)),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("list_with_dynamics"), knownvalue.ListExact(
						[]knownvalue.Check{
							knownvalue.StringExact("hello"),
							knownvalue.StringExact("world"),
							knownvalue.StringExact("true"),
						},
					)),
				},
			},
			{
				Config: `resource "examplecloud_thing" "test" {
					list_with_dynamics = [123, 456, 789]
				}
				
				# Used to check the Terraform type
				output "type_test" {
					value = examplecloud_thing.test.list_with_dynamics
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					// Now it's a list of numbers!
					sandboxtesting.ExpectOutputType("type_test", cty.List(cty.Number)),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("list_with_dynamics"), knownvalue.ListExact(
						[]knownvalue.Check{
							knownvalue.Int64Exact(123),
							knownvalue.Int64Exact(456),
							knownvalue.Int64Exact(789),
						},
					)),
				},
			},
			{
				Config: `resource "examplecloud_thing" "test" {
					list_with_dynamics = [true, false, true]
				}
				
				# Used to check the Terraform type
				output "type_test" {
					value = examplecloud_thing.test.list_with_dynamics
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					// Now it's a list of booleans!
					sandboxtesting.ExpectOutputType("type_test", cty.List(cty.Bool)),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("list_with_dynamics"), knownvalue.ListExact(
						[]knownvalue.Check{
							knownvalue.Bool(true),
							knownvalue.Bool(false),
							knownvalue.Bool(true),
						},
					)),
				},
			},
		},
	})
}
