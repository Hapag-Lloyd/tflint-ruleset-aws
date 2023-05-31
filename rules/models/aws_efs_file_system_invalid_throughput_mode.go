// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEfsFileSystemInvalidThroughputModeRule checks the pattern is valid
type AwsEfsFileSystemInvalidThroughputModeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEfsFileSystemInvalidThroughputModeRule returns new rule with default attributes
func NewAwsEfsFileSystemInvalidThroughputModeRule() *AwsEfsFileSystemInvalidThroughputModeRule {
	return &AwsEfsFileSystemInvalidThroughputModeRule{
		resourceType:  "aws_efs_file_system",
		attributeName: "throughput_mode",
		enum: []string{
			"bursting",
			"provisioned",
			"elastic",
		},
	}
}

// Name returns the rule name
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Name() string {
	return "aws_efs_file_system_invalid_throughput_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as throughput_mode`, truncateLongMessage(val)),
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
