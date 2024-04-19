package sdkprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThing() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceThingCreate,
		ReadContext:   resourceThingRead,
		UpdateContext: resourceThingUpdate,
		DeleteContext: resourceThingDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceThingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceThingRead(ctx, d, meta)
}

func resourceThingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string) //nolint
	d.SetId(name)

	return nil
}

func resourceThingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceThingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
