package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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

	// iptypes
	IPv4AddressBefore types.String        `tfsdk:"ipv4_address_before"`
	IPv6AddressBefore types.String        `tfsdk:"ipv6_address_before"`
	IPv4Address       iptypes.IPv4Address `tfsdk:"ipv4_address"`
	IPv6Address       iptypes.IPv6Address `tfsdk:"ipv6_address"`

	// cidrtypes
	IPv4CidrBefore types.String         `tfsdk:"ipv4_cidr_before"`
	IPv6CidrBefore types.String         `tfsdk:"ipv6_cidr_before"`
	IPv4Cidr       cidrtypes.IPv4Prefix `tfsdk:"ipv4_cidr"`
	IPv6Cidr       cidrtypes.IPv6Prefix `tfsdk:"ipv6_cidr"`
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

			// iptypes
			"ipv4_address_before": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ipv6_address_before": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ipv4_address": schema.StringAttribute{
				CustomType: iptypes.IPv4AddressType{},
				Optional:   true,
				Computed:   true,
			},
			"ipv6_address": schema.StringAttribute{
				CustomType: iptypes.IPv6AddressType{},
				Optional:   true,
				Computed:   true,
			},

			// cidrtypes
			"ipv4_cidr_before": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ipv6_cidr_before": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"ipv4_cidr": schema.StringAttribute{
				CustomType: cidrtypes.IPv4PrefixType{},
				Optional:   true,
				Computed:   true,
			},
			"ipv6_cidr": schema.StringAttribute{
				CustomType: cidrtypes.IPv6PrefixType{},
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

func (r *thingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Simulate normalizing JSON by re-ordering properties
	var obj any
	data.JsonNormalized.Unmarshal(&obj)
	normalizedJSON, _ := json.Marshal(obj)
	data.JsonNormalized = jsontypes.NewNormalizedValue(string(normalizedJSON))

	var obj1 any
	data.JsonExact.Unmarshal(&obj1)
	normalizedJSON1, _ := json.Marshal(obj1)
	data.JsonExact = jsontypes.NewExactValue(string(normalizedJSON1))

	// Simulate normalizing IPv6 by shorthanding/expanding
	if data.IPv6Address.ValueString() == "0:0:0:0:0:0:0:0" {
		data.IPv6Address = iptypes.NewIPv6AddressValue("::")
	} else if data.IPv6Address.ValueString() == "::1" {
		data.IPv6Address = iptypes.NewIPv6AddressValue("0:0:0:0:0:0:0:1")
	} else if data.IPv6Address.ValueString() == "0:0:0:0:0:FFFF:129.144.52.38" {
		data.IPv6Address = iptypes.NewIPv6AddressValue("::FFFF:129.144.52.38")
	}

	// Simulate normalizing IPv6 CIDR by shorthanding/expanding
	if data.IPv6Cidr.ValueString() == "2001:db8:0:0:0:0:0:0/117" {
		data.IPv6Cidr = cidrtypes.NewIPv6PrefixValue("2001:db8::/117")
	} else if data.IPv6Cidr.ValueString() == "2001:db8::/115" {
		data.IPv6Cidr = cidrtypes.NewIPv6PrefixValue("2001:db8:0:0:0:0:0:0/115")
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *thingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data thingResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
