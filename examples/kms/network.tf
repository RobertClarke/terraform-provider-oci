// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_virtual_network" "my_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_id}"
  display_name   = "myvcn"
  dns_label      = "myvcn"
}

resource "oci_core_internet_gateway" "my_internet_gateway" {
  compartment_id = "${var.compartment_id}"
  display_name   = "my internet gateway"
  vcn_id         = "${oci_core_virtual_network.my_vcn.id}"
}

resource "oci_core_route_table" "my_route_table" {
  compartment_id = "${var.compartment_id}"
  vcn_id         = "${oci_core_virtual_network.my_vcn.id}"
  display_name   = "my route table"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.my_internet_gateway.id}"
  }
}

resource "oci_core_subnet" "my_subnet" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "mysubnet"
  dns_label           = "mysubnet"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.my_vcn.id}"
  security_list_ids   = ["${oci_core_virtual_network.my_vcn.default_security_list_id}"]
  route_table_id      = "${oci_core_route_table.my_route_table.id}"
}
