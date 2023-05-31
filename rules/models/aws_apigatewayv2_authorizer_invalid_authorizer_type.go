// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule checks the pattern is valid
type AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule returns new rule with default attributes
func NewAwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule() *AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule {
	return &AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule{
		resourceType:  "aws_apigatewayv2_authorizer",
		attributeName: "authorizer_type",
		enum: []string{
			"REQUEST",
			"JWT",
		},
	}
}

// Name returns the rule name
func (r *AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule) Name() string {
	return "aws_apigatewayv2_authorizer_invalid_authorizer_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsApigatewayv2AuthorizerInvalidAuthorizerTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as authorizer_type`, truncateLongMessage(val)),
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
