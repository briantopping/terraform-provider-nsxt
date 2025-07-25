// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	utl "github.com/vmware/terraform-provider-nsxt/api/utl"
)

func dataSourceNsxtPolicySegment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtPolicySegmentRead,

		Schema: map[string]*schema.Schema{
			"id":           getDataSourceIDSchema(),
			"display_name": getDataSourceExtendedDisplayNameSchema(),
			"description":  getDataSourceDescriptionSchema(),
			"path":         getPathSchema(),
			"context":      getContextSchemaWithSpec(utl.SessionContextSpec{IsRequired: false, IsComputed: false, IsVpc: false, AllowDefaultProject: false, FromGlobal: true}),
		},
	}
}

func dataSourceNsxtPolicySegmentRead(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	_, err := policyDataSourceResourceRead(d, connector, getSessionContext(d, m), "Segment", nil)
	if err != nil {
		return err
	}

	return nil
}
