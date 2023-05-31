// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3ObjectCopyInvalidServerSideEncryptionRule checks the pattern is valid
type AwsS3ObjectCopyInvalidServerSideEncryptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3ObjectCopyInvalidServerSideEncryptionRule returns new rule with default attributes
func NewAwsS3ObjectCopyInvalidServerSideEncryptionRule() *AwsS3ObjectCopyInvalidServerSideEncryptionRule {
	return &AwsS3ObjectCopyInvalidServerSideEncryptionRule{
		resourceType:  "aws_s3_object_copy",
		attributeName: "server_side_encryption",
		enum: []string{
			"AES256",
			"aws:kms",
		},
	}
}

// Name returns the rule name
func (r *AwsS3ObjectCopyInvalidServerSideEncryptionRule) Name() string {
	return "aws_s3_object_copy_invalid_server_side_encryption"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3ObjectCopyInvalidServerSideEncryptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3ObjectCopyInvalidServerSideEncryptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3ObjectCopyInvalidServerSideEncryptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3ObjectCopyInvalidServerSideEncryptionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as server_side_encryption`, truncateLongMessage(val)),
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
