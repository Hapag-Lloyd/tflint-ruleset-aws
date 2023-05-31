// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3BucketInvalidRequestPayerRule checks the pattern is valid
type AwsS3BucketInvalidRequestPayerRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketInvalidRequestPayerRule returns new rule with default attributes
func NewAwsS3BucketInvalidRequestPayerRule() *AwsS3BucketInvalidRequestPayerRule {
	return &AwsS3BucketInvalidRequestPayerRule{
		resourceType:  "aws_s3_bucket",
		attributeName: "request_payer",
		enum: []string{
			"Requester",
			"BucketOwner",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketInvalidRequestPayerRule) Name() string {
	return "aws_s3_bucket_invalid_request_payer"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketInvalidRequestPayerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3BucketInvalidRequestPayerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketInvalidRequestPayerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketInvalidRequestPayerRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as request_payer`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
