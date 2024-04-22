package sdkprovider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FakeAPIClient struct {
	APIKey string
}

func New() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"examplecloud_sdkv2_thing": dataSourceThing(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"examplecloud_sdkv2_thing": resourceThing(),
		},
		ConfigureProvider: func(ctx context.Context, req schema.ConfigureProviderRequest, resp *schema.ConfigureProviderResponse) {
			providerConfig := req.ResourceData.GetRawConfig()

			if providerConfig.IsWhollyKnown() {
				apiKey := req.ResourceData.Get("api_key").(string) //nolint
				resp.Meta = &FakeAPIClient{
					APIKey: apiKey,
				}

				return
			}

			if !req.DeferralAllowed {
				resp.Diagnostics = diag.FromErr(errors.New("provider configuration is unknown and deferral is not allowed!"))
				return
			}

			// Provider configuration is unknown, defer all resources and data sources
			resp.DeferralResponse = &schema.DeferralResponse{
				Reason: schema.DeferralReasonProviderConfigUnknown,
			}
		},
	}

	return p
}
