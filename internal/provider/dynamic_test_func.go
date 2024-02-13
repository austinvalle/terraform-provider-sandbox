package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &DynamicTestFunction{}

type DynamicTestFunction struct{}

func NewDynamicTestFunction() function.Function {
	return &DynamicTestFunction{}
}

func (f *DynamicTestFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "dynamic_test"
}

func (f *DynamicTestFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Doing dynamic things...",
		Description: "With dynamic flair.",

		Parameters: []function.Parameter{
			function.DynamicParameter{
				Name:        "dynam_1",
				Description: "It could be anything!",
			},
			function.DynamicParameter{
				Name:        "dynam_2",
				Description: "It could even be a boat!",
			},
		},
		Return: function.DynamicReturn{},
	}
}

func (f *DynamicTestFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var dynam1, dynam2 types.Dynamic

	resp.Diagnostics.Append(req.Arguments.Get(ctx, &dynam1, &dynam2)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dynamObj, objDiags := types.ObjectValue(
		map[string]attr.Type{
			"dynam_1": dynam1.UnderlyingValue().Type(ctx),
			"dynam_2": dynam2.UnderlyingValue().Type(ctx),
		},
		map[string]attr.Value{
			"dynam_1": dynam1.UnderlyingValue(),
			"dynam_2": dynam2.UnderlyingValue(),
		},
	)

	resp.Diagnostics.Append(objDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dynamResp := types.DynamicValue(dynamObj)
	resp.Diagnostics.Append(resp.Result.Set(ctx, &dynamResp)...)
}
