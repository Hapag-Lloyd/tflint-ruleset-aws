// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferUserInvalidServerIDRule checks the pattern is valid
type AwsTransferUserInvalidServerIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsTransferUserInvalidServerIDRule returns new rule with default attributes
func NewAwsTransferUserInvalidServerIDRule() *AwsTransferUserInvalidServerIDRule {
	return &AwsTransferUserInvalidServerIDRule{
		resourceType:  "aws_transfer_user",
		attributeName: "server_id",
		max:           19,
		min:           19,
		pattern:       regexp.MustCompile(`^s-([0-9a-f]{17})$`),
	}
}

// Name returns the rule name
func (r *AwsTransferUserInvalidServerIDRule) Name() string {
	return "aws_transfer_user_invalid_server_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferUserInvalidServerIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferUserInvalidServerIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferUserInvalidServerIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferUserInvalidServerIDRule) Check(runner tflint.Runner) error {
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
					"server_id must be 19 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"server_id must be 19 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^s-([0-9a-f]{17})$`),
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
