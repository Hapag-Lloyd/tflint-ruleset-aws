// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsIAMSamlProviderInvalidSamlMetadataDocumentRule checks the pattern is valid
type AwsIAMSamlProviderInvalidSamlMetadataDocumentRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsIAMSamlProviderInvalidSamlMetadataDocumentRule returns new rule with default attributes
func NewAwsIAMSamlProviderInvalidSamlMetadataDocumentRule() *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule {
	return &AwsIAMSamlProviderInvalidSamlMetadataDocumentRule{
		resourceType:  "aws_iam_saml_provider",
		attributeName: "saml_metadata_document",
		max:           10000000,
		min:           1000,
	}
}

// Name returns the rule name
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Name() string {
	return "aws_iam_saml_provider_invalid_saml_metadata_document"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMSamlProviderInvalidSamlMetadataDocumentRule) Check(runner tflint.Runner) error {
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
					"saml_metadata_document must be 10000000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"saml_metadata_document must be 1000 characters or higher",
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
