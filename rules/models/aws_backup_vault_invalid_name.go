// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsBackupVaultInvalidNameRule checks the pattern is valid
type AwsBackupVaultInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsBackupVaultInvalidNameRule returns new rule with default attributes
func NewAwsBackupVaultInvalidNameRule() *AwsBackupVaultInvalidNameRule {
	return &AwsBackupVaultInvalidNameRule{
		resourceType:  "aws_backup_vault",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9\-\_]{2,50}$`),
	}
}

// Name returns the rule name
func (r *AwsBackupVaultInvalidNameRule) Name() string {
	return "aws_backup_vault_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBackupVaultInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBackupVaultInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBackupVaultInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBackupVaultInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9\-\_]{2,50}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
