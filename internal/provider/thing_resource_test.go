package provider_test

import (
	"testing"

	"github.com/austinvalle/terraform-provider-sandbox/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestThingResource(t *testing.T) {
	t.Parallel()

	r.UnitTest(t, r.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"examplecloud": providerserver.NewProtocol6WithError(provider.New()()),
		},
		Steps: []r.TestStep{
			{
				Config: `resource "examplecloud_thing" "this" {}`,
				// This plan check asserts that the value is unknown before apply
				// ConfigPlanChecks: r.ConfigPlanChecks{
				// 	PreApply: []plancheck.PlanCheck{
				// 		plancheck.ExpectUnknownValue("examplecloud_thing.this", tfjsonpath.New("bar")),
				// 	},
				// },
				// Assert that the computed value has been updated in state after apply
				Check: r.TestCheckResourceAttr("examplecloud_thing.this", "bar.baz", "baz-for-the-foo"),
			},
		},
	})
}
