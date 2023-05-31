// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecurityhubInsightInvalidNameRule checks the pattern is valid
type AwsSecurityhubInsightInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSecurityhubInsightInvalidNameRule returns new rule with default attributes
func NewAwsSecurityhubInsightInvalidNameRule() *AwsSecurityhubInsightInvalidNameRule {
	return &AwsSecurityhubInsightInvalidNameRule{
		resourceType:  "aws_securityhub_insight",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsSecurityhubInsightInvalidNameRule) Name() string {
	return "aws_securityhub_insight_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecurityhubInsightInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecurityhubInsightInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecurityhubInsightInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecurityhubInsightInvalidNameRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
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
