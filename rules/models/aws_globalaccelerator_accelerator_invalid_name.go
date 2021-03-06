// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlobalacceleratorAcceleratorInvalidNameRule checks the pattern is valid
type AwsGlobalacceleratorAcceleratorInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsGlobalacceleratorAcceleratorInvalidNameRule returns new rule with default attributes
func NewAwsGlobalacceleratorAcceleratorInvalidNameRule() *AwsGlobalacceleratorAcceleratorInvalidNameRule {
	return &AwsGlobalacceleratorAcceleratorInvalidNameRule{
		resourceType:  "aws_globalaccelerator_accelerator",
		attributeName: "name",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorAcceleratorInvalidNameRule) Name() string {
	return "aws_globalaccelerator_accelerator_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorAcceleratorInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorAcceleratorInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorAcceleratorInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorAcceleratorInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 255 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
