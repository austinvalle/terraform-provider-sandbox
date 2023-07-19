package jsontypes

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable                   = (*Normalized)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*Normalized)(nil)
)

// TODO: docs.
type Normalized struct {
	basetypes.StringValue
}

func (v Normalized) Type(_ context.Context) attr.Type {
	return NormalizedType{}
}

func (v Normalized) Equal(o attr.Value) bool {
	other, ok := o.(Normalized)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v Normalized) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(Normalized)
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

	result, err := jsonEqual(newValue.ValueString(), v.ValueString())

	if err != nil {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected error occurred while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)

		return false, diags
	}

	return result, diags
}

func jsonEqual(s1, s2 string) (bool, error) {
	var j1, j2 interface{}
	if err := json.Unmarshal([]byte(s1), &j1); err != nil {
		return false, err
	}
	if err := json.Unmarshal([]byte(s2), &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j1, j2), nil
}

func NewNormalizedNull() Normalized {
	return Normalized{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewNormalizedUnknown() Normalized {
	return Normalized{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewNormalizedValue(value string) Normalized {
	return Normalized{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewNormalizedPointerValue(value *string) Normalized {
	return Normalized{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
