// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWorklinkFleetInvalidDisplayNameRule checks the pattern is valid
type AwsWorklinkFleetInvalidDisplayNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsWorklinkFleetInvalidDisplayNameRule returns new rule with default attributes
func NewAwsWorklinkFleetInvalidDisplayNameRule() *AwsWorklinkFleetInvalidDisplayNameRule {
	return &AwsWorklinkFleetInvalidDisplayNameRule{
		resourceType:  "aws_worklink_fleet",
		attributeName: "display_name",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Name() string {
	return "aws_worklink_fleet_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorklinkFleetInvalidDisplayNameRule) Check(runner tflint.Runner) error {
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
					"display_name must be 100 characters or less",
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
