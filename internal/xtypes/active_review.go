package xtypes

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.ObjectTypable = ActiveReviewType{}

type ActiveReviewType struct {
	basetypes.ObjectType
}

func (t ActiveReviewType) Equal(o attr.Type) bool {
	other, ok := o.(ActiveReviewType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ActiveReviewType) String() string {
	return "ActiveReviewType"
}

func (t ActiveReviewType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return nil, diags
	}

	descriptionVal, ok := descriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be basetypes.StringValue, was: %T`, descriptionAttribute))
	}

	reviewIdAttribute, ok := attributes["review_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`review_id is missing from object`)

		return nil, diags
	}

	reviewIdVal, ok := reviewIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`review_id expected to be basetypes.StringValue, was: %T`, reviewIdAttribute))
	}

	reviewStatusAttribute, ok := attributes["review_status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`review_status is missing from object`)

		return nil, diags
	}

	reviewStatusVal, ok := reviewStatusAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`review_status expected to be basetypes.StringValue, was: %T`, reviewStatusAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ActiveReviewValue{
		Description:  descriptionVal,
		ReviewId:     reviewIdVal,
		ReviewStatus: reviewStatusVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewActiveReviewValueNull() ActiveReviewValue {
	return ActiveReviewValue{
		state: attr.ValueStateNull,
	}
}

func NewActiveReviewValueUnknown() ActiveReviewValue {
	return ActiveReviewValue{
		state: attr.ValueStateUnknown,
	}
}

func NewActiveReviewValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ActiveReviewValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ActiveReviewValue Attribute Value",
				"While creating a ActiveReviewValue value, a missing attribute value was detected. "+
					"A ActiveReviewValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ActiveReviewValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ActiveReviewValue Attribute Type",
				"While creating a ActiveReviewValue value, an invalid attribute value was detected. "+
					"A ActiveReviewValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ActiveReviewValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ActiveReviewValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ActiveReviewValue Attribute Value",
				"While creating a ActiveReviewValue value, an extra attribute value was detected. "+
					"A ActiveReviewValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ActiveReviewValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewActiveReviewValueUnknown(), diags
	}

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return NewActiveReviewValueUnknown(), diags
	}

	descriptionVal, ok := descriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be basetypes.StringValue, was: %T`, descriptionAttribute))
	}

	reviewIdAttribute, ok := attributes["review_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`review_id is missing from object`)

		return NewActiveReviewValueUnknown(), diags
	}

	reviewIdVal, ok := reviewIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`review_id expected to be basetypes.StringValue, was: %T`, reviewIdAttribute))
	}

	reviewStatusAttribute, ok := attributes["review_status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`review_status is missing from object`)

		return NewActiveReviewValueUnknown(), diags
	}

	reviewStatusVal, ok := reviewStatusAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`review_status expected to be basetypes.StringValue, was: %T`, reviewStatusAttribute))
	}

	if diags.HasError() {
		return NewActiveReviewValueUnknown(), diags
	}

	return ActiveReviewValue{
		Description:  descriptionVal,
		ReviewId:     reviewIdVal,
		ReviewStatus: reviewStatusVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewActiveReviewValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ActiveReviewValue {
	object, diags := NewActiveReviewValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewActiveReviewValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ActiveReviewType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewActiveReviewValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewActiveReviewValueUnknown(), nil
	}

	if in.IsNull() {
		return NewActiveReviewValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewActiveReviewValueMust(ActiveReviewValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ActiveReviewType) ValueType(ctx context.Context) attr.Value {
	return ActiveReviewValue{}
}

var _ basetypes.ObjectValuable = ActiveReviewValue{}

type ActiveReviewValue struct {
	Description  basetypes.StringValue `tfsdk:"description"`
	ReviewId     basetypes.StringValue `tfsdk:"review_id"`
	ReviewStatus basetypes.StringValue `tfsdk:"review_status"`
	state        attr.ValueState
}

func (v ActiveReviewValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["description"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["review_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["review_status"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.Description.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["description"] = val

		val, err = v.ReviewId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["review_id"] = val

		val, err = v.ReviewStatus.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["review_status"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ActiveReviewValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ActiveReviewValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ActiveReviewValue) String() string {
	return "ActiveReviewValue"
}

func (v ActiveReviewValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		return types.ObjectNull(v.AttributeTypes(ctx)), nil
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(v.AttributeTypes(ctx)), nil
	}

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"description":   basetypes.StringType{},
			"review_id":     basetypes.StringType{},
			"review_status": basetypes.StringType{},
		},
		map[string]attr.Value{
			"description":   v.Description,
			"review_id":     v.ReviewId,
			"review_status": v.ReviewStatus,
		})

	return objVal, diags
}

func (v ActiveReviewValue) Equal(o attr.Value) bool {
	other, ok := o.(ActiveReviewValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Description.Equal(other.Description) {
		return false
	}

	if !v.ReviewId.Equal(other.ReviewId) {
		return false
	}

	if !v.ReviewStatus.Equal(other.ReviewStatus) {
		return false
	}

	return true
}

func (v ActiveReviewValue) Type(ctx context.Context) attr.Type {
	return ActiveReviewType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ActiveReviewValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"description":   basetypes.StringType{},
		"review_id":     basetypes.StringType{},
		"review_status": basetypes.StringType{},
	}
}
