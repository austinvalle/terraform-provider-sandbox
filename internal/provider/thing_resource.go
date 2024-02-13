package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/dynamicdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/dynamicplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*thingResource)(nil)

func NewThingResource() resource.Resource {
	return &thingResource{}
}

type thingResource struct{}

type thingResourceModel struct {
	DynamicAttrRequired         types.Dynamic `tfsdk:"dynamic_attr_required"`
	DynamicAttrComputed         types.Dynamic `tfsdk:"dynamic_attr_computed"`
	DynamicAttrOptional         types.Dynamic `tfsdk:"dynamic_attr_optional"`
	DynamicAttrComputedOptional types.Dynamic `tfsdk:"dynamic_attr_computed_optional"`
}

func (r *thingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dynamic_attr_required": schema.DynamicAttribute{
				Required: true,
			},
			"dynamic_attr_computed": schema.DynamicAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Dynamic{
					dynamicplanmodifier.UseStateForUnknown(),
				},
			},
			"dynamic_attr_optional": schema.DynamicAttribute{
				Optional: true,
			},
			"dynamic_attr_computed_optional": schema.DynamicAttribute{
				Computed: true,
				Optional: true,
				Default: dynamicdefault.StaticDynamic(
					types.DynamicValue(types.StringValue("hello default")),
				),
			},
		},
	}
}

func (r *thingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data thingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.DynamicAttrComputed = types.DynamicValue(types.BoolValue(true))
	if data.DynamicAttrComputedOptional.IsUnknown() {
		data.DynamicAttrComputedOptional = types.DynamicValue(types.StringValue("computed/optional created!"))
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data thingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.DynamicAttrComputed = types.DynamicValue(types.StringValue("it got refreshed!"))
	data.DynamicAttrComputedOptional = types.DynamicValue(types.StringValue("creating some drift"))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data thingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}
