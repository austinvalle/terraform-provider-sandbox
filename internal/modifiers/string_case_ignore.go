package modifiers

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// Yoinked from:
// https://github.com/oracle/terraform-provider-oci/blob/a50f7729707668d88e51e642178464a1985717ba/internal/service-framework/vault/vault_secret_resource.go#L277-L298
func StringEqualIgnoreCasePlanModifier() planmodifier.String {
	return &stringEqualIgnoreCasePlanModifier{}
}

type stringEqualIgnoreCasePlanModifier struct {
}

func (d *stringEqualIgnoreCasePlanModifier) Description(ctx context.Context) string {
	return "This attribute is case insensitive"
}

func (d *stringEqualIgnoreCasePlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *stringEqualIgnoreCasePlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.PlanValue.IsUnknown() && !req.PlanValue.IsNull() && !req.StateValue.IsNull() {
		if strings.EqualFold(req.PlanValue.ValueString(), req.StateValue.ValueString()) {
			resp.PlanValue = req.StateValue
		}
	}
}
