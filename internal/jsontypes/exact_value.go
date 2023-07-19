package jsontypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable = (*Exact)(nil)
)

// TODO: docs.
type Exact struct {
	basetypes.StringValue
}

func (v Exact) Type(_ context.Context) attr.Type {
	return ExactType{}
}

func (v Exact) Equal(o attr.Value) bool {
	other, ok := o.(Exact)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func NewExactNull() Exact {
	return Exact{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewExactUnknown() Exact {
	return Exact{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewExactValue(value string) Exact {
	return Exact{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewExactPointerValue(value *string) Exact {
	return Exact{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
