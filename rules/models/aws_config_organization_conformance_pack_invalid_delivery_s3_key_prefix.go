// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule checks the pattern is valid
type AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule returns new rule with default attributes
func NewAwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule() *AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule {
	return &AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule{
		resourceType:  "aws_config_organization_conformance_pack",
		attributeName: "delivery_s3_key_prefix",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule) Name() string {
	return "aws_config_organization_conformance_pack_invalid_delivery_s3_key_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationConformancePackInvalidDeliveryS3KeyPrefixRule) Check(runner tflint.Runner) error {
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
					"delivery_s3_key_prefix must be 1024 characters or less",
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
