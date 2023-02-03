package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*exampleResource)(nil)

type exampleResource struct {
	provider exampleProvider
}

func NewResource() resource.Resource {
	return &exampleResource{}
}

func (e *exampleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

type CustomListType struct {
	types.ListType
}

func (c CustomListType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	val, err := c.ListType.ValueFromTerraform(ctx, in)

	return CustomListValue{
		// unchecked type assertion
		val.(types.List),
	}, err
}

type CustomListValue struct {
	types.List
}

func (c CustomListValue) DoSomething(ctx context.Context) {
	tflog.Info(ctx, "called DoSomething on CustomListValue")
}

func (e *exampleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				CustomType: CustomListType{
					types.ListType{
						ElemType: types.StringType,
					},
				},
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
				CustomType: CustomListType{
					types.ListType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"int64_attribute": types.Int64Type,
								"list_attribute": types.ListType{
									ElemType: types.StringType,
								},
							},
						},
					},
				},
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
				CustomType: CustomListType{
					types.ListType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"bool_attribute":    types.BoolType,
								"float64_attribute": types.Float64Type,
								"int64_attribute":   types.Int64Type,
								"list_attribute": types.ListType{
									ElemType: types.StringType,
								},
								"list_nested_nested_block": types.ListType{
									ElemType: types.ObjectType{
										AttrTypes: map[string]attr.Type{
											"bool_attribute": types.BoolType,
										},
									},
								},
							},
						},
					},
				},
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

type exampleResourceData struct {
	Id types.String `tfsdk:"id"`

	// Simple/primitive attributes
	BoolAttribute    types.Bool      `tfsdk:"bool_attribute"`
	Float64Attribute types.Float64   `tfsdk:"float64_attribute"`
	Int64Attribute   types.Int64     `tfsdk:"int64_attribute"`
	ListAttribute    CustomListValue `tfsdk:"list_attribute"`
	MapAttribute     types.Map       `tfsdk:"map_attribute"`
	NumberAttribute  types.Number    `tfsdk:"number_attribute"`
	ObjectAttribute  types.Object    `tfsdk:"object_attribute"`
	SetAttribute     types.Set       `tfsdk:"set_attribute"`
	StringAttribute  types.String    `tfsdk:"string_attribute"`

	// Nested Attributes
	ListNestedAttribute   CustomListValue `tfsdk:"list_nested_attribute"`
	MapNestedAttribute    types.Map       `tfsdk:"map_nested_attribute"`
	SetNestedAttribute    types.Set       `tfsdk:"set_nested_attribute"`
	SingleNestedAttribute types.Object    `tfsdk:"single_nested_attribute"`

	// Nested Blocks
	ListNestedBlock   CustomListValue `tfsdk:"list_nested_block"`
	SetNestedBlock    types.Set       `tfsdk:"set_nested_block"`
	SingleNestedBlock types.Object    `tfsdk:"single_nested_block"`
}

// Create is unmarshalling the config onto exampleResourceData and persisting to the state.
func (e *exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data exampleResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.ListAttribute.DoSomething(ctx)
	data.ListNestedAttribute.DoSomething(ctx)
	data.ListNestedBlock.DoSomething(ctx)

	data.Id = types.StringValue("example-id")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Read is returning the contents of the state for this resource and the State field within resource.ReadResponse
// is pre-populated so no action is required in this function.
func (e *exampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update overwrites the state with the plan. This is required in order that optional attributes in the config
// are updated in place.
func (e *exampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data exampleResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

// Delete is automatically handled by the Framework which sets an empty state. This function does not need to
// be populated unless additional handling needs to be implemented over and above setting an empty state.
func (e *exampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
