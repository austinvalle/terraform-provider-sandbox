package provider

import (
	"context"
	"encoding/json"

	"github.com/austinvalle/terraform-provider-sandbox/internal/jsontypes"
	"github.com/austinvalle/terraform-provider-sandbox/internal/nettypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*thingResource)(nil)

func NewThingResource() resource.Resource {
	return &thingResource{}
}

type thingResource struct{}

type thingResourceModel struct {
	// jsontypes
	JsonBefore     types.String         `tfsdk:"json_before"`
	JsonExact      jsontypes.Exact      `tfsdk:"json_exact"`
	JsonNormalized jsontypes.Normalized `tfsdk:"json_normalized"`

	// nettypes
	IPv4AddressBefore types.String `tfsdk:"ipv4_address_before"`
	IPv6AddressBefore types.String `tfsdk:"ipv6_address_before"`

	IPv4Address nettypes.IPv4Address `tfsdk:"ipv4_address"`
	IPv6Address nettypes.IPv6Address `tfsdk:"ipv6_address"`
}

func (r *thingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// jsontypes
			"json_before": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"json_exact": schema.StringAttribute{
				CustomType: jsontypes.ExactType{},
				Optional:   true,
				Computed:   true,
			},
			"json_normalized": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
				Computed:   true,
			},

			// nettypes
			"ipv4_address_before": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ipv6_address_before": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ipv4_address": schema.StringAttribute{
				CustomType: nettypes.IPv4AddressType{},
				Optional:   true,
				Computed:   true,
			},
			"ipv6_address": schema.StringAttribute{
				CustomType: nettypes.IPv6AddressType{},
				Optional:   true,
				Computed:   true,
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type Example struct {
	Hello   string  `json:"hello"`
	Num     int64   `json:"num"`
	Decimal float64 `json:"decimal"`
	Itdobe  bool    `json:"itdobe"`
	Obj     Obj     `json:"obj"`
	Arr     []Arr   `json:"arr"`
}

type Arr struct {
	Ayo string `json:"ayo"`
}

type Obj struct {
	Wewant string `json:"wewant"`
}

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Normalized simulation of JSON
	var obj Example
	//nolint
	json.Unmarshal([]byte(data.JsonNormalized.ValueString()), &obj)
	normalizedJSON, _ := json.Marshal(obj)
	data.JsonNormalized = jsontypes.NewNormalizedValue(string(normalizedJSON))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
