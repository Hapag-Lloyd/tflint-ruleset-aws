// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEcsServiceInvalidLaunchTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ecs_service" "foo" {
	launch_type = "POD"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEcsServiceInvalidLaunchTypeRule(),
					Message: `"POD" is an invalid value as launch_type`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ecs_service" "foo" {
	launch_type = "FARGATE"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEcsServiceInvalidLaunchTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}