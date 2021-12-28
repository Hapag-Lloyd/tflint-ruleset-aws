// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsChimeVoiceConnectorGroupInvalidNameRule checks the pattern is valid
type AwsChimeVoiceConnectorGroupInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsChimeVoiceConnectorGroupInvalidNameRule returns new rule with default attributes
func NewAwsChimeVoiceConnectorGroupInvalidNameRule() *AwsChimeVoiceConnectorGroupInvalidNameRule {
	return &AwsChimeVoiceConnectorGroupInvalidNameRule{
		resourceType:  "aws_chime_voice_connector_group",
		attributeName: "name",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsChimeVoiceConnectorGroupInvalidNameRule) Name() string {
	return "aws_chime_voice_connector_group_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsChimeVoiceConnectorGroupInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsChimeVoiceConnectorGroupInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsChimeVoiceConnectorGroupInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsChimeVoiceConnectorGroupInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 256 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}