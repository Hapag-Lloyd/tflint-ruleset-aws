// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamImageBuilderInvalidDisplayNameRule checks the pattern is valid
type AwsAppstreamImageBuilderInvalidDisplayNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsAppstreamImageBuilderInvalidDisplayNameRule returns new rule with default attributes
func NewAwsAppstreamImageBuilderInvalidDisplayNameRule() *AwsAppstreamImageBuilderInvalidDisplayNameRule {
	return &AwsAppstreamImageBuilderInvalidDisplayNameRule{
		resourceType:  "aws_appstream_image_builder",
		attributeName: "display_name",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsAppstreamImageBuilderInvalidDisplayNameRule) Name() string {
	return "aws_appstream_image_builder_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamImageBuilderInvalidDisplayNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamImageBuilderInvalidDisplayNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamImageBuilderInvalidDisplayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamImageBuilderInvalidDisplayNameRule) Check(runner tflint.Runner) error {
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
