---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vcn"
sidebar_current: "docs-oci-resource-core-vcn"
description: |-
  Provides the Vcn resource in Oracle Cloud Infrastructure Core service
---

# oci_core_vcn
This resource provides the Vcn resource in Oracle Cloud Infrastructure Core service.

The VCN automatically comes with a default route table, default security list, and default set of DHCP options.
For managing these resources, see [Managing Default VCN Resources](/docs/providers/oci/guides/managing_default_resources.html)

Creates a new Virtual Cloud Network (VCN). For more information, see
[VCNs and Subnets](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVCNs.htm).

For the VCN you must specify a single, contiguous IPv4 CIDR block. Oracle recommends using one of the
private IP address ranges specified in [RFC 1918](https://tools.ietf.org/html/rfc1918) (10.0.0.0/8,
172.16/12, and 192.168/16). Example: 172.16.0.0/16. The CIDR block can range from /16 to /30, and it
must not overlap with your on-premises network. You can't change the size of the VCN after creation.

For the purposes of access control, you must provide the OCID of the compartment where you want the VCN to
reside. Consult an Oracle Cloud Infrastructure administrator in your organization if you're not sure which
compartment to use. Notice that the VCN doesn't have to be in the same compartment as the subnets or other
Networking Service components. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the VCN, otherwise a default is provided. It does not have to
be unique, and you can change it. Avoid entering confidential information.

You can also add a DNS label for the VCN, which is required if you want the instances to use the
Interent and VCN Resolver option for DNS in the VCN. For more information, see
[DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

The VCN automatically comes with a default route table, default security list, and default set of DHCP options.
The OCID for each is returned in the response. You can't delete these default objects, but you can change their
contents (that is, change the route rules, security list rules, and so on).

The VCN and subnets you create are not accessible until you attach an internet gateway or set up an IPSec VPN
or FastConnect. For more information, see
[Overview of the Networking Service](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm).


## Supported Aliases

* `oci_core_virtual_network`

## Example Usage

```hcl
resource "oci_core_vcn" "test_vcn" {
	#Required
	cidr_block = "${var.vcn_cidr_block}"
	compartment_id = "${var.compartment_id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.vcn_display_name}"
	dns_label = "${var.vcn_dns_label}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `cidr_block` - (Required) The CIDR IP address block of the VCN.  Example: `172.16.0.0/16` 
* `compartment_id` - (Required) The OCID of the compartment to contain the VCN.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_label` - (Optional) A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Not required to be unique, but it's a best practice to set unique DNS labels for VCNs in your tenancy. Must be an alphanumeric string that begins with a letter. The value cannot be changed.

	You must set this value if you want instances to be able to use hostnames to resolve other instances in the VCN. Otherwise the Internet and VCN Resolver will not work.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cidr_block` - The CIDR IP address block of the VCN.  Example: `172.16.0.0/16` 
* `compartment_id` - The OCID of the compartment containing the VCN.
* `default_dhcp_options_id` - The OCID for the VCN's default set of DHCP options. 
* `default_route_table_id` - The OCID for the VCN's default route table.
* `default_security_list_id` - The OCID for the VCN's default security list.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_label` - A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter. The value cannot be changed.

	The absence of this parameter means the Internet and VCN Resolver will not work for this VCN.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The VCN's Oracle ID (OCID).
* `state` - The VCN's current state.
* `time_created` - The date and time the VCN was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_domain_name` - The VCN's domain name, which consists of the VCN's DNS label, and the `oraclevcn.com` domain.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1.oraclevcn.com` 

## Import

Vcns can be imported using the `id`, e.g.

```
$ terraform import oci_core_vcn.test_vcn "id"
```

