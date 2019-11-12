// Code generated by "mapstructure-to-hcl2 -type Config,ModuleDir"; DO NOT EDIT.
package converge

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName      *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType    *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug          *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce          *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError        *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars       map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars  []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Bootstrap            *bool             `mapstructure:"bootstrap" cty:"bootstrap"`
	Version              *string           `mapstructure:"version" cty:"version"`
	BootstrapCommand     *string           `mapstructure:"bootstrap_command" cty:"bootstrap_command"`
	PreventBootstrapSudo *bool             `mapstructure:"prevent_bootstrap_sudo" cty:"prevent_bootstrap_sudo"`
	ModuleDirs           []ModuleDir       `mapstructure:"module_dirs" cty:"module_dirs"`
	Module               *string           `mapstructure:"module" cty:"module"`
	WorkingDirectory     *string           `mapstructure:"working_directory" cty:"working_directory"`
	Params               map[string]string `mapstructure:"params" cty:"params"`
	ExecuteCommand       *string           `mapstructure:"execute_command" cty:"execute_command"`
	PreventSudo          *bool             `mapstructure:"prevent_sudo" cty:"prevent_sudo"`
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
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"bootstrap":                  &hcldec.AttrSpec{Name: "bootstrap", Type: cty.Bool, Required: false},
		"version":                    &hcldec.AttrSpec{Name: "version", Type: cty.String, Required: false},
		"bootstrap_command":          &hcldec.AttrSpec{Name: "bootstrap_command", Type: cty.String, Required: false},
		"prevent_bootstrap_sudo":     &hcldec.AttrSpec{Name: "prevent_bootstrap_sudo", Type: cty.Bool, Required: false},
		"module_dirs":                &hcldec.BlockListSpec{TypeName: "module_dirs", Nested: &hcldec.BlockSpec{TypeName: "module_dirs", Nested: hcldec.ObjectSpec((*ModuleDir)(nil).HCL2Spec())}},
		"module":                     &hcldec.AttrSpec{Name: "module", Type: cty.String, Required: false},
		"working_directory":          &hcldec.AttrSpec{Name: "working_directory", Type: cty.String, Required: false},
		"params":                     &hcldec.BlockAttrsSpec{TypeName: "params", ElementType: cty.String, Required: false},
		"execute_command":            &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"prevent_sudo":               &hcldec.AttrSpec{Name: "prevent_sudo", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatModuleDir is an auto-generated flat version of ModuleDir.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatModuleDir struct {
	Source      *string  `mapstructure:"source" cty:"source"`
	Destination *string  `mapstructure:"destination" cty:"destination"`
	Exclude     []string `mapstructure:"exclude" cty:"exclude"`
}

// FlatMapstructure returns a new FlatModuleDir.
// FlatModuleDir is an auto-generated flat version of ModuleDir.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ModuleDir) FlatMapstructure() interface{} { return new(FlatModuleDir) }

// HCL2Spec returns the hcl spec of a ModuleDir.
// This spec is used by HCL to read the fields of ModuleDir.
// The decoded values from this spec will then be applied to a FlatModuleDir.
func (*ModuleDir) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"source":      &hcldec.AttrSpec{Name: "source", Type: cty.String, Required: false},
		"destination": &hcldec.AttrSpec{Name: "destination", Type: cty.String, Required: false},
		"exclude":     &hcldec.AttrSpec{Name: "exclude", Type: cty.List(cty.String), Required: false},
	}
	return s
}
