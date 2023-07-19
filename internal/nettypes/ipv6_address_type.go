package nettypes

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable = (*IPv6AddressType)(nil)
	_ xattr.TypeWithValidate  = (*IPv6AddressType)(nil)
)

// TODO: docs. Also supports:
// -  IPv4 compatible addresses like ::13.1.68.3
// -  IPv4 mapped addresses like ::FFFF:129.144.52.38 .
type IPv6AddressType struct {
	basetypes.StringType
}

func (t IPv6AddressType) ValueType(ctx context.Context) attr.Value {
	return IPv6Address{}
}

func (t IPv6AddressType) String() string {
	return "nettypes.IPv6AddressType"
}

func (t IPv6AddressType) Equal(o attr.Type) bool {
	other, ok := o.(IPv6AddressType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t IPv6AddressType) Validate(ctx context.Context, value tftypes.Value, valuePath path.Path) diag.Diagnostics {
	if value.IsNull() || !value.IsKnown() {
		return nil
	}

	var diags diag.Diagnostics
	var valueString string

	if err := value.As(&valueString); err != nil {
		diags.AddAttributeError(
			valuePath,
			"IPv6 Address Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)

		return diags
	}

	ipAddr, err := netip.ParseAddr(valueString)
	if err != nil {
		// TODO: error message clean-up
		diags.AddAttributeError(
			valuePath,
			"Invalid IPv6 Address String Value",
			"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n"+
				"Error: "+err.Error(),
		)

		return diags
	}

	if ipAddr.Is4() {
		// TODO: error message clean-up
		diags.AddAttributeError(
			valuePath,
			"Invalid IPv6 Address String Value",
			"An IPv4 string format was provided, string value must be IPv6 string format or IPv4-Mapped IPv6 string format (RFC 4291).\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n",
		)

		return diags
	}

	if !ipAddr.IsValid() || !ipAddr.Is6() {
		// TODO: error message clean-up
		diags.AddAttributeError(
			valuePath,
			"Invalid IPv6 Address String Value",
			"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n",
		)

		return diags
	}

	return diags
}

func (t IPv6AddressType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPv6Address{
		StringValue: in,
	}, nil
}

func (t IPv6AddressType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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