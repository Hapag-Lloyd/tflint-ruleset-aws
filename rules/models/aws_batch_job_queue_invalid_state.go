// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsBatchJobQueueInvalidStateRule checks the pattern is valid
type AwsBatchJobQueueInvalidStateRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsBatchJobQueueInvalidStateRule returns new rule with default attributes
func NewAwsBatchJobQueueInvalidStateRule() *AwsBatchJobQueueInvalidStateRule {
	return &AwsBatchJobQueueInvalidStateRule{
		resourceType:  "aws_batch_job_queue",
		attributeName: "state",
		enum: []string{
			"ENABLED",
			"DISABLED",
		},
	}
}

// Name returns the rule name
func (r *AwsBatchJobQueueInvalidStateRule) Name() string {
	return "aws_batch_job_queue_invalid_state"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBatchJobQueueInvalidStateRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBatchJobQueueInvalidStateRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBatchJobQueueInvalidStateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBatchJobQueueInvalidStateRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as state`, truncateLongMessage(val)),
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
