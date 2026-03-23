package provider

import (
	"context"

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

	BreakDataConsistency types.Bool `tfsdk:"break_data_consistency"`
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
			"break_data_consistency": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}

func (r *thingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// This intentionally produces an invalid resource state by changing a user-configured value
	if data.BreakDataConsistency.ValueBool() {
		data.ConfigAttr = types.StringValue(data.ConfigAttr.ValueString() + " - data from the provider!")
	}

	data.ComputedAttr = types.StringValue("computed value")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// This intentionally produces an invalid resource state by changing a user-configured value
	if data.BreakDataConsistency.ValueBool() {
		data.ConfigAttr = types.StringValue(data.ConfigAttr.ValueString() + " - data from the provider!")
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

	// This intentionally produces an invalid resource state by changing a user-configured value
	if data.BreakDataConsistency.ValueBool() {
		data.ConfigAttr = types.StringValue(data.ConfigAttr.ValueString() + " - data from the provider!")
	}

	data.ComputedAttr = types.StringValue("computed value")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
