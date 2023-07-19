package jsontypes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable = (*ExactType)(nil)
	_ xattr.TypeWithValidate  = (*ExactType)(nil)
)

// TODO: docs.
type ExactType struct {
	basetypes.StringType
}

func (t ExactType) ValueType(ctx context.Context) attr.Value {
	return Exact{}
}

func (t ExactType) String() string {
	return "jsontypes.ExactType"
}

func (t ExactType) Equal(o attr.Type) bool {
	other, ok := o.(ExactType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t ExactType) Validate(ctx context.Context, value tftypes.Value, valuePath path.Path) diag.Diagnostics {
	if value.IsNull() || !value.IsKnown() {
		return nil
	}

	var diags diag.Diagnostics
	var valueString string

	if err := value.As(&valueString); err != nil {
		diags.AddAttributeError(
			valuePath,
			"Invalid Terraform Value",
			"An unexpected error occurred while attempting to convert a Terraform value to a string. "+
				"This generally is an issue with the provider schema implementation. "+
				"Please contact the provider developers.\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Error: "+err.Error(),
		)

		return diags
	}

	if ok := json.Valid([]byte(valueString)); !ok {
		// TODO: error message clean-up
		diags.AddAttributeError(
			valuePath,
			"Invalid JSON String Value",
			"A string value was provided that is not valid JSON string format (RFC 7159).\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n",
		)

		return diags
	}

	return diags
}

func (t ExactType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return Exact{
		StringValue: in,
	}, nil
}

func (t ExactType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}
