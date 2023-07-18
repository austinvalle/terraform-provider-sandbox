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
	_ basetypes.StringValuable = (*IPv6Address)(nil)
)

// TODO: docs.
type IPv6Address struct {
	basetypes.StringValue
}

func (v IPv6Address) Equal(o attr.Value) bool {
	other, ok := o.(IPv6Address)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v IPv6Address) Type(_ context.Context) attr.Type {
	return IPv6AddressType{}
}

func NewIPv6AddressNull() IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewIPv6AddressUnknown() IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewIPv6AddressValue(value string) IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewIPv6AddressPointerValue(value *string) IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

func (v IPv6Address) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv6Address)
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
