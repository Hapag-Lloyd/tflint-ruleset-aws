// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmAssociationInvalidMaxErrorsRule checks the pattern is valid
type AwsSsmAssociationInvalidMaxErrorsRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidMaxErrorsRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidMaxErrorsRule() *AwsSsmAssociationInvalidMaxErrorsRule {
	return &AwsSsmAssociationInvalidMaxErrorsRule{
		resourceType:  "aws_ssm_association",
		attributeName: "max_errors",
		max:           7,
		min:           1,
		pattern:       regexp.MustCompile(`^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Name() string {
	return "aws_ssm_association_invalid_max_errors"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Check(runner tflint.Runner) error {
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
					"max_errors must be 7 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"max_errors must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`),
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
