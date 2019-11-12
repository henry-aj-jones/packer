// Code generated by "mapstructure-to-hcl2 -type tencentCloudDataDisk"; DO NOT EDIT.
package cvm

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlattencentCloudDataDisk is an auto-generated flat version of tencentCloudDataDisk.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlattencentCloudDataDisk struct {
	DiskType   *string `mapstructure:"disk_type" cty:"disk_type"`
	DiskSize   *int64  `mapstructure:"disk_size" cty:"disk_size"`
	SnapshotId *string `mapstructure:"disk_snapshot_id" cty:"disk_snapshot_id"`
}

// FlatMapstructure returns a new FlattencentCloudDataDisk.
// FlattencentCloudDataDisk is an auto-generated flat version of tencentCloudDataDisk.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*tencentCloudDataDisk) FlatMapstructure() interface{} { return new(FlattencentCloudDataDisk) }

// HCL2Spec returns the hcl spec of a tencentCloudDataDisk.
// This spec is used by HCL to read the fields of tencentCloudDataDisk.
// The decoded values from this spec will then be applied to a FlattencentCloudDataDisk.
func (*tencentCloudDataDisk) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"disk_type":        &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"disk_size":        &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"disk_snapshot_id": &hcldec.AttrSpec{Name: "disk_snapshot_id", Type: cty.String, Required: false},
	}
	return s
}
