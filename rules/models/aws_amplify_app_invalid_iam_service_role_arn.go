// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyAppInvalidIAMServiceRoleArnRule checks the pattern is valid
type AwsAmplifyAppInvalidIAMServiceRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAmplifyAppInvalidIAMServiceRoleArnRule returns new rule with default attributes
func NewAwsAmplifyAppInvalidIAMServiceRoleArnRule() *AwsAmplifyAppInvalidIAMServiceRoleArnRule {
	return &AwsAmplifyAppInvalidIAMServiceRoleArnRule{
		resourceType:  "aws_amplify_app",
		attributeName: "iam_service_role_arn",
		max:           1000,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Name() string {
	return "aws_amplify_app_invalid_iam_service_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyAppInvalidIAMServiceRoleArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"iam_service_role_arn must be 1000 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"iam_service_role_arn must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}