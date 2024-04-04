package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestThingResource_no_replace(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"examplecloud": providerserver.NewProtocol6WithError(New()()),
		},
		Steps: []resource.TestStep{
			{
				// Initial creation
				Config: `resource "examplecloud_thing" "test" {
					name = "john"
					cities = {
					  season = "spring"
					}
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("id"), knownvalue.StringExact("123")),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("name"), knownvalue.StringExact("john")),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("cities"), knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"season":   knownvalue.StringExact("spring"),
							"computed": knownvalue.StringExact("true"),
						},
					)),
				},
			},
			{
				Config: `resource "examplecloud_thing" "test" {
					name = "bob"
					cities = {
					  season = "spring"
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Should cause an update in-place, as only name has changed
						plancheck.ExpectResourceAction("examplecloud_thing.test", plancheck.ResourceActionUpdate),

						// Computed value in object should use prior state
						plancheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("cities").AtMapKey("computed"), knownvalue.StringExact("true")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("id"), knownvalue.StringExact("123")),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("name"), knownvalue.StringExact("bob")),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("cities"), knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"season":   knownvalue.StringExact("spring"),
							"computed": knownvalue.StringExact("true"),
						},
					)),
				},
			},
			{
				Config: `resource "examplecloud_thing" "test" {
					name = "bob"
					cities = {
					  season = "fall"
					}
				}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// Should cause an replace (create/destroy), as cities.season has changed
						plancheck.ExpectResourceAction("examplecloud_thing.test", plancheck.ResourceActionReplace),

						// Computed value in object is unknown here because the object is being recreated
						plancheck.ExpectUnknownValue("examplecloud_thing.test", tfjsonpath.New("cities").AtMapKey("computed")),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("id"), knownvalue.StringExact("123")),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("name"), knownvalue.StringExact("bob")),
					statecheck.ExpectKnownValue("examplecloud_thing.test", tfjsonpath.New("cities"), knownvalue.ObjectExact(
						map[string]knownvalue.Check{
							"season":   knownvalue.StringExact("fall"),
							"computed": knownvalue.StringExact("true"),
						},
					)),
				},
			},
		},
	})
}
