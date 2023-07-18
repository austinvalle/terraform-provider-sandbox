package cidrtypes

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable = (*IPv4Prefix)(nil)
)

// TODO: docs. MENTION LEADING ZEROES are not allowed :) .
type IPv4Prefix struct {
	basetypes.StringValue
}

func (v IPv4Prefix) Equal(o attr.Value) bool {
	other, ok := o.(IPv4Prefix)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v IPv4Prefix) Type(_ context.Context) attr.Type {
	return IPv4PrefixType{}
}

func NewIPv4PrefixNull() IPv4Prefix {
	return IPv4Prefix{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewIPv4PrefixUnknown() IPv4Prefix {
	return IPv4Prefix{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewIPv4PrefixValue(value string) IPv4Prefix {
	return IPv4Prefix{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewIPv4PrefixPointerValue(value *string) IPv4Prefix {
	return IPv4Prefix{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

func (v IPv4Prefix) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv4Prefix)
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
	newIpPrefix, _ := netip.ParsePrefix(newValue.ValueString())
	// TODO: better name for vIpPrefix?
	vIpPrefix, _ := netip.ParsePrefix(v.ValueString())

	// TODO: is this the correct way to perform equality?
	cidrMatch := newIpPrefix.Addr() == vIpPrefix.Addr() && newIpPrefix.Bits() == vIpPrefix.Bits()

	return cidrMatch, diags
}
