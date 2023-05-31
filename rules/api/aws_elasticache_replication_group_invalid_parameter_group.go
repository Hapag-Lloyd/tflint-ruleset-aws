// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsElastiCacheReplicationGroupInvalidParameterGroupRule checks whether attribute value actually exists
type AwsElastiCacheReplicationGroupInvalidParameterGroupRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsElastiCacheReplicationGroupInvalidParameterGroupRule returns new rule with default attributes
func NewAwsElastiCacheReplicationGroupInvalidParameterGroupRule() *AwsElastiCacheReplicationGroupInvalidParameterGroupRule {
	return &AwsElastiCacheReplicationGroupInvalidParameterGroupRule{
		resourceType:  "aws_elasticache_replication_group",
		attributeName: "parameter_group_name",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsElastiCacheReplicationGroupInvalidParameterGroupRule) Name() string {
	return "aws_elasticache_replication_group_invalid_parameter_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheReplicationGroupInvalidParameterGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheReplicationGroupInvalidParameterGroupRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheReplicationGroupInvalidParameterGroupRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsElastiCacheReplicationGroupInvalidParameterGroupRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeCacheParameterGroups
func (r *AwsElastiCacheReplicationGroupInvalidParameterGroupRule) Check(rr tflint.Runner) error {
	runner := rr.(*aws.Runner)

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
			{Name: "provider"},
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

		if !r.dataPrepared {
			awsClient, err := runner.AwsClient(resource.Body.Attributes)
			if err != nil {
				return err
			}
			logger.Debug("invoking DescribeCacheParameterGroups")
			r.data, err = awsClient.DescribeCacheParameterGroups()
			if err != nil {
				err := fmt.Errorf("An error occurred while invoking DescribeCacheParameterGroups; %w", err)
				logger.Error("%s", err)
				return err
			}
			r.dataPrepared = true
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid parameter group name.`, val),
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
