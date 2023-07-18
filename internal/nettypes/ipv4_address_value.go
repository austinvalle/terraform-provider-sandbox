package nettypes

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable = (*IPv4Address)(nil)
)

// TODO: docs. MENTION LEADING ZEROES are not allowed :) .
type IPv4Address struct {
	basetypes.StringValue
}

func (v IPv4Address) Equal(o attr.Value) bool {
	other, ok := o.(IPv4Address)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v IPv4Address) Type(_ context.Context) attr.Type {
	return IPv4AddressType{}
}

func NewIPv4AddressNull() IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewIPv4AddressUnknown() IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewIPv4AddressValue(value string) IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewIPv4AddressPointerValue(value *string) IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

func (v IPv4Address) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv4Address)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	// TODO: are ignoring these errors okay?
	newIpAddr, _ := netip.ParseAddr(newValue.ValueString())
	// TODO: better name for vIpAddr?
	vIpAddr, _ := netip.ParseAddr(v.ValueString())

	return vIpAddr == newIpAddr, diags
}
