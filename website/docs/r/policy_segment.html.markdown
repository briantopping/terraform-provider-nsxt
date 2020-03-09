---
layout: "nsxt"
page_title: "NSXT: nsxt_policy_segment"
sidebar_current: "docs-nsxt-resource-policy-segment"
description: A resource to configure a network Segment.
---

# nsxt_policy_segment

This resource provides a method for the management of Segments.
 
## Example Usage

```hcl
resource "nsxt_policy_segment" "segment1" {
    display_name      = "segment1"
    description       = "Terraform provisioned Segment"
    connectivity_path = nsxt_policy_tier1_gateway.mygateway.path

    subnet {
      cidr        = "12.12.2.1/24"
      dhcp_ranges = ["12.12.2.100-12.12.2.160"]

      dhcp_v4_config {
        server_address = "12.12.2.2/24"
        lease_time     = 36000

        dhcp_option_121 {
          network  = "6.6.6.0/24"
          next_hop = "1.1.1.21"
        }

        dhcp_generic_option {
          code = "119"
          values = ["abc"]
        }
      }
    }

    advanced_config {
      connectivity = "OFF"
      local_egress = true
    }
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name of the resource.
* `description` - (Optional) Description of the resource.
* `tag` - (Optional) A list of scope + tag pairs to associate with this policy.
* `nsx_id` - (Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.
* `connectivity_path` - (Optional) Policy path to the connecting Tier-0 or Tier-1.
* `domain_name`- (Optional) DNS domain names.
* `overlay_id` - (Optional) Overlay connectivity ID for this Segment.
* `transport_zone_path` - (Optional) Policy path to the Overlay transport zone. This property is required if more than one overlay transport zone is defined, and none is marked as default.
* `dhcp_config_path` - (Optional) Policy path to DHCP server or relay configuration to use for subnets configured on this segment. This attribute is supported with NSX 3.0.0 onwards.
* `subnet` - (Required) Subnet configuration block.
  * `cidr` - (Required) Gateway IP address CIDR. This argument can not be changed if DHCP is enabled for the subnet.
  * `dhcp_ranges` - (Optional) List of DHCP address ranges for dynamic IP allocation.
  * `dhcp_v4_config` - (Optional) DHCPv4 config for IPv4 subnet. This clause is supported with NSX 3.0.0 onwards.
    * `server_address` - (Optional) IP address of the DHCP server in CIDR format. This attribute is required if segment has provided dhcp_config_path and it represents a DHCP server config.
    * `dns_servers` - (Optional) List of IP addresses of DNS servers for the subnet.
    * `lease_time`  - (Optional) DHCP lease time in seconds.
    * `dhcp_option_121` - (Optional) DHCP classless static routes.
      * `network` - (Required) Destination in cidr format.
      * `next_hop` - (Required) IP address of next hop.
    * `dhcp_generic_option` - (Optional) Generic DHCP options.
      * `code` - (Required) DHCP option code. Valid values are from 0 to 255.
      * `values` - (Required) List of DHCP option values.
  * `dhcp_v6_config` - (Optional) DHCPv6 config for IPv6 subnet. This clause is supported with NSX 3.0.0 onwards.
    * `server_address` - (Optional) IP address of the DHCP server in CIDR format. This attribute is required if segment has provided dhcp_config_path and it represents a DHCP server config.
    * `dns_servers` - (Optional) List of IP addresses of DNS servers for the subnet.
    * `lease_time`  - (Optional) DHCP lease time in seconds.
    * `preferred_time` - (Optional) The time interval in seconds, in which the prefix is advertised as preferred.
    * `domain_names` - (Optional) List of domain names for this subnet.
    * `excluded_range` - (Optional) List of excluded address ranges to define dynamic ip allocation ranges.
      * `start` - (Required) IPv6 address that marks beginning of the range.
      * `end` - (Required) IPv6 address that marks end of the range.
    * `sntp_servers` - (Optional) IPv6 address of SNTP servers for the subnet.
* `l2_extension` - (Optional) Configuration for extending Segment through L2 VPN.
  * `l2vpn_paths` - (Optional) Policy paths of associated L2 VPN sessions.
  * `tunnel_id` - (Optional) The Tunnel ID that's a int value between 1 and 4093.
* `advanced_config` - (Optional) Advanced Segment configuration.
  * `address_pool_paths` - (Optional) List of Policy path to IP address pools.
  * `connectivity` - (Optional) Connectivity configuration to manually connect (ON) or disconnect (OFF).
  * `hybrid` - (Optional) Boolean flag to identify a hybrid logical switch.
  * `local_egress` (Optional) Boolean flag to enable local egress.

## Attributes Reference

In addition to arguments listed above, the following attributes are exported:

* `id` - ID of the Secuirty Policy.
* `revision` - Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `path` - The NSX path of the policy resource.
* In the `subnet`:
  * `network` The network CIDR for the subnet.

## Importing

An existing segment can be [imported][docs-import] into this resource, via the following command:

[docs-import]: /docs/import/index.html

```
terraform import nsxt_policy_segment.segment1 ID
```

The above command imports the segment  named `segment1` with the NSX Segment ID `ID`.