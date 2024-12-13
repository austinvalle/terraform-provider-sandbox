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
					Default:  true,
				},
			},
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap: map[string]*schema.Resource{
				"examplecloud_sdkv2_thing": thingResource(),
			},
			ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
				val := d.Get("set_namespace_from_token").(bool) //nolint
				return FakeMeta{SetNameSpaceFromToken: val}, nil
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
