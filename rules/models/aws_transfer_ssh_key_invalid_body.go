// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferSSHKeyInvalidBodyRule checks the pattern is valid
type AwsTransferSSHKeyInvalidBodyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsTransferSSHKeyInvalidBodyRule returns new rule with default attributes
func NewAwsTransferSSHKeyInvalidBodyRule() *AwsTransferSSHKeyInvalidBodyRule {
	return &AwsTransferSSHKeyInvalidBodyRule{
		resourceType:  "aws_transfer_ssh_key",
		attributeName: "body",
		max:           2048,
		pattern:       regexp.MustCompile(`^\s*(ssh|ecdsa)-[a-z0-9-]+[ \t]+(([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{1,3})?(={0,3})?)(\s*|[ \t]+[\S \t]*\s*)$`),
	}
}

// Name returns the rule name
func (r *AwsTransferSSHKeyInvalidBodyRule) Name() string {
	return "aws_transfer_ssh_key_invalid_body"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferSSHKeyInvalidBodyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferSSHKeyInvalidBodyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferSSHKeyInvalidBodyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferSSHKeyInvalidBodyRule) Check(runner tflint.Runner) error {
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
					"body must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\s*(ssh|ecdsa)-[a-z0-9-]+[ \t]+(([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{1,3})?(={0,3})?)(\s*|[ \t]+[\S \t]*\s*)$`),
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
