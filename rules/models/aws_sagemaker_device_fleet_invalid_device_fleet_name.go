// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule checks the pattern is valid
type AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerDeviceFleetInvalidDeviceFleetNameRule returns new rule with default attributes
func NewAwsSagemakerDeviceFleetInvalidDeviceFleetNameRule() *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule {
	return &AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule{
		resourceType:  "aws_sagemaker_device_fleet",
		attributeName: "device_fleet_name",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Name() string {
	return "aws_sagemaker_device_fleet_invalid_device_fleet_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerDeviceFleetInvalidDeviceFleetNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"device_fleet_name must be 63 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"device_fleet_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}