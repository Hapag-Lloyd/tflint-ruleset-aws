// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferServerInvalidURLRule checks the pattern is valid
type AwsTransferServerInvalidURLRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsTransferServerInvalidURLRule returns new rule with default attributes
func NewAwsTransferServerInvalidURLRule() *AwsTransferServerInvalidURLRule {
	return &AwsTransferServerInvalidURLRule{
		resourceType:  "aws_transfer_server",
		attributeName: "url",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidURLRule) Name() string {
	return "aws_transfer_server_invalid_url"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidURLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidURLRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidURLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidURLRule) Check(runner tflint.Runner) error {
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
					"url must be 255 characters or less",
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
