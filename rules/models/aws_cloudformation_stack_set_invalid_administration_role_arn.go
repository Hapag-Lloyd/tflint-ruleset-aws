// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudformationStackSetInvalidAdministrationRoleArnRule checks the pattern is valid
type AwsCloudformationStackSetInvalidAdministrationRoleArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackSetInvalidAdministrationRoleArnRule returns new rule with default attributes
func NewAwsCloudformationStackSetInvalidAdministrationRoleArnRule() *AwsCloudformationStackSetInvalidAdministrationRoleArnRule {
	return &AwsCloudformationStackSetInvalidAdministrationRoleArnRule{
		resourceType:  "aws_cloudformation_stack_set",
		attributeName: "administration_role_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Name() string {
	return "aws_cloudformation_stack_set_invalid_administration_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Check(runner tflint.Runner) error {
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
					"administration_role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"administration_role_arn must be 20 characters or higher",
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
