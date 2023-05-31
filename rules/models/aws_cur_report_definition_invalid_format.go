// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCurReportDefinitionInvalidFormatRule checks the pattern is valid
type AwsCurReportDefinitionInvalidFormatRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCurReportDefinitionInvalidFormatRule returns new rule with default attributes
func NewAwsCurReportDefinitionInvalidFormatRule() *AwsCurReportDefinitionInvalidFormatRule {
	return &AwsCurReportDefinitionInvalidFormatRule{
		resourceType:  "aws_cur_report_definition",
		attributeName: "format",
		enum: []string{
			"textORcsv",
			"Parquet",
		},
	}
}

// Name returns the rule name
func (r *AwsCurReportDefinitionInvalidFormatRule) Name() string {
	return "aws_cur_report_definition_invalid_format"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCurReportDefinitionInvalidFormatRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCurReportDefinitionInvalidFormatRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCurReportDefinitionInvalidFormatRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCurReportDefinitionInvalidFormatRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as format`, truncateLongMessage(val)),
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
