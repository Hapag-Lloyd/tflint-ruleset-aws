// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudhsmV2HsmInvalidAvailabilityZoneRule checks the pattern is valid
type AwsCloudhsmV2HsmInvalidAvailabilityZoneRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudhsmV2HsmInvalidAvailabilityZoneRule returns new rule with default attributes
func NewAwsCloudhsmV2HsmInvalidAvailabilityZoneRule() *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule {
	return &AwsCloudhsmV2HsmInvalidAvailabilityZoneRule{
		resourceType:  "aws_cloudhsm_v2_hsm",
		attributeName: "availability_zone",
		pattern:       regexp.MustCompile(`^[a-z]{2}(-(gov))?-(east|west|north|south|central){1,2}-\d[a-z]$`),
	}
}

// Name returns the rule name
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Name() string {
	return "aws_cloudhsm_v2_hsm_invalid_availability_zone"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudhsmV2HsmInvalidAvailabilityZoneRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z]{2}(-(gov))?-(east|west|north|south|central){1,2}-\d[a-z]$`),
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
