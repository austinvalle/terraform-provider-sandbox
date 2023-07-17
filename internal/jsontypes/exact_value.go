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

func (v Exact) Equal(o attr.Value) bool {
	other, ok := o.(Exact)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v Exact) Type(_ context.Context) attr.Type {
	return ExactType{}
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

// StringSemanticEquals checks if two Exact objects have
// equivalent values, even if they are not equal.
// func (v Exact) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
// 	var diags diag.Diagnostics

// 	newValue, ok := newValuable.(Exact)
// 	if !ok {
// 		diags.AddError(
// 			"Semantic Equality Check Error",
// 			fmt.Sprintf("Expected value type %T but got value type %T. Please report this to the provider developers.", v, newValuable),
// 		)

// 		return false, diags
// 	}

// 	result, err := JSONEqual(newValue.ValueString(), v.ValueString())

// 	if err != nil {
// 		diags.AddError(
// 			"Semantic Equality Check Error",
// 			err.Error(),
// 		)

// 		return false, diags
// 	}

// 	return result, diags
// }

// func JSONBytesEqual(a, b []byte) (bool, error) {
// 	var j, j2 interface{}
// 	if err := json.Unmarshal(a, &j); err != nil {
// 		return false, err
// 	}
// 	if err := json.Unmarshal(b, &j2); err != nil {
// 		return false, err
// 	}
// 	return reflect.DeepEqual(j2, j), nil
// }

// func JSONEqual(a, b string) (bool, error) {
// 	return JSONBytesEqual([]byte(a), []byte(b))
// }
