// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDxBgpPeerInvalidAddressFamilyRule checks the pattern is valid
type AwsDxBgpPeerInvalidAddressFamilyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDxBgpPeerInvalidAddressFamilyRule returns new rule with default attributes
func NewAwsDxBgpPeerInvalidAddressFamilyRule() *AwsDxBgpPeerInvalidAddressFamilyRule {
	return &AwsDxBgpPeerInvalidAddressFamilyRule{
		resourceType:  "aws_dx_bgp_peer",
		attributeName: "address_family",
		enum: []string{
			"ipv4",
			"ipv6",
		},
	}
}

// Name returns the rule name
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Name() string {
	return "aws_dx_bgp_peer_invalid_address_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDxBgpPeerInvalidAddressFamilyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as address_family`, truncateLongMessage(val)),
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
