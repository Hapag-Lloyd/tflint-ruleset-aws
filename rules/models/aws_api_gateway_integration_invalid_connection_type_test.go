// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsAPIGatewayIntegrationInvalidConnectionTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_api_gateway_integration" "foo" {
	connection_type = "INTRANET"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsAPIGatewayIntegrationInvalidConnectionTypeRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("INTRANET"), "connection_type"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_api_gateway_integration" "foo" {
	connection_type = "INTERNET"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsAPIGatewayIntegrationInvalidConnectionTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
