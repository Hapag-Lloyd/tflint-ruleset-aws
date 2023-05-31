// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule checks the pattern is valid
type AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule returns new rule with default attributes
func NewAwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule() *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule {
	return &AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule{
		resourceType:  "aws_config_organization_managed_rule",
		attributeName: "tag_key_scope",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Name() string {
	return "aws_config_organization_managed_rule_invalid_tag_key_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule) Check(runner tflint.Runner) error {
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
					"tag_key_scope must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"tag_key_scope must be 1 characters or higher",
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
