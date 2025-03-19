package modifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func UseStateNoMatterWhatForUnknown() planmodifier.String {
	return &useStateNoMatterWhatForUnknown{}
}

type useStateNoMatterWhatForUnknown struct {
}

func (d *useStateNoMatterWhatForUnknown) Description(ctx context.Context) string {
	return ""
}

func (d *useStateNoMatterWhatForUnknown) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *useStateNoMatterWhatForUnknown) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}
