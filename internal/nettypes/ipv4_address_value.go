package nettypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable = (*IPv4Address)(nil)
)

// TODO: docs. MENTION LEADING ZEROES are not allowed :) .
type IPv4Address struct {
	basetypes.StringValue
}

func (v IPv4Address) Type(_ context.Context) attr.Type {
	return IPv4AddressType{}
}

func (v IPv4Address) Equal(o attr.Value) bool {
	other, ok := o.(IPv4Address)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
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
