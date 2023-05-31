// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCodepipelineWebhookInvalidAuthenticationRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	authentication = "GITLAB_HMAC"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCodepipelineWebhookInvalidAuthenticationRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("GITLAB_HMAC"), "authentication"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_codepipeline_webhook" "foo" {
	authentication = "GITHUB_HMAC"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCodepipelineWebhookInvalidAuthenticationRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
