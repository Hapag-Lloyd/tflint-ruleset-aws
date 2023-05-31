// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule checks the pattern is valid
type AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule returns new rule with default attributes
func NewAwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule() *AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule {
	return &AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule{
		resourceType:  "aws_servicecatalog_budget_resource_association",
		attributeName: "resource_id",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-]*`),
	}
}

// Name returns the rule name
func (r *AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule) Name() string {
	return "aws_servicecatalog_budget_resource_association_invalid_resource_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogBudgetResourceAssociationInvalidResourceIDRule) Check(runner tflint.Runner) error {
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
					"resource_id must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-]*`),
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
