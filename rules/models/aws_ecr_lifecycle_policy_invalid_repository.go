// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcrLifecyclePolicyInvalidRepositoryRule checks the pattern is valid
type AwsEcrLifecyclePolicyInvalidRepositoryRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEcrLifecyclePolicyInvalidRepositoryRule returns new rule with default attributes
func NewAwsEcrLifecyclePolicyInvalidRepositoryRule() *AwsEcrLifecyclePolicyInvalidRepositoryRule {
	return &AwsEcrLifecyclePolicyInvalidRepositoryRule{
		resourceType:  "aws_ecr_lifecycle_policy",
		attributeName: "repository",
		max:           256,
		min:           2,
		pattern:       regexp.MustCompile(`^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Name() string {
	return "aws_ecr_lifecycle_policy_invalid_repository"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"repository must be 256 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"repository must be 2 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}