// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsDatasyncLocationS3InvalidS3BucketArnRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_datasync_location_s3" "foo" {
	s3_bucket_arn = "arn:aws:eks:us-east-1:123456789012:cluster/my-cluster"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsDatasyncLocationS3InvalidS3BucketArnRule(),
					Message: fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage("arn:aws:eks:us-east-1:123456789012:cluster/my-cluster"), `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):s3:[a-z\-0-9]*:[0-9]{12}:accesspoint[/:][a-zA-Z0-9\-.]{1,63}$|^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):s3-outposts:[a-z\-0-9]+:[0-9]{12}:outpost[/:][a-zA-Z0-9\-]{1,63}[/:]accesspoint[/:][a-zA-Z0-9\-]{1,63}$|^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):s3:::[a-zA-Z0-9.\-_]{1,255}$`),
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_datasync_location_s3" "foo" {
	s3_bucket_arn = "arn:aws:s3:::my_corporate_bucket"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsDatasyncLocationS3InvalidS3BucketArnRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
