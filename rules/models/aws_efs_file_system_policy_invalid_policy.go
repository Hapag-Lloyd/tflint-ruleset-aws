// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEfsFileSystemPolicyInvalidPolicyRule checks the pattern is valid
type AwsEfsFileSystemPolicyInvalidPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEfsFileSystemPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsEfsFileSystemPolicyInvalidPolicyRule() *AwsEfsFileSystemPolicyInvalidPolicyRule {
	return &AwsEfsFileSystemPolicyInvalidPolicyRule{
		resourceType:  "aws_efs_file_system_policy",
		attributeName: "policy",
		max:           20000,
		min:           1,
		pattern:       regexp.MustCompile(`^[\s\S]+$`),
	}
}

// Name returns the rule name
func (r *AwsEfsFileSystemPolicyInvalidPolicyRule) Name() string {
	return "aws_efs_file_system_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEfsFileSystemPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEfsFileSystemPolicyInvalidPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEfsFileSystemPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEfsFileSystemPolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
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
					"policy must be 20000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\s\S]+$`),
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
