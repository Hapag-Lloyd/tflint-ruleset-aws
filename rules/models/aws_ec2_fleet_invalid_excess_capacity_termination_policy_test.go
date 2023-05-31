// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEc2FleetInvalidExcessCapacityTerminationPolicyRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_fleet" "foo" {
	excess_capacity_termination_policy = "remain"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEc2FleetInvalidExcessCapacityTerminationPolicyRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("remain"), "excess_capacity_termination_policy"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_fleet" "foo" {
	excess_capacity_termination_policy = "termination"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEc2FleetInvalidExcessCapacityTerminationPolicyRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
