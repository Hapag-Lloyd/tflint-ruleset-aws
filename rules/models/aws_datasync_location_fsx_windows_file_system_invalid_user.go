// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule checks the pattern is valid
type AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule returns new rule with default attributes
func NewAwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule() *AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule {
	return &AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule{
		resourceType:  "aws_datasync_location_fsx_windows_file_system",
		attributeName: "user",
		max:           104,
		pattern:       regexp.MustCompile(`^[^\x22\x5B\x5D/\\:;|=,+*?\x3C\x3E]{1,104}$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule) Name() string {
	return "aws_datasync_location_fsx_windows_file_system_invalid_user"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidUserRule) Check(runner tflint.Runner) error {
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
					"user must be 104 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[^\x22\x5B\x5D/\\:;|=,+*?\x3C\x3E]{1,104}$`),
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
