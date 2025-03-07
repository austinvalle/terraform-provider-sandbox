package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*thingDataSource)(nil)

func NewThingDataSource() datasource.DataSource {
	return &thingDataSource{}
}

type thingDataSource struct{}

type thingDataSourceModel struct {
	NestedMap types.Map `tfsdk:"nested_map"`
	NestedObj types.Map `tfsdk:"nested_obj"`
}

func (r *thingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"nested_map": schema.MapNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"required_attr": schema.StringAttribute{
							Required: true,
						},
						"computed_attr": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"nested_obj": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"required_attr": schema.StringAttribute{
						Required: true,
					},
					"computed_attr": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *thingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data thingDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("nested_map").AtMapKey("computed_attr"), types.StringValue("world!"))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("nested_obj").AtName("computed_attr"), types.StringValue("world!"))...)
}
