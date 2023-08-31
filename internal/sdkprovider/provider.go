package sdkprovider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap: map[string]*schema.Resource{
				"examplecloud_thing": thingResource(),
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
			"word": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
					return strings.EqualFold(oldValue, newValue)
				},
				DiffSuppressOnRefresh: true,
			},
		},
	}
}

func thingResourceCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	d.SetId("thing-id")

	return nil
}

func thingResourceRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// This simulates an API returning the 'word' attribute as all uppercase,
	// which is stored to state even if it doesn't match the config or prior value.
	word := d.Get("word").(string)       //nolint
	d.Set("word", strings.ToUpper(word)) //nolint

	return nil
}

func thingResourceUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func thingResourceDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}
