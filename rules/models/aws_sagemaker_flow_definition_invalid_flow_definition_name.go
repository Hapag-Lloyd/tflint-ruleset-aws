// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule checks the pattern is valid
type AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule returns new rule with default attributes
func NewAwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule() *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule {
	return &AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule{
		resourceType:  "aws_sagemaker_flow_definition",
		attributeName: "flow_definition_name",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-z0-9](-*[a-z0-9]){0,62}`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Name() string {
	return "aws_sagemaker_flow_definition_invalid_flow_definition_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerFlowDefinitionInvalidFlowDefinitionNameRule) Check(runner tflint.Runner) error {
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
					"flow_definition_name must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"flow_definition_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9](-*[a-z0-9]){0,62}`),
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
