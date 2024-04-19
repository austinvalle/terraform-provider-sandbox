package sdkprovider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"examplecloud_sdkv2_thing": resourceThing(),
		},
	}

	return p
}
