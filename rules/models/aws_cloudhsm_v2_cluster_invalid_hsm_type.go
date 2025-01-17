// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudhsmV2ClusterInvalidHsmTypeRule checks the pattern is valid
type AwsCloudhsmV2ClusterInvalidHsmTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsCloudhsmV2ClusterInvalidHsmTypeRule returns new rule with default attributes
func NewAwsCloudhsmV2ClusterInvalidHsmTypeRule() *AwsCloudhsmV2ClusterInvalidHsmTypeRule {
	return &AwsCloudhsmV2ClusterInvalidHsmTypeRule{
		resourceType:  "aws_cloudhsm_v2_cluster",
		attributeName: "hsm_type",
		max:           32,
		pattern:       regexp.MustCompile(`^((p|)hsm[0-9][a-z.]*\.[a-zA-Z]+)$`),
	}
}

// Name returns the rule name
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Name() string {
	return "aws_cloudhsm_v2_cluster_invalid_hsm_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Check(runner tflint.Runner) error {
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
					"hsm_type must be 32 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^((p|)hsm[0-9][a-z.]*\.[a-zA-Z]+)$`),
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
