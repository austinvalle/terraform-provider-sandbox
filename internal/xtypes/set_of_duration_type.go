package xtypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.SetTypable = SetOfDurationType{}

type SetOfDurationType struct {
	basetypes.SetType
}

func (s SetOfDurationType) Equal(o attr.Type) bool {
	_, ok := o.(SetOfDurationType)

	return ok
}

func (s SetOfDurationType) ElementType() attr.Type {
	return basetypes.StringType{}
}

func (s SetOfDurationType) WithElementType(typ attr.Type) attr.TypeWithElementType {
	return SetOfDurationType{
		SetType: basetypes.SetType{
			ElemType: basetypes.StringType{},
		},
	}
}

func (s SetOfDurationType) String() string {
	return "xtypes.SetOfDurationType"
}

func (s SetOfDurationType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := s.SetType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	setValue, ok := attrValue.(basetypes.SetValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	setValuable, diags := s.ValueFromSet(ctx, setValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return setValuable, nil
}

func (s SetOfDurationType) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	value := SetOfDurationValue{
		SetValue: in,
	}

	return value, nil
}

func (s SetOfDurationType) ValueType(ctx context.Context) attr.Value {
	return SetOfDurationValue{}
}
