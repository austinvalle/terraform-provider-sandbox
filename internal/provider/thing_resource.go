package provider

import (
	"context"

	"github.com/austinvalle/terraform-provider-sandbox/internal/modifiers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*thingResource)(nil)
var _ resource.ResourceWithModifyPlan = (*thingResource)(nil)

func NewThingResource() resource.Resource {
	return &thingResource{}
}

type thingResource struct{}

type thingResourceModel struct {
	CaseInsensitiveString types.String `tfsdk:"case_insensitive_string"`
	ComputedAttrOne       types.String `tfsdk:"computed_attr_one"`
	ComputedAttrTwo       types.String `tfsdk:"computed_attr_two"`
}

func (r *thingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"case_insensitive_string": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					// Always takes the prior state value if the plan value (config) is equal w/ case-insensitive compare
					// https://github.com/oracle/terraform-provider-oci/blob/a50f7729707668d88e51e642178464a1985717ba/internal/service-framework/vault/vault_secret_resource.go#L277-L298
					modifiers.StringEqualIgnoreCasePlanModifier(),
				},
			},
			// Always set state to null
			"computed_attr_one": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					modifiers.UseStateForUnknown_Fixed(),
				},
			},
			// Create => "create value"
			// Update => "update value"
			"computed_attr_two": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					modifiers.UseStateForUnknown_Fixed(),
				},
			},
		},
	}
}

// This ModifyPlan is responsible for
func (r *thingResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Return early if we are deleting (plan is null) or creating (state is null)
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

	// After running all the provider plan modification (include case ignore), if there are no differences/updates, return.
	if req.Plan.Raw.Equal(req.State.Raw) {
		return
	}

	// An update is happening, so mark any computed attributes we know could change during apply as unknown
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("computed_attr_two"), types.StringUnknown())...)

	// We know "computed_attr_one" will not change, so we can skip that.
}

// Wrote a function that can mark any computed value as unknown
//
// func MarkComputedAsUnknown(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse, attrPath path.Path) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	var cfgVal attr.Value
// 	diags.Append(req.Config.GetAttribute(ctx, attrPath, &cfgVal)...)
// 	if diags.HasError() {
// 		return diags
// 	}

// 	// Has a config value, can't mark as unknown (for example, if the value is optional+computed)
// 	if !cfgVal.IsNull() {
// 		return diags
// 	}

// 	valType := cfgVal.Type(ctx)
// 	tfType := valType.TerraformType(ctx)
// 	newUnknownVal, err := valType.ValueFromTerraform(ctx, tftypes.NewValue(tfType, tftypes.UnknownValue))
// 	if err != nil {
// 		diags.Append(diag.NewErrorDiagnostic("error creating unknown val", err.Error()))
// 		return diags
// 	}

// 	diags.Append(resp.Plan.SetAttribute(ctx, attrPath, newUnknownVal)...)
// 	return diags
// }

func (r *thingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ComputedAttrOne = types.StringNull()
	data.ComputedAttrTwo = types.StringValue("create value")

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ComputedAttrOne = types.StringNull()

	data.ComputedAttrTwo = types.StringValue("update value")

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *thingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
