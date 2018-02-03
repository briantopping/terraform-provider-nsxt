---
layout: "nsxt"
page_title: "NSXT: nsxt_firewall_section"
sidebar_current: "docs-nsxt-resource-firewall-section"
description: |-
  Provides a resource to configure firewall section on NSX-T manager
---

# nsxt_firewall_section

Provides a resource to configure firewall section on NSX-T manager

## Example Usage

```hcl
resource "nsxt_firewall_section" "FS" {
  description = "FS provisioned by Terraform"
  display_name = "FS"
  tags = [{ scope = "color"
            tag = "red" }
  ]
  applied_tos = [{target_type = "NSGroup",
                  target_id = "${nsxt_ns_group.GRP2.id}"}]
  section_type = "LAYER3"
  stateful = true
  rules = [{display_name = "out_rule",
            description = "Out going rule",
            action = "ALLOW",
            logged = "true",
            ip_protocol = "IPV4",
            direction = "OUT",
            sources = [{target_type = "LogicalSwitch", target_id = "${nsxt_logical_switch.LS1.id}"}],
            destinations = [{target_type = "LogicalSwitch", target_id = "${nsxt_logical_switch.LS2.id}"}],
           },
           {display_name = "in_rule",
            description = "In going rule",
            action = "DROP",
            logged = "true",
            ip_protocol = "IPV4",
            direction = "IN",
            services = [{target_type = "NSService", target_id = "e8d59e13-484b-4825-ae3b-4c11f83249d9"},
                        {target_type = "NSService", target_id = "${nsxt_l4_port_set_ns_service.S2.id}"}]
           }
          ]
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) Defaults to ID if not set.
* `description` - (Optional) Description of this resource.
* `tags` - (Optional) A list of scope + tag pairs to associate with this firewall_section.
* `applied_tos` - (Optional) List of objects where the rules in this section will be enforced. This will take precedence over rule level appliedTo. [allowed target types: "LogicalPort", "LogicalSwitch", "NSGroup"]
* `section_type` - (Required) Type of the rules which a section can contain. Either LAYER2 or LAYER3. Only homogeneous sections are supported.
* `stateful` - (Required) Stateful or Stateless nature of firewall section is enforced on all rules inside the section. Layer3 sections can be stateful or stateless. Layer2 sections can only be stateless.
* `rules` - (Optional) A list of rules to be applied in this section. each rule has the following arguments"
  * `display_name` - (Optional) Defaults to ID if not set.
  * `description` - (Optional) Description of this resource.
  * `action` - (Required) Action enforced on the packets which matches the firewall rule.
  * `applied_tos` - (Optional) List of object where rule will be enforced. The section level field overrides this one. Null will be treated as any. [Allowed target types: "LogicalPort", "LogicalSwitch", "NSGroup"]
  * `destinations` - (Optional) List of the destinations. Null will be treated as any. [Allowed target types: "IPSet", "LogicalPort", "LogicalSwitch", "NSGroup", "MACSet" (depending on the section type)]
  * `destinations_excluded` - (Optional) Negation of the destination.
  * `direction` - (Optional) Rule direction in case of stateless firewall rules. This will only considered if section level parameter is set to stateless. Default to IN_OUT if not specified. [Allowed values: "ALLOW", "DROP", "REJECT"]
  * `disabled` - (Optional) Flag to disable rule. Disabled will only be persisted but never provisioned/realized.
  * `ip_protocol` - (Optional) Type of IP packet that should be matched while enforcing the rule. [allowed values: "IPV4", "IPV6", "IPV4_IPV6"]
  * `is_default` - (Optional) Flag to indicate whether rule is default.
  * `logged` - (Optional) Flag to enable packet logging. Default is disabled.
  * `notes` - (Optional) User notes specific to the rule.
  * `rule_tag` - (Optional) User level field which will be printed in CLI and packet logs.
  * `services` - (Optional) List of the services. Null will be treated as any. [Allowed target types: "NSService", "NSServiceGroup"]
  * `sources` - (Optional) List of sources. Null will be treated as any. [Allowed target types: "IPSet", "LogicalPort", "LogicalSwitch", "NSGroup", "MACSet" (depending on the section type)]
  * `sources_excluded` - (Optional) Negation of the source.

## Attributes Reference

In addition to arguments listed above, the following attributes are exported:

* `id` - ID of the firewall_section.
* `revision` - Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.
* `system_owned` - A boolean that indicates whether this resource is system-owned and thus read-only.
* `is_default` - It is a boolean flag which reflects whether a firewall section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section.
* `rule_count` - (Optional) Number of rules in this section.