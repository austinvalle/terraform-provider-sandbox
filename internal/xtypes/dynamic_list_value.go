package xtypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListValuable = (*DynamicList)(nil)
)

type DynamicList struct {
	basetypes.ListValue
}

func (v DynamicList) Type(_ context.Context) attr.Type {
	return DynamicListType{
		ListType: basetypes.ListType{
			ElemType: basetypes.DynamicType{},
		},
	}
}

func (v DynamicList) Equal(o attr.Value) bool {
	other, ok := o.(DynamicList)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v DynamicList) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	listType := tftypes.List{ElementType: v.ElementType(ctx).TerraformType(ctx)}

	if v.ListValue.IsNull() {
		return tftypes.NewValue(tftypes.List{ElementType: tftypes.DynamicPseudoType}, nil), nil
	}

	if v.ListValue.IsUnknown() {
		return tftypes.NewValue(tftypes.List{ElementType: tftypes.DynamicPseudoType}, tftypes.UnknownValue), nil
	}

	var elemTfType tftypes.Type = tftypes.DynamicPseudoType

	// Since the element type is dynamic, the final list element type will be determined by the value.
	for _, elem := range v.Elements() {
		val, err := elem.ToTerraformValue(ctx)
		// Find the first non-dynamic value and use that as the type
		if err == nil && !val.Type().Is(tftypes.DynamicPseudoType) {
			elemTfType = val.Type()
			break
		}
	}

	vals := make([]tftypes.Value, 0, len(v.Elements()))

	for _, elem := range v.Elements() {
		val, err := elem.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(listType, tftypes.UnknownValue), err
		}

		// If the value is an unknown/nil DynamicPseudoType, we need to append a unknown/nil that matches the concrete value type
		if val.Type().Is(tftypes.DynamicPseudoType) {
			if val.IsNull() {
				val = tftypes.NewValue(elemTfType, nil)
			} else if !val.IsKnown() {
				val = tftypes.NewValue(elemTfType, tftypes.UnknownValue)
			}
		}

		vals = append(vals, val)
	}

	if err := tftypes.ValidateValue(listType, vals); err != nil {
		return tftypes.NewValue(listType, tftypes.UnknownValue), err
	}

	return tftypes.NewValue(listType, vals), nil
}

func NewDynamicListNull() DynamicList {
	return DynamicList{
		ListValue: basetypes.NewListNull(basetypes.DynamicType{}),
	}
}

func NewDynamicListUnknown() DynamicList {
	return DynamicList{
		ListValue: basetypes.NewListUnknown(basetypes.DynamicType{}),
	}
}

func NewDynamicListValue(elements []attr.Value) (DynamicList, diag.Diagnostics) {
	listValue, diags := basetypes.NewListValue(basetypes.DynamicType{}, elements)
	if diags.HasError() {
		return NewDynamicListUnknown(), diags
	}

	return DynamicList{
		ListValue: listValue,
	}, nil
}
