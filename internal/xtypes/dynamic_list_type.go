package xtypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable  = (*DynamicListType)(nil)
	_ xattr.TypeWithValidate = (*DynamicListType)(nil)
)

type DynamicListType struct {
	basetypes.ListType
}

func (t DynamicListType) String() string {
	return "xtypes.DynamicList"
}

func (l DynamicListType) ElementType() attr.Type {
	return basetypes.DynamicType{}
}

func (l DynamicListType) WithElementType(typ attr.Type) attr.TypeWithElementType {
	return DynamicListType{
		ListType: basetypes.ListType{
			ElemType: basetypes.DynamicType{},
		},
	}
}

func (t DynamicListType) ValueType(ctx context.Context) attr.Value {
	return DynamicList{}
}

func (t DynamicListType) Equal(o attr.Type) bool {
	_, ok := o.(DynamicListType)

	return ok
}

func (t DynamicListType) Validate(ctx context.Context, in tftypes.Value, path path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if in.Type() == nil {
		return diags
	}

	return diags
}

func (t DynamicListType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	return DynamicList{
		ListValue: in,
	}, nil
}

func (t DynamicListType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return basetypes.NewListNull(t.ElementType()), nil
	}

	if !in.Type().Is(tftypes.List{}) {
		return nil, fmt.Errorf("can't use %s as value of List with ElementType %T", in.String(), t.ElementType())
	}

	if !in.IsKnown() {
		return basetypes.NewListUnknown(t.ElementType()), nil
	}

	if in.IsNull() {
		return basetypes.NewListNull(t.ElementType()), nil
	}

	val := []tftypes.Value{}
	err := in.As(&val)
	if err != nil {
		return nil, err
	}
	elems := make([]attr.Value, 0, len(val))
	for _, elem := range val {
		av, err := t.ElementType().ValueFromTerraform(ctx, elem)
		if err != nil {
			return nil, err
		}
		elems = append(elems, av)
	}

	listValue := basetypes.NewListValueMust(t.ElementType(), elems)

	listValuable, diags := t.ValueFromList(ctx, listValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}
