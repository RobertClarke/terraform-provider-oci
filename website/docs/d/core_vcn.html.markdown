---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vcn"
sidebar_current: "docs-oci-datasource-core-vcn"
description: |-
  Provides details about a specific Vcn in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_vcn
This data source provides details about a specific Vcn resource in Oracle Cloud Infrastructure Core service.

Gets the specified VCN's information.

## Example Usage

```hcl
data "oci_core_vcn" "test_vcn" {
	#Required
	vcn_id = "${oci_core_vcn.test_vcn.id}"
}
```

## Argument Reference

The following arguments are supported:

* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


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

