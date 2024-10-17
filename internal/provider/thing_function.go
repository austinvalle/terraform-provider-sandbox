package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &DoThingFunction{}

type DoThingFunction struct{}

func NewDoThingFunction() function.Function {
	return &DoThingFunction{}
}

func (f *DoThingFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "do_thing"
}

func (f *DoThingFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Parameters: []function.Parameter{
			function.StringParameter{
				Name: "name",
				Validators: []function.StringParameterValidator{
					stringvalidator.LengthBetween(0, 2),
				},
			},
			function.ListParameter{
				Name:        "list",
				ElementType: types.StringType,
				Validators: []function.ListParameterValidator{
					listvalidator.UniqueValues(),
				},
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *DoThingFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	resp.Error = resp.Result.Set(ctx, true)
}
