// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWafregionalRateBasedRuleInvalidRateKeyRule checks the pattern is valid
type AwsWafregionalRateBasedRuleInvalidRateKeyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsWafregionalRateBasedRuleInvalidRateKeyRule returns new rule with default attributes
func NewAwsWafregionalRateBasedRuleInvalidRateKeyRule() *AwsWafregionalRateBasedRuleInvalidRateKeyRule {
	return &AwsWafregionalRateBasedRuleInvalidRateKeyRule{
		resourceType:  "aws_wafregional_rate_based_rule",
		attributeName: "rate_key",
		enum: []string{
			"IP",
		},
	}
}

// Name returns the rule name
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Name() string {
	return "aws_wafregional_rate_based_rule_invalid_rate_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafregionalRateBasedRuleInvalidRateKeyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as rate_key`, truncateLongMessage(val)),
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
