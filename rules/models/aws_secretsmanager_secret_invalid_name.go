// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecretsmanagerSecretInvalidNameRule checks the pattern is valid
type AwsSecretsmanagerSecretInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSecretsmanagerSecretInvalidNameRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretInvalidNameRule() *AwsSecretsmanagerSecretInvalidNameRule {
	return &AwsSecretsmanagerSecretInvalidNameRule{
		resourceType:  "aws_secretsmanager_secret",
		attributeName: "name",
		max:           512,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretInvalidNameRule) Name() string {
	return "aws_secretsmanager_secret_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
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
