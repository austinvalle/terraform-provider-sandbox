package sdkprovider

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"default_str": {
				Default:  "hello there!",
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		CustomizeDiff: customdiff.All(
			customdiff.ValidateChange("size", func(ctx context.Context, oldVal, newVal, meta any) error {
				if newVal.(int) > 100 { //nolint
					return fmt.Errorf("size value must be less than 100")
				}
				return nil
			}),
		),

		// This allows CustomizeDiff to be called even if a deferral response will be returned.
		//
		// In the case of "size", it will validate that the value is < 100 and return an error even if a deferral is returned.
		// If the size is invalid, an error diagnostic will be displayed, rather than a deferral and an eventual diagnostic (on a later round).
		ResourceBehavior: schema.ResourceBehavior{
			ProviderDeferred: schema.ProviderDeferredBehavior{
				EnablePlanModification: true,
			},
		},
	}
}

func resourceThingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceThingRead(ctx, d, meta)
}

func resourceThingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, ok := meta.(*FakeAPIClient)

	if !ok || client == nil || client.APIKey == "" {
		return diag.FromErr(errors.New("API client isn't set! Failed to read thing data source"))
	}

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
