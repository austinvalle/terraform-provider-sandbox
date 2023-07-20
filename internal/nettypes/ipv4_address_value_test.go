package nettypes_test

import (
	"net/netip"
	"testing"

	"github.com/austinvalle/terraform-provider-sandbox/internal/nettypes"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TestIPv4AddressValueIPv4Address(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ipValue        nettypes.IPv4Address
		expectedIpAddr netip.Addr
		expectedDiags  diag.Diagnostics
	}{
		"IPv4 address value is null ": {
			ipValue: nettypes.NewIPv4AddressNull(),
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"IPv4Address ValueIPv4Address Error",
					"IPv4 address string value is null",
				),
			},
		},
		"IPv4 address value is unknown ": {
			ipValue: nettypes.NewIPv4AddressUnknown(),
			expectedDiags: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"IPv4Address ValueIPv4Address Error",
					"IPv4 address string value is unknown",
				),
			},
		},
		"valid IPv4 address ": {
			ipValue:        nettypes.NewIPv4AddressValue("127.0.0.1"),
			expectedIpAddr: netip.MustParseAddr("127.0.0.1"),
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ipAddr, diags := testCase.ipValue.ValueIPv4Address()

			if ipAddr != testCase.expectedIpAddr {
				t.Errorf("Unexpected difference in netip.Addr, got: %s, expected: %s", ipAddr, testCase.expectedIpAddr)
			}

			if diff := cmp.Diff(diags, testCase.expectedDiags); diff != "" {
				t.Errorf("Unexpected diagnostics (+got, -expected): %s", diff)
			}
		})
	}
}
