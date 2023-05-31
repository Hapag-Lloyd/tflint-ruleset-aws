// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule checks the pattern is valid
type AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule returns new rule with default attributes
func NewAwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule() *AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule {
	return &AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule{
		resourceType:  "aws_organizations_delegated_administrator",
		attributeName: "service_principal",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]*$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule) Name() string {
	return "aws_organizations_delegated_administrator_invalid_service_principal"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsDelegatedAdministratorInvalidServicePrincipalRule) Check(runner tflint.Runner) error {
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
					"service_principal must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"service_principal must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.@-]*$`),
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
