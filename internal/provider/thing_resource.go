package provider

import (
	"context"

	"github.com/austinvalle/terraform-provider-sandbox/internal/xtypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*thingResource)(nil)
var _ resource.ResourceWithValidateConfig = (*thingResource)(nil)

func NewThingResource() resource.Resource {
	return &thingResource{}
}

type thingResource struct{}

// ValidateConfig implements resource.ResourceWithValidateConfig.
func (r *thingResource) ValidateConfig(context.Context, resource.ValidateConfigRequest, *resource.ValidateConfigResponse) {
	return
}

type thingResourceModel struct {
	ActiveReview xtypes.ActiveReviewValue `tfsdk:"active_review"`
}

func (r *thingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"active_review": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Required: true,
					},
					"review_id": schema.StringAttribute{
						Required: true,
					},
					"review_status": schema.StringAttribute{
						Required: true,
					},
				},
				CustomType: xtypes.ActiveReviewType{
					ObjectType: types.ObjectType{
						AttrTypes: xtypes.ActiveReviewValue{}.AttributeTypes(ctx),
					},
				},
				Optional: true,
				Computed: true,
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

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

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
