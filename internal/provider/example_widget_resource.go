package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*exampleWidgetResource)(nil)

type exampleWidgetResource struct{}

func NewWidgetResource() resource.Resource {
	return &exampleWidgetResource{}
}

func (e *exampleWidgetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_widget"
}

func (e *exampleWidgetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// ... other configuration ...

		Attributes: map[string]schema.Attribute{
			// ... other attributes ...

			"existing_attribute": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "use new_attribute instead",
			},

			"new_attribute": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(path.Expressions{
						path.MatchRoot("existing_attribute"),
					}...),
				},
			},
		},
	}
}

type exampleWidgetResourceData struct {
	// ... other attributes ...

	ExistingAttribute types.String `tfsdk:"existing_attribute"`
	NewAttribute      types.String `tfsdk:"new_attribute"`
}

func (e *exampleWidgetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data exampleWidgetResourceData

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.NewAttribute.IsNull() {
		// add NewAttribute to provider create API call
	} else {
		// add ExistingAttribute to provider create API call
	}

	// ... other logic ...
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (e *exampleWidgetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// ... other logic ...
}

func (e *exampleWidgetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data exampleWidgetResourceData

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.NewAttribute.IsNull() {
		// add NewAttribute to provider create API call
	} else {
		// add ExistingAttribute to provider create API call
	}

	// ... other logic ...
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (e *exampleWidgetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// ... other logic ...
}
