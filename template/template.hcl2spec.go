// Code generated by "mapstructure-to-hcl2 -type Provisioner"; DO NOT EDIT.
package template

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatProvisioner is an auto-generated flat version of Provisioner.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatProvisioner struct {
	Only        []string               `json:"only,omitempty" cty:"only"`
	Except      []string               `json:"except,omitempty" cty:"except"`
	Type        *string                `json:"type" cty:"type"`
	Config      map[string]interface{} `json:"config,omitempty" cty:"config"`
	Override    map[string]interface{} `json:"override,omitempty" cty:"override"`
	PauseBefore *string                `mapstructure:"pause_before" json:"pause_before,omitempty" cty:"pause_before"`
	Timeout     *string                `mapstructure:"timeout" json:"timeout,omitempty" cty:"timeout"`
}

// FlatMapstructure returns a new FlatProvisioner.
// FlatProvisioner is an auto-generated flat version of Provisioner.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Provisioner) FlatMapstructure() interface{} { return new(FlatProvisioner) }

// HCL2Spec returns the hcl spec of a Provisioner.
// This spec is used by HCL to read the fields of Provisioner.
// The decoded values from this spec will then be applied to a FlatProvisioner.
func (*Provisioner) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"only":         &hcldec.AttrSpec{Name: "only", Type: cty.List(cty.String), Required: false},
		"except":       &hcldec.AttrSpec{Name: "except", Type: cty.List(cty.String), Required: false},
		"type":         &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
		"config":       &hcldec.BlockAttrsSpec{TypeName: "config", ElementType: cty.String, Required: false},
		"override":     &hcldec.BlockAttrsSpec{TypeName: "override", ElementType: cty.String, Required: false},
		"pause_before": &hcldec.AttrSpec{Name: "pause_before", Type: cty.String, Required: false},
		"timeout":      &hcldec.AttrSpec{Name: "timeout", Type: cty.String, Required: false},
	}
	return s
}
