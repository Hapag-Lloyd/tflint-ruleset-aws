// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIotPolicyInvalidPolicyRule checks the pattern is valid
type AwsIotPolicyInvalidPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsIotPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsIotPolicyInvalidPolicyRule() *AwsIotPolicyInvalidPolicyRule {
	return &AwsIotPolicyInvalidPolicyRule{
		resourceType:  "aws_iot_policy",
		attributeName: "policy",
		max:           404600,
		pattern:       regexp.MustCompile(`^[\s\S]*$`),
	}
}

// Name returns the rule name
func (r *AwsIotPolicyInvalidPolicyRule) Name() string {
	return "aws_iot_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIotPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIotPolicyInvalidPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIotPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIotPolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
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
					"policy must be 404600 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\s\S]*$`),
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
