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
		Description: "This function has one static type parameter and a dynamic variadic parameter. The return is a " +
			"string that contains all the concrete values passed to the function",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name: "param1",
			},
		},
		VariadicParameter: function.DynamicParameter{
			Name: "varparam",
		},
		Return: function.StringReturn{},
	}
}

func (f *DoThingFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var stringParam types.String
	var variadicParam types.Tuple
	resp.Error = req.Arguments.Get(ctx, &stringParam, &variadicParam)
	if resp.Error != nil {
		return
	}

	var valueTypes []string
	// Grab all the concrete values of the variadic element
	for _, value := range variadicParam.Elements() {
		dynValue, ok := value.(types.Dynamic)
		if !ok {
			continue
		}

		valueTypes = append(valueTypes, dynValue.UnderlyingValue().Type(ctx).String())
	}

	returnVal := types.StringValue(
		fmt.Sprintf("param1: %s, varparam: [%s]", stringParam.Type(ctx), strings.Join(valueTypes, ", ")),
	)

	resp.Error = resp.Result.Set(ctx, returnVal)
}
