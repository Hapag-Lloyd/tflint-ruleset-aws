// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3ObjectCopyInvalidSourceRule checks the pattern is valid
type AwsS3ObjectCopyInvalidSourceRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsS3ObjectCopyInvalidSourceRule returns new rule with default attributes
func NewAwsS3ObjectCopyInvalidSourceRule() *AwsS3ObjectCopyInvalidSourceRule {
	return &AwsS3ObjectCopyInvalidSourceRule{
		resourceType:  "aws_s3_object_copy",
		attributeName: "source",
		pattern:       regexp.MustCompile(`^\/.+\/.+$`),
	}
}

// Name returns the rule name
func (r *AwsS3ObjectCopyInvalidSourceRule) Name() string {
	return "aws_s3_object_copy_invalid_source"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3ObjectCopyInvalidSourceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3ObjectCopyInvalidSourceRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3ObjectCopyInvalidSourceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3ObjectCopyInvalidSourceRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\/.+\/.+$`),
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
