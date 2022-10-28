/* Copyright © 2022 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceNsxtPolicyIPSecVpnLocalEndpoint(t *testing.T) {
	name := getAccTestResourceName()
	testResourceName := "data.nsxt_policy_ipsec_vpn_local_endpoint.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccOnlyLocalManager(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNsxtPolicyIPSecVpnLocalEndpointCheckDestroy(state, name)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNsxtPolicyIPSecVpnLocalEndpointDataSourceTemplate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "display_name", name),
					resource.TestCheckResourceAttrSet(testResourceName, "path"),
				),
			},
		},
	})
}

func testAccNsxtPolicyIPSecVpnLocalEndpointDataSourceTemplate(name string) string {
	return testAccNsxtPolicyDataSourceIPSecVpnLocalEndpointPreconfig(name) + `
data "nsxt_policy_ipsec_vpn_local_endpoint" "test" {
  display_name = nsxt_policy_ipsec_vpn_local_endpoint.test.display_name
}`
}

func testAccNsxtPolicyDataSourceIPSecVpnLocalEndpointPreconfig(name string) string {
	return testAccNsxtPolicyEdgeClusterReadTemplate(getEdgeClusterName()) +
		testAccNsxtPolicyTier0WithEdgeClusterForVPN("test") + fmt.Sprintf(`
resource "nsxt_policy_ipsec_vpn_service" "test" {
	display_name        = "%s"
	locale_service_path =  one(nsxt_policy_tier0_gateway.test.locale_service).path
}

resource "nsxt_policy_ipsec_vpn_local_endpoint" "test" {
	service_path  = nsxt_policy_ipsec_vpn_service.test.path
	display_name  = "%s"
	local_address = "%s"
  }`, name, name, "20.20.0.10")
}