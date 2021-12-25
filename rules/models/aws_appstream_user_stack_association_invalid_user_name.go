// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamUserStackAssociationInvalidUserNameRule checks the pattern is valid
type AwsAppstreamUserStackAssociationInvalidUserNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAppstreamUserStackAssociationInvalidUserNameRule returns new rule with default attributes
func NewAwsAppstreamUserStackAssociationInvalidUserNameRule() *AwsAppstreamUserStackAssociationInvalidUserNameRule {
	return &AwsAppstreamUserStackAssociationInvalidUserNameRule{
		resourceType:  "aws_appstream_user_stack_association",
		attributeName: "user_name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{S}\p{N}\p{P}]+$`),
	}
}

// Name returns the rule name
func (r *AwsAppstreamUserStackAssociationInvalidUserNameRule) Name() string {
	return "aws_appstream_user_stack_association_invalid_user_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamUserStackAssociationInvalidUserNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamUserStackAssociationInvalidUserNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamUserStackAssociationInvalidUserNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamUserStackAssociationInvalidUserNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"user_name must be 128 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"user_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\p{L}\p{M}\p{S}\p{N}\p{P}]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}