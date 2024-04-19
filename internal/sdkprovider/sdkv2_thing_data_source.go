package sdkprovider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceThing() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceThingRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceThingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, ok := meta.(*FakeAPIClient)

	if !ok || client == nil || client.APIKey == "" {
		return diag.FromErr(errors.New("API client isn't set! Failed to read thing data source"))
	}

	d.SetId("id-12345")
	err := d.Set("name", "austin")
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
