// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewaySmbFileShareInvalidAuthenticationRule checks the pattern is valid
type AwsStoragegatewaySmbFileShareInvalidAuthenticationRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewaySmbFileShareInvalidAuthenticationRule returns new rule with default attributes
func NewAwsStoragegatewaySmbFileShareInvalidAuthenticationRule() *AwsStoragegatewaySmbFileShareInvalidAuthenticationRule {
	return &AwsStoragegatewaySmbFileShareInvalidAuthenticationRule{
		resourceType:  "aws_storagegateway_smb_file_share",
		attributeName: "authentication",
		max:           15,
		min:           5,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewaySmbFileShareInvalidAuthenticationRule) Name() string {
	return "aws_storagegateway_smb_file_share_invalid_authentication"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewaySmbFileShareInvalidAuthenticationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewaySmbFileShareInvalidAuthenticationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewaySmbFileShareInvalidAuthenticationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewaySmbFileShareInvalidAuthenticationRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"authentication must be 15 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"authentication must be 5 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}