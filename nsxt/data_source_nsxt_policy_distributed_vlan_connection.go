// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNsxtPolicyDistributedVlanConnection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtPolicyDistributedVlanConnectionRead,

		Schema: map[string]*schema.Schema{
			"id":           getDataSourceIDSchema(),
			"path":         getPathSchema(),
			"display_name": getOptionalDisplayNameSchema(true),
			"description":  getDescriptionSchema(),
		},
	}
}

func dataSourceNsxtPolicyDistributedVlanConnectionRead(d *schema.ResourceData, m interface{}) error {
	connector := getPolicyConnector(m)

	_, err := policyDataSourceResourceReadWithValidation(d, connector, getSessionContext(d, m), "DistributedVlanConnection", nil, false)
	if err != nil {
		return err
	}

	return nil
}
