// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ZoneInvalidNameRule checks the pattern is valid
type AwsRoute53ZoneInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53ZoneInvalidNameRule returns new rule with default attributes
func NewAwsRoute53ZoneInvalidNameRule() *AwsRoute53ZoneInvalidNameRule {
	return &AwsRoute53ZoneInvalidNameRule{
		resourceType:  "aws_route53_zone",
		attributeName: "name",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsRoute53ZoneInvalidNameRule) Name() string {
	return "aws_route53_zone_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ZoneInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ZoneInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ZoneInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ZoneInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 1024 characters or less",
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
