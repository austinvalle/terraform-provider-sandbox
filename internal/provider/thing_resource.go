package provider

import (
	"context"

	"github.com/austinvalle/terraform-provider-sandbox/internal/modifiers"
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
					// https://github.com/oracle/terraform-provider-oci/blob/a50f7729707668d88e51e642178464a1985717ba/internal/service-framework/vault/vault_secret_resource.go#L277-L298
					modifiers.StringEqualIgnoreCasePlanModifier(),
				},
			},
			// Always set state to null
			"computed_attr_one": schema.StringAttribute{
				Computed: true,
				// PlanModifiers: []planmodifier.String{
				// 	modifiers.UseStateNoMatterWhatForUnknown(), // option 3
				// 	// stringplanmodifier.UseStateForUnknown(), // The state value will be null, so this won't work.
				// },
			},
			// Always set state to "known value"
			"computed_attr_two": schema.StringAttribute{
				Computed: true,
				// PlanModifiers: []planmodifier.String{
				// 	stringplanmodifier.UseStateForUnknown(),
				// },
			},
		},
	}
}

// ModifyPlan implements resource.ResourceWithModifyPlan.
func (r *thingResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Plan = has the computed unknown values set in plan, but the rest of the plan is equal
	// - computed_attr_one = <unknown>
	// - computed_attr_two = <unknown>
	// - case_insensitive_string = "test" (prior state value, was modified from "Test" by "modifiers.StringEqualIgnoreCasePlanModifier()")

	// State = has the correct state values, rest of plan is equal
	// - computed_attr_one = <null>
	// - computed_attr_two = "known value"
	// - case_insensitive_string = "test"

	// Desired Plan/State
	// - computed_attr_one = <null>
	// - computed_attr_two = "known value"
	// - case_insensitive_string = "test"

	// Option 1. Walk raw plan to check if Plan == State
	// 		(noop, minus the unknowns that were marked by framework - this could cause problems if the unknowns were marked by the provider itself)
	//   a. If it's a noop, take the entire prior state (effectively undoing the framework's computed marking)
	//
	// Option 2. Blindly mark all computed attributes (with config value of null) as their prior state - bad idea
	//
	// Option 3. Add new plan modifier that will carry over null values as well (modifiers.UseStateNoMatterWhatForUnknown())
}

func (r *thingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ComputedAttrOne = types.StringNull()
	data.ComputedAttrTwo = types.StringValue("known value")

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ComputedAttrOne = types.StringNull()
	data.ComputedAttrTwo = types.StringValue("known value")

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
