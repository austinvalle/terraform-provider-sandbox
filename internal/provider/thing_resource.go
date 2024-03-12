package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	Dynamic         types.Dynamic `tfsdk:"dynamic"`
	DynamicComputed types.Dynamic `tfsdk:"dynamic_computed"`

	StaticObj  types.Object `tfsdk:"static_obj"`
	StaticList types.List   `tfsdk:"static_list"`
}

func (r *thingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dynamic": schema.DynamicAttribute{
				Optional: true,
			},
			"dynamic_computed": schema.DynamicAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Dynamic{
					dynamicplanmodifier.UseStateForUnknown(),
				},
			},
			"static_obj": schema.ObjectAttribute{
				AttributeTypes: map[string]attr.Type{
					"dynamic": types.DynamicType,
				},
				Optional: true,
			},
			"static_list": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}

func (r *thingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Set the computed attribute to a list of numbers
	data.DynamicComputed = types.DynamicValue(
		types.ListValueMust(
			types.Int64Type,
			[]attr.Value{
				types.Int64Value(1),
				types.Int64Value(2),
				types.Int64Value(3),
				types.Int64Value(4),
				types.Int64Value(5),
			},
		),
	)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Refresh the computed attribute from a list of number to a boolean
	//
	// This is totally valid on a refresh! (╯°□°)╯︵ ┻━┻
	// - Things like this could break a config :(
	// data.DynamicComputed = types.DynamicValue(types.BoolValue(true))

	// data.DynamicComputed = types.DynamicValue(
	// 	types.SetValueMust(
	// 		types.Int64Type,
	// 		[]attr.Value{
	// 			types.Int64Value(1),
	// 			types.Int64Value(2),
	// 			types.Int64Value(3),
	// 			types.Int64Value(4),
	// 		},
	// 	),
	// )

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data thingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}
