package jsontypes_test

import (
	"testing"

	"github.com/austinvalle/terraform-provider-sandbox/internal/jsontypes"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TestNormalizedUnmarshal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		json          jsontypes.Normalized
		target        any
		output        any
		expectedDiags diag.Diagnostics
	}{
		"normalized value is null ": {
			json: jsontypes.NewNormalizedNull(),
			target: struct {
				Hello string `json:"hello"`
			}{},
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Normalized JSON Unmarshal Error",
					"json string value is null",
				),
			},
		},
		"normalized value is unknown ": {
			json: jsontypes.NewNormalizedUnknown(),
			target: struct {
				Hello string `json:"hello"`
			}{},
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Normalized JSON Unmarshal Error",
					"json string value is unknown",
				),
			},
		},
		"invalid target - not a pointer ": {
			json: jsontypes.NewNormalizedValue(`{"hello": "world"}`),
			target: struct {
				Hello string `json:"hello"`
			}{},
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Normalized JSON Unmarshal Error",
					"json: Unmarshal(non-pointer struct { Hello string \"json:\\\"hello\\\"\" })",
				),
			},
		},
		"valid target ": {
			json: jsontypes.NewNormalizedValue(`{"hello": "world", "nums": [1, 2, 3], "test-bool": true}`),
			target: &struct {
				Hello   string `json:"hello"`
				Numbers []int  `json:"nums"`
				Test    bool   `json:"test-bool"`
			}{},
			output: &struct {
				Hello   string `json:"hello"`
				Numbers []int  `json:"nums"`
				Test    bool   `json:"test-bool"`
			}{
				Hello:   "world",
				Numbers: []int{1, 2, 3},
				Test:    true,
			},
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			diags := testCase.json.Unmarshal(testCase.target)

			if diff := cmp.Diff(diags, testCase.expectedDiags); diff != "" {
				t.Errorf("Unexpected diagnostics (+got, -expected): %s", diff)
			}
		})
	}
}
