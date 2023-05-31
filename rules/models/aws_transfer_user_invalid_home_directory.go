// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferUserInvalidHomeDirectoryRule checks the pattern is valid
type AwsTransferUserInvalidHomeDirectoryRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsTransferUserInvalidHomeDirectoryRule returns new rule with default attributes
func NewAwsTransferUserInvalidHomeDirectoryRule() *AwsTransferUserInvalidHomeDirectoryRule {
	return &AwsTransferUserInvalidHomeDirectoryRule{
		resourceType:  "aws_transfer_user",
		attributeName: "home_directory",
		max:           1024,
		pattern:       regexp.MustCompile(`^$|/.*`),
	}
}

// Name returns the rule name
func (r *AwsTransferUserInvalidHomeDirectoryRule) Name() string {
	return "aws_transfer_user_invalid_home_directory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferUserInvalidHomeDirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferUserInvalidHomeDirectoryRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferUserInvalidHomeDirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferUserInvalidHomeDirectoryRule) Check(runner tflint.Runner) error {
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
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"home_directory must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^$|/.*`),
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
