// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOntapVolumeInvalidJunctionPathRule checks the pattern is valid
type AwsFsxOntapVolumeInvalidJunctionPathRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxOntapVolumeInvalidJunctionPathRule returns new rule with default attributes
func NewAwsFsxOntapVolumeInvalidJunctionPathRule() *AwsFsxOntapVolumeInvalidJunctionPathRule {
	return &AwsFsxOntapVolumeInvalidJunctionPathRule{
		resourceType:  "aws_fsx_ontap_volume",
		attributeName: "junction_path",
		max:           255,
		min:           1,
		pattern:       regexp.MustCompile(`^[^\x{0000}\x{0085}\x{2028}\x{2029}\r\n]{1,255}$`),
	}
}

// Name returns the rule name
func (r *AwsFsxOntapVolumeInvalidJunctionPathRule) Name() string {
	return "aws_fsx_ontap_volume_invalid_junction_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOntapVolumeInvalidJunctionPathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOntapVolumeInvalidJunctionPathRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOntapVolumeInvalidJunctionPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOntapVolumeInvalidJunctionPathRule) Check(runner tflint.Runner) error {
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
					"junction_path must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"junction_path must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\x{0000}\x{0085}\x{2028}\x{2029}\r\n]{1,255}$`),
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
