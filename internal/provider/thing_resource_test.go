package provider_test

import (
	"testing"

	"github.com/austinvalle/terraform-provider-sandbox/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	r "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func Test(t *testing.T) {
	t.Parallel()

	r.UnitTest(t, r.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"examplecloud": providerserver.NewProtocol6WithError(provider.New()()),
		},
		Steps: []r.TestStep{
			{
				ConfigPlanChecks: r.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectUnknownValue("examplecloud_thing.this", tfjsonpath.New("bar")),
					},
				},
				Config: `resource "examplecloud_thing" "this" {}`,
			},
		},
	})
}
