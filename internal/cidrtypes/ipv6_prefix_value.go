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
	_ basetypes.StringValuable = (*IPv6Prefix)(nil)
)

// TODO: docs.
type IPv6Prefix struct {
	basetypes.StringValue
}

func (v IPv6Prefix) Equal(o attr.Value) bool {
	other, ok := o.(IPv6Prefix)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v IPv6Prefix) Type(_ context.Context) attr.Type {
	return IPv6PrefixType{}
}

func NewIPv6PrefixNull() IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewIPv6PrefixUnknown() IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewIPv6PrefixValue(value string) IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewIPv6PrefixPointerValue(value *string) IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}

func (v IPv6Prefix) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv6Prefix)
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
