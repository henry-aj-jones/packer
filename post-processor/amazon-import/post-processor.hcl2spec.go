// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package amazonimport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/builder/amazon/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName       *string                       `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType     *string                       `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug           *bool                         `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce           *bool                         `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError         *string                       `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars        map[string]string             `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars   []string                      `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	AccessKey             *string                       `mapstructure:"access_key" required:"true" cty:"access_key"`
	CustomEndpointEc2     *string                       `mapstructure:"custom_endpoint_ec2" required:"false" cty:"custom_endpoint_ec2"`
	DecodeAuthZMessages   *bool                         `mapstructure:"decode_authorization_messages" required:"false" cty:"decode_authorization_messages"`
	InsecureSkipTLSVerify *bool                         `mapstructure:"insecure_skip_tls_verify" required:"false" cty:"insecure_skip_tls_verify"`
	MFACode               *string                       `mapstructure:"mfa_code" required:"false" cty:"mfa_code"`
	ProfileName           *string                       `mapstructure:"profile" required:"false" cty:"profile"`
	RawRegion             *string                       `mapstructure:"region" required:"true" cty:"region"`
	SecretKey             *string                       `mapstructure:"secret_key" required:"true" cty:"secret_key"`
	SkipValidation        *bool                         `mapstructure:"skip_region_validation" required:"false" cty:"skip_region_validation"`
	SkipMetadataApiCheck  *bool                         `mapstructure:"skip_metadata_api_check" cty:"skip_metadata_api_check"`
	Token                 *string                       `mapstructure:"token" required:"false" cty:"token"`
	VaultAWSEngine        *common.VaultAWSEngineOptions `mapstructure:"vault_aws_engine" required:"false" cty:"vault_aws_engine"`
	S3Bucket              *string                       `mapstructure:"s3_bucket_name" cty:"s3_bucket_name"`
	S3Key                 *string                       `mapstructure:"s3_key_name" cty:"s3_key_name"`
	S3Encryption          *string                       `mapstructure:"s3_encryption" cty:"s3_encryption"`
	S3EncryptionKey       *string                       `mapstructure:"s3_encryption_key" cty:"s3_encryption_key"`
	SkipClean             *bool                         `mapstructure:"skip_clean" cty:"skip_clean"`
	Tags                  map[string]string             `mapstructure:"tags" cty:"tags"`
	Name                  *string                       `mapstructure:"ami_name" cty:"ami_name"`
	Description           *string                       `mapstructure:"ami_description" cty:"ami_description"`
	Users                 []string                      `mapstructure:"ami_users" cty:"ami_users"`
	Groups                []string                      `mapstructure:"ami_groups" cty:"ami_groups"`
	Encrypt               *bool                         `mapstructure:"ami_encrypt" cty:"ami_encrypt"`
	KMSKey                *string                       `mapstructure:"ami_kms_key" cty:"ami_kms_key"`
	LicenseType           *string                       `mapstructure:"license_type" cty:"license_type"`
	RoleName              *string                       `mapstructure:"role_name" cty:"role_name"`
	Format                *string                       `mapstructure:"format" cty:"format"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":             &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":           &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                  &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                  &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":               &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":         &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":    &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"access_key":                    &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"custom_endpoint_ec2":           &hcldec.AttrSpec{Name: "custom_endpoint_ec2", Type: cty.String, Required: false},
		"decode_authorization_messages": &hcldec.AttrSpec{Name: "decode_authorization_messages", Type: cty.Bool, Required: false},
		"insecure_skip_tls_verify":      &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"mfa_code":                      &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"profile":                       &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"region":                        &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"secret_key":                    &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"skip_region_validation":        &hcldec.AttrSpec{Name: "skip_region_validation", Type: cty.Bool, Required: false},
		"skip_metadata_api_check":       &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"token":                         &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"vault_aws_engine":              &hcldec.BlockSpec{TypeName: "vault_aws_engine", Nested: hcldec.ObjectSpec((*common.VaultAWSEngineOptions)(nil).HCL2Spec())},
		"s3_bucket_name":                &hcldec.AttrSpec{Name: "s3_bucket_name", Type: cty.String, Required: false},
		"s3_key_name":                   &hcldec.AttrSpec{Name: "s3_key_name", Type: cty.String, Required: false},
		"s3_encryption":                 &hcldec.AttrSpec{Name: "s3_encryption", Type: cty.String, Required: false},
		"s3_encryption_key":             &hcldec.AttrSpec{Name: "s3_encryption_key", Type: cty.String, Required: false},
		"skip_clean":                    &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
		"tags":                          &hcldec.BlockAttrsSpec{TypeName: "tags", ElementType: cty.String, Required: false},
		"ami_name":                      &hcldec.AttrSpec{Name: "ami_name", Type: cty.String, Required: false},
		"ami_description":               &hcldec.AttrSpec{Name: "ami_description", Type: cty.String, Required: false},
		"ami_users":                     &hcldec.AttrSpec{Name: "ami_users", Type: cty.List(cty.String), Required: false},
		"ami_groups":                    &hcldec.AttrSpec{Name: "ami_groups", Type: cty.List(cty.String), Required: false},
		"ami_encrypt":                   &hcldec.AttrSpec{Name: "ami_encrypt", Type: cty.Bool, Required: false},
		"ami_kms_key":                   &hcldec.AttrSpec{Name: "ami_kms_key", Type: cty.String, Required: false},
		"license_type":                  &hcldec.AttrSpec{Name: "license_type", Type: cty.String, Required: false},
		"role_name":                     &hcldec.AttrSpec{Name: "role_name", Type: cty.String, Required: false},
		"format":                        &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
	}
	return s
}
