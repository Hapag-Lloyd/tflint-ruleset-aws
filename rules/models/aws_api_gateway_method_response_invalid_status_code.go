// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAPIGatewayMethodResponseInvalidStatusCodeRule checks the pattern is valid
type AwsAPIGatewayMethodResponseInvalidStatusCodeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAPIGatewayMethodResponseInvalidStatusCodeRule returns new rule with default attributes
func NewAwsAPIGatewayMethodResponseInvalidStatusCodeRule() *AwsAPIGatewayMethodResponseInvalidStatusCodeRule {
	return &AwsAPIGatewayMethodResponseInvalidStatusCodeRule{
		resourceType:  "aws_api_gateway_method_response",
		attributeName: "status_code",
		pattern:       regexp.MustCompile(`^[1-5]\d\d$`),
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayMethodResponseInvalidStatusCodeRule) Name() string {
	return "aws_api_gateway_method_response_invalid_status_code"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayMethodResponseInvalidStatusCodeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayMethodResponseInvalidStatusCodeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayMethodResponseInvalidStatusCodeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayMethodResponseInvalidStatusCodeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[1-5]\d\d$`),
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
