package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*thingResource)(nil)

func NewThingResource() resource.Resource {
	return &thingResource{}
}

type thingResource struct{}

type thingResourceModel struct {
	ConfigAttr   types.String `tfsdk:"config_attr"`
	ComputedAttr types.String `tfsdk:"computed_attr"`

	IPAddress iptypes.IPv4Address `tfsdk:"ip_address"`
}

func (r *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Example resource.",
		Attributes: map[string]schema.Attribute{
			"config_attr": schema.StringAttribute{
				Required: true,
			},
			"computed_attr": schema.StringAttribute{
				Computed: true,
			},
			"ip_address": schema.StringAttribute{
				Optional:   true,
				CustomType: iptypes.IPv4AddressType{},
			},
		},
	}
}

func (r *thingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Example 1: Getting all data from the plan
	var data thingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Example 2: Getting a single attribute from the plan with a framework type
	var fwConfigAttr types.String // <- target == framework type
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("config_attr"), &fwConfigAttr)...)
	if resp.Diagnostics.HasError() {
		return
	}

	fmt.Println(fwConfigAttr)

	// Example 3: Getting a single attribute from the plan with a custom type
	var fwIPAddr iptypes.IPv4Address // <- target == custom type (defined by package outside of framework)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("ip_address"), &fwIPAddr)...)
	if resp.Diagnostics.HasError() {
		return
	}

	fmt.Println(fwIPAddr)

	// Example 4: Getting a single attribute from the plan with a Go native type
	var stringIPAddr *string // <- target == Go native type that supports nil (pointer)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("ip_address"), &stringIPAddr)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if stringIPAddr != nil {
		fmt.Println(*stringIPAddr)
	}

	// Typically would call an API here, get some returned data

	// Set computed attributes from API
	data.ComputedAttr = types.StringValue("computed value")

	// Set to state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ComputedAttr = types.StringValue("computed value")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ComputedAttr = types.StringValue("computed value")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
