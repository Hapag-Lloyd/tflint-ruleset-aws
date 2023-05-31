// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMRolePolicyInvalidRoleRule checks the pattern is valid
type AwsIAMRolePolicyInvalidRoleRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMRolePolicyInvalidRoleRule returns new rule with default attributes
func NewAwsIAMRolePolicyInvalidRoleRule() *AwsIAMRolePolicyInvalidRoleRule {
	return &AwsIAMRolePolicyInvalidRoleRule{
		resourceType:  "aws_iam_role_policy",
		attributeName: "role",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMRolePolicyInvalidRoleRule) Name() string {
	return "aws_iam_role_policy_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMRolePolicyInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMRolePolicyInvalidRoleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMRolePolicyInvalidRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMRolePolicyInvalidRoleRule) Check(runner tflint.Runner) error {
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
					"role must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.@-]+$`),
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
