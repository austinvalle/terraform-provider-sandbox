package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestThingResource_cross_validation_with_unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source: "hashicorp/random",
			},
		},
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"examplecloud": providerserver.NewProtocol6WithError(New()()),
		},
		Steps: []resource.TestStep{
			{
				Config: `resource "random_string" "a" {
					length = 10
				}

				resource "examplecloud_thing" "test" {
					str_value  = random_string.a.result
					max_length = 5
				}`,
				ExpectError: regexp.MustCompile(`string length must be less then 5`),
			},
		},
	})
}
