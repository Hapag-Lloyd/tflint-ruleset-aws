// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSfnStateMachineInvalidRoleArnRule checks the pattern is valid
type AwsSfnStateMachineInvalidRoleArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSfnStateMachineInvalidRoleArnRule returns new rule with default attributes
func NewAwsSfnStateMachineInvalidRoleArnRule() *AwsSfnStateMachineInvalidRoleArnRule {
	return &AwsSfnStateMachineInvalidRoleArnRule{
		resourceType:  "aws_sfn_state_machine",
		attributeName: "role_arn",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSfnStateMachineInvalidRoleArnRule) Name() string {
	return "aws_sfn_state_machine_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSfnStateMachineInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSfnStateMachineInvalidRoleArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSfnStateMachineInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSfnStateMachineInvalidRoleArnRule) Check(runner tflint.Runner) error {
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
					"role_arn must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 1 characters or higher",
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
