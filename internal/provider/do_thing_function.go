package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &DoThingFunction{}

type DoThingFunction struct{}

func NewDoThingFunction() function.Function {
	return &DoThingFunction{}
}

func (f *DoThingFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "check_things"
}

func (f *DoThingFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		VariadicParameter: function.DynamicParameter{
			AllowNullValue: true,
		},
		Return: function.StringReturn{},
	}
}

func (f *DoThingFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var dynValues []types.Dynamic
	resp.Error = req.Arguments.Get(ctx, &dynValues)
	if resp.Error != nil {
		return
	}

	var valueTypes []string
	// Grab all the concrete values of the variadic element
	for _, dynValue := range dynValues {
		if dynValue.IsNull() || dynValue.IsUnknown() {
			continue
		}

		switch dynValue.UnderlyingValue().(type) {
		case types.String:
			// Handle string
		case types.List:
			// Handle list
		}

		valueTypes = append(valueTypes, dynValue.UnderlyingValue().Type(ctx).String())
	}

	returnVal := types.StringValue(
		fmt.Sprintf("varparam: [%s]", strings.Join(valueTypes, ", ")),
	)

	resp.Error = resp.Result.Set(ctx, returnVal)
}
