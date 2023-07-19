package cidrtypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable = (*IPv4Prefix)(nil)
)

// TODO: docs. MENTION LEADING ZEROES are not allowed :) .
type IPv4Prefix struct {
	basetypes.StringValue
}

func (v IPv4Prefix) Type(_ context.Context) attr.Type {
	return IPv4PrefixType{}
}

func (v IPv4Prefix) Equal(o attr.Value) bool {
	other, ok := o.(IPv4Prefix)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
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
