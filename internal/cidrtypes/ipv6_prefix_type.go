package cidrtypes

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
	_ basetypes.StringTypable = (*IPv6PrefixType)(nil)
	_ xattr.TypeWithValidate  = (*IPv6PrefixType)(nil)
)

// TODO: docs.
type IPv6PrefixType struct {
	basetypes.StringType
}

func (t IPv6PrefixType) ValueType(ctx context.Context) attr.Value {
	return IPv6Prefix{}
}

func (t IPv6PrefixType) String() string {
	return "cidrtypes.IPv6PrefixType"
}

func (t IPv6PrefixType) Equal(o attr.Type) bool {
	other, ok := o.(IPv6PrefixType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t IPv6PrefixType) Validate(ctx context.Context, value tftypes.Value, valuePath path.Path) diag.Diagnostics {
	if value.IsNull() || !value.IsKnown() {
		return nil
	}

	var diags diag.Diagnostics
	var valueString string

	if err := value.As(&valueString); err != nil {
		diags.AddAttributeError(
			valuePath,
			"IPv6 Prefix Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)

		return diags
	}

	ipPrefix, err := netip.ParsePrefix(valueString)
	if err != nil {
		// TODO: error message clean-up
		diags.AddAttributeError(
			valuePath,
			"Invalid IPv6 CIDR String Value",
			"A string value was provided that is not valid IPv6 CIDR string format (RFC 4291).\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n"+
				"Error: "+err.Error(),
		)

		return diags
	}

	// TODO: is this correct way to determine IPv4 CIDR?
	if ipPrefix.Addr().Is4() {
		// TODO: error message clean-up
		diags.AddAttributeError(
			valuePath,
			"Invalid IPv6 CIDR String Value",
			"An IPv4 CIDR string format was provided, string value must be IPv6 CIDR string format (RFC 4291).\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n",
		)

		return diags
	}

	// TODO: is this correct way to determine IPv6 CIDR?
	if !ipPrefix.IsValid() || !ipPrefix.Addr().Is6() {
		// TODO: error message clean-up
		diags.AddAttributeError(
			valuePath,
			"Invalid IPv6 CIDR String Value",
			"A string value was provided that is not valid IPv6 CIDR string format (RFC 4291).\n\n"+
				"Path: "+valuePath.String()+"\n"+
				"Given Value: "+valueString+"\n",
		)

		return diags
	}

	return diags
}

func (t IPv6PrefixType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPv6Prefix{
		StringValue: in,
	}, nil
}

func (t IPv6PrefixType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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