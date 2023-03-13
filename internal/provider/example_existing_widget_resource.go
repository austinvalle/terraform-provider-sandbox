package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*exampleExistingWidgetResource)(nil)

type exampleExistingWidgetResource struct{}

func NewExistingWidgetResource() resource.Resource {
	return &exampleExistingWidgetResource{}
}

func (e *exampleExistingWidgetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_existing_widget"
}

func (e *exampleExistingWidgetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// ... other configuration ...

		Attributes: map[string]schema.Attribute{
			// ... other attributes ...

			"string_attribute": schema.StringAttribute{
				Required: true,
			},
		},
		DeprecationMessage: "use example_new_widget resource instead",
	}
}

type exampleExistingWidgetResourceData struct {
	// ... other attributes ...

	StringAttribute types.String `tfsdk:"string_attribute"`
}

func (e *exampleExistingWidgetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.Append(
		diag.NewErrorDiagnostic("example_existing_widget resource deprecated", "use example_new_widget resource instead"),
	)
}

func (e *exampleExistingWidgetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.Append(
		diag.NewErrorDiagnostic("example_existing_widget resource deprecated", "use example_new_widget resource instead"),
	)
}

func (e *exampleExistingWidgetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.Append(
		diag.NewErrorDiagnostic("example_existing_widget resource deprecated", "use example_new_widget resource instead"),
	)
}

func (e *exampleExistingWidgetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.Append(
		diag.NewErrorDiagnostic("example_existing_widget resource deprecated", "use example_new_widget resource instead"),
	)
}
