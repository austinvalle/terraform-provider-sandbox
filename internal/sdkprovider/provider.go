package sdkprovider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FakeMeta struct {
	SetNameSpaceFromToken bool
}

func New() func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"set_namespace_from_token": {
					Type:     schema.TypeBool,
					Optional: true,
					// Default: true,
				},
			},
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap: map[string]*schema.Resource{
				"examplecloud_sdkv2_thing": thingResource(),
			},
			ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

				// Raw config is the actual data that is sent over the protocol. It uses go-cty (the same type system that core uses)
				// and is relatively awkward to use. But for retrieving a simple attribute at the root of an object, this should be fine.
				//
				// Cty favors using panics over error handling, so if the attribute name is mistyped, operations like (cty.Value).GetAttr will panic
				rawConfig := d.GetRawConfig()
				rawVal := rawConfig.GetAttr("set_namespace_from_token")

				// We don't care about the underlying value, just detecting if the config value is null (unset) or not.
				if rawVal.IsNull() {
					// The value is null (unset) in config, do our defaulting logic, which in this case is just returning a client with true
					return FakeMeta{SetNameSpaceFromToken: true}, nil
				}

				// If the value is set in config, just use d.Get like normal
				return FakeMeta{SetNameSpaceFromToken: d.Get("set_namespace_from_token").(bool)}, nil //nolint
			},
		}
		return p
	}
}
func thingResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: thingResourceCreate,
		ReadContext:   thingResourceRead,
		UpdateContext: thingResourceUpdate,
		DeleteContext: thingResourceDelete,
		Schema: map[string]*schema.Schema{
			"trigger": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "just used to trigger updates and stop the internal validate from complaining :)",
			},
			"computed_attr": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "well reflect what is in the provider config",
			},
		},
	}
}
func thingResourceCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	fakeMeta := meta.(FakeMeta) //nolint
	d.SetId("thing-123")
	d.Set("computed_attr", fmt.Sprintf("The value passed into the provider config was %t", fakeMeta.SetNameSpaceFromToken)) //nolint
	return nil
}
func thingResourceRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}
func thingResourceUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	fakeMeta := meta.(FakeMeta)                                                                                             //nolint
	d.Set("computed_attr", fmt.Sprintf("The value passed into the provider config was %t", fakeMeta.SetNameSpaceFromToken)) //nolint
	return nil
}
func thingResourceDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}
