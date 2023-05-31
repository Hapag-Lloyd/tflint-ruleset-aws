// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	comparison_operator = "GreaterThanOrEqual"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudwatchMetricAlarmInvalidComparisonOperatorRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("GreaterThanOrEqual"), "comparison_operator"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	comparison_operator = "GreaterThanOrEqualToThreshold"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCloudwatchMetricAlarmInvalidComparisonOperatorRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
