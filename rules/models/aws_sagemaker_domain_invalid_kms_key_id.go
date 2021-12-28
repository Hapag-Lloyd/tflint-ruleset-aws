// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerDomainInvalidKmsKeyIDRule checks the pattern is valid
type AwsSagemakerDomainInvalidKmsKeyIDRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerDomainInvalidKmsKeyIDRule returns new rule with default attributes
func NewAwsSagemakerDomainInvalidKmsKeyIDRule() *AwsSagemakerDomainInvalidKmsKeyIDRule {
	return &AwsSagemakerDomainInvalidKmsKeyIDRule{
		resourceType:  "aws_sagemaker_domain",
		attributeName: "kms_key_id",
		max:           2048,
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerDomainInvalidKmsKeyIDRule) Name() string {
	return "aws_sagemaker_domain_invalid_kms_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerDomainInvalidKmsKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerDomainInvalidKmsKeyIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerDomainInvalidKmsKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerDomainInvalidKmsKeyIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"kms_key_id must be 2048 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}