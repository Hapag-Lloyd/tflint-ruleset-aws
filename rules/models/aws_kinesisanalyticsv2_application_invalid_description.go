// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule checks the pattern is valid
type AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsKinesisanalyticsv2ApplicationInvalidDescriptionRule returns new rule with default attributes
func NewAwsKinesisanalyticsv2ApplicationInvalidDescriptionRule() *AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule {
	return &AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule{
		resourceType:  "aws_kinesisanalyticsv2_application",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule) Name() string {
	return "aws_kinesisanalyticsv2_application_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKinesisanalyticsv2ApplicationInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 1024 characters or less",
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
