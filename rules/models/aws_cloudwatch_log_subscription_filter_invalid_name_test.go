// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCloudwatchLogSubscriptionFilterInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_log_subscription_filter" "foo" {
	name = "test_lambdafunction_logfilter:test"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCloudwatchLogSubscriptionFilterInvalidNameRule(),
					Message: fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage("test_lambdafunction_logfilter:test"), `^[^:*]*$`),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_log_subscription_filter" "foo" {
	name = "test_lambdafunction_logfilter"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCloudwatchLogSubscriptionFilterInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
