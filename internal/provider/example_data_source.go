package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*exampleDataSource)(nil)

type exampleDataSource struct {
	provider exampleProvider
}

func NewDataSource() datasource.DataSource {
	return &exampleDataSource{}
}

func (e *exampleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datasource"
}

func (e *exampleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},

			// Simple/primitive attributes
			"bool_attribute": schema.BoolAttribute{
				Optional: true,
			},

			"float64_attribute": schema.Float64Attribute{
				Optional: true,
			},

			"int64_attribute": schema.Int64Attribute{
				Optional: true,
			},

			"list_attribute": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},

			"map_attribute": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},

			"number_attribute": schema.NumberAttribute{
				Optional: true,
			},

			"object_attribute": schema.ObjectAttribute{
				Optional: true,
				AttributeTypes: map[string]attr.Type{
					"bool_attribute":    types.BoolType,
					"float64_attribute": types.Float64Type,
					"int64_attribute":   types.Int64Type,
					"list_attribute":    types.ListType{ElemType: types.StringType},
					"map_attribute":     types.MapType{ElemType: types.StringType},
					"number_attribute":  types.NumberType,
					"set_attribute":     types.ListType{ElemType: types.StringType},
					"string_attribute":  types.StringType,
				},
			},

			"set_attribute": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},

			"string_attribute": schema.StringAttribute{
				Optional: true,
			},

			// Nested Attributes
			"list_nested_attribute": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"int64_attribute": schema.Int64Attribute{
							Optional: true,
						},
						"list_attribute": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
					},
				},
			},

			"map_nested_attribute": schema.MapNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"map_attribute": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"number_attribute": schema.NumberAttribute{
							Optional: true,
						},
					},
				},
			},

			"set_nested_attribute": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object_attribute": schema.ObjectAttribute{
							Optional: true,
							AttributeTypes: map[string]attr.Type{
								"bool_attribute":    types.BoolType,
								"float64_attribute": types.Float64Type,
								"int64_attribute":   types.Int64Type,
								"list_attribute":    types.ListType{ElemType: types.StringType},
								"map_attribute":     types.MapType{ElemType: types.StringType},
								"number_attribute":  types.NumberType,
								"set_attribute":     types.ListType{ElemType: types.StringType},
								"string_attribute":  types.StringType,
							},
						},

						"set_attribute": schema.SetAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"string_attribute": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},

			"single_nested_attribute": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"bool_attribute": schema.BoolAttribute{
						Optional: true,
					},
					"float64_attribute": schema.Float64Attribute{
						Optional: true,
					},
				},
			},
		},


		// Nested Blocks
		Blocks: map[string]schema.Block{
			"list_nested_block": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bool_attribute": schema.BoolAttribute{
							Optional: true,
						},
						"float64_attribute": schema.Float64Attribute{
							Optional: true,
						},

						"int64_attribute": schema.Int64Attribute{
							Optional: true,
						},
						"list_attribute": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
					},
					Blocks: map[string]schema.Block{
						"list_nested_nested_block": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"bool_attribute": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"set_nested_block": schema.SetNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"map_attribute": schema.MapAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"number_attribute": schema.NumberAttribute{
							Optional: true,
						},
						"object_attribute": schema.ObjectAttribute{
							Optional: true,
							AttributeTypes: map[string]attr.Type{
								"bool_attribute":    types.BoolType,
								"float64_attribute": types.Float64Type,
								"int64_attribute":   types.Int64Type,
								"list_attribute":    types.ListType{ElemType: types.StringType},
								"map_attribute":     types.MapType{ElemType: types.StringType},
								"number_attribute":  types.NumberType,
								"set_attribute":     types.ListType{ElemType: types.StringType},
								"string_attribute":  types.StringType,
							},
						},
						"set_attribute": schema.SetAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
					},
					Blocks: map[string]schema.Block{
						"set_nested_nested_block": schema.SetNestedBlock{

							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"bool_attribute": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"single_nested_block": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"bool_attribute": schema.BoolAttribute{
						Optional: true,
					},
					"float64_attribute": schema.Float64Attribute{
						Optional: true,
					},
				},
			},
		},
	}
}

type exampleDataSourceData struct {
	Id types.String `tfsdk:"id"`

	BoolAttribute    types.Bool    `tfsdk:"bool_attribute"`
	Float64Attribute types.Float64 `tfsdk:"float64_attribute"`
	Int64Attribute   types.Int64   `tfsdk:"int64_attribute"`
	ListAttribute    types.List    `tfsdk:"list_attribute"`
	MapAttribute     types.Map     `tfsdk:"map_attribute"`
	NumberAttribute  types.Number  `tfsdk:"number_attribute"`
	ObjectAttribute  types.Object  `tfsdk:"object_attribute"`
	SetAttribute     types.Set     `tfsdk:"set_attribute"`
	StringAttribute  types.String  `tfsdk:"string_attribute"`

	// Nested Attributes
	ListNestedAttribute   types.List   `tfsdk:"list_nested_attribute"`
	MapNestedAttribute    types.Map    `tfsdk:"map_nested_attribute"`
	SetNestedAttribute    types.Set    `tfsdk:"set_nested_attribute"`
	SingleNestedAttribute types.Object `tfsdk:"single_nested_attribute"`

	// Nested Blocks
	ListNestedBlock   types.List   `tfsdk:"list_nested_block"`
	SetNestedBlock    types.Set    `tfsdk:"set_nested_block"`
	SingleNestedBlock types.Object `tfsdk:"single_nested_block"`
}

func (e *exampleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data exampleDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue("example-id")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
