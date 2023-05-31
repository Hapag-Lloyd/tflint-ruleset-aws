// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderComponentInvalidDataRule checks the pattern is valid
type AwsImagebuilderComponentInvalidDataRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsImagebuilderComponentInvalidDataRule returns new rule with default attributes
func NewAwsImagebuilderComponentInvalidDataRule() *AwsImagebuilderComponentInvalidDataRule {
	return &AwsImagebuilderComponentInvalidDataRule{
		resourceType:  "aws_imagebuilder_component",
		attributeName: "data",
		max:           16000,
		min:           1,
		pattern:       regexp.MustCompile(`^[^\x00]+$`),
	}
}

// Name returns the rule name
func (r *AwsImagebuilderComponentInvalidDataRule) Name() string {
	return "aws_imagebuilder_component_invalid_data"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderComponentInvalidDataRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderComponentInvalidDataRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderComponentInvalidDataRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderComponentInvalidDataRule) Check(runner tflint.Runner) error {
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
					"data must be 16000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"data must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\x00]+$`),
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
