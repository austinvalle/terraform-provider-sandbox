package xtypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ basetypes.SetValuable = SetOfDurationValue{}
var _ basetypes.SetValuableWithSemanticEquals = SetOfDurationValue{}

type SetOfDurationValue struct {
	basetypes.SetValue
}

func (s SetOfDurationValue) Equal(o attr.Value) bool {
	_, ok := o.(SetOfDurationValue)

	return ok
}

func (s SetOfDurationValue) Type(context.Context) attr.Type {
	return SetOfDurationType{
		SetType: basetypes.SetType{
			ElemType: basetypes.StringType{},
		},
	}
}

func (s SetOfDurationValue) SetSemanticEquals(ctx context.Context, newVal basetypes.SetValuable) (bool, diag.Diagnostics) {
	// Use prior value always
	return true, nil
}
