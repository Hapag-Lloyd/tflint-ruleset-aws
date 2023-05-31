// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudformationStackSetInstanceInvalidAccountIDRule checks the pattern is valid
type AwsCloudformationStackSetInstanceInvalidAccountIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudformationStackSetInstanceInvalidAccountIDRule returns new rule with default attributes
func NewAwsCloudformationStackSetInstanceInvalidAccountIDRule() *AwsCloudformationStackSetInstanceInvalidAccountIDRule {
	return &AwsCloudformationStackSetInstanceInvalidAccountIDRule{
		resourceType:  "aws_cloudformation_stack_set_instance",
		attributeName: "account_id",
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Name() string {
	return "aws_cloudformation_stack_set_instance_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInstanceInvalidAccountIDRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]{12}$`),
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
