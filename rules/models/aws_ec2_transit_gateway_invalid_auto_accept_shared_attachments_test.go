// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	auto_accept_shared_attachments = "true"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule(),
					Message: fmt.Sprintf(`"%s" is an invalid value as %s`, truncateLongMessage("true"), "auto_accept_shared_attachments"),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	auto_accept_shared_attachments = "enable"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
