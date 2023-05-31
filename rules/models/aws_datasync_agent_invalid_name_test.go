// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDatasyncAgentInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_datasync_agent" "foo" {
	name = "example^example"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDatasyncAgentInvalidNameRule(),
					Message: fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage("example^example"), `^[a-zA-Z0-9\s+=._:@/-]+$`),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_datasync_agent" "foo" {
	name = "example"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsDatasyncAgentInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
