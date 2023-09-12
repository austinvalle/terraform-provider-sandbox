package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*fooResource)(nil)

func NewFooResource() resource.Resource {
	return &fooResource{}
}

type fooResource struct{}

// go sdk
// this is also the response body for later reference in provider plugin Create()/Read()
type Foo struct {
	Bar Bar `json:"bar"`
}

type Bar struct {
	Baz string `json:"baz"`
}

// provider plugin
// tf go models
type fooModel struct {
	ID  types.String `tfsdk:"id"`
	Bar types.Object `tfsdk:"bar"`
}

type barModel struct {
	Baz types.String `tfsdk:"baz"`
}

// tf model
var barModelTF = map[string]attr.Type{
	"baz": types.StringType,
}

func (r *fooResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *fooResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bar": schema.SingleNestedAttribute{
				Computed: true,
				// This will break the apply
				Default:     objectdefault.StaticValue(types.ObjectNull(barModelTF)),
				Description: "The foo attributes.",
				Attributes: map[string]schema.Attribute{
					"baz": schema.StringAttribute{
						Computed:    true,
						Description: "The baz.",
					},
				},
			},
		},
	}
}

// create
func (resource *fooResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// read Terraform plan data into the model
	var data fooModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// assume below is executing sdk successfully and returning Foo type struct; note at this point data.ID is empty string
	responseBody, _ := DoFoo(struct{}{}, "POST", data.ID.ValueString())
	tflog.Info(ctx, "foo created", map[string]any{"success": true})
	// error handling etc. omitted

	// assign response fields to schema values
	// assign bar attr as id
	data.ID = types.StringValue("test")
	// assign foo attributes
	var barConvertDiags diag.Diagnostics
	data.Bar, barConvertDiags = BarGoToTerraform(ctx, responseBody.Bar)
	resp.Diagnostics.Append(barConvertDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// more conversions here analogous to above, but the above alone is enough to trigger this error

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// read
func (resource *fooResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// determine input values
	var state fooModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// assume below is executing sdk successfully and returning Foo type struct
	// responseBody, _ := DoFoo(struct{}{}, "GET", state.ID.ValueString())
	// error handling etc. omitted

	// assign response fields to schema values
	// assign foo attributes
	// var barConvertDiags diag.Diagnostics
	// state.Bar, barConvertDiags = BarGoToTerraform(ctx, responseBody.Bar)
	// resp.Diagnostics.Append(barConvertDiags...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// more conversions here analogous to above, but the above alone is enough to trigger this error

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if !resp.Diagnostics.HasError() {
		tflog.Info(ctx, "Determined Foo information", map[string]any{"success": true})
	}
}

func (r *fooResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data fooModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *fooResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data fooModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func DoFoo(client any, method string, id string) (Foo, error) {
	return Foo{
		Bar: Bar{
			Baz: "using-object-from-now",
		},
	}, nil
}

func BarGoToTerraform(ctx context.Context, b Bar) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(ctx, barModelTF, barModel{
		Baz: types.StringValue(b.Baz),
	})
}
