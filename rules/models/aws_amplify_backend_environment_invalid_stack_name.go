// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyBackendEnvironmentInvalidStackNameRule checks the pattern is valid
type AwsAmplifyBackendEnvironmentInvalidStackNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAmplifyBackendEnvironmentInvalidStackNameRule returns new rule with default attributes
func NewAwsAmplifyBackendEnvironmentInvalidStackNameRule() *AwsAmplifyBackendEnvironmentInvalidStackNameRule {
	return &AwsAmplifyBackendEnvironmentInvalidStackNameRule{
		resourceType:  "aws_amplify_backend_environment",
		attributeName: "stack_name",
		max:           255,
		min:           1,
		pattern:       regexp.MustCompile(`^(?s).+$`),
	}
}

// Name returns the rule name
func (r *AwsAmplifyBackendEnvironmentInvalidStackNameRule) Name() string {
	return "aws_amplify_backend_environment_invalid_stack_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyBackendEnvironmentInvalidStackNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyBackendEnvironmentInvalidStackNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyBackendEnvironmentInvalidStackNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyBackendEnvironmentInvalidStackNameRule) Check(runner tflint.Runner) error {
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
					"stack_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"stack_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(?s).+$`),
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
