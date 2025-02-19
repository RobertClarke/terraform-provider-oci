// Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "ssh_public_key" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

resource "oci_core_virtual_network" "vcn1" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.vcn1.id}"
  display_name   = "TFExampleSecurityList"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"
  }
}

// An AD based subnet will supply an Availability Domain
resource "oci_core_subnet" "test_subnet" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  cidr_block          = "10.0.2.0/24"
  display_name        = "TFADSubnet"
  dns_label           = "adsubnet"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.vcn1.id}"
  security_list_ids   = ["${oci_core_security_list.test_security_list.id}"]
  route_table_id      = "${oci_core_virtual_network.vcn1.default_route_table_id}"
  dhcp_options_id     = "${oci_core_virtual_network.vcn1.default_dhcp_options_id}"
}

resource "oci_database_db_system" "test_db_system" {
  availability_domain     = "${oci_core_subnet.test_subnet.availability_domain}"
  compartment_id          = "${var.compartment_ocid}"
  subnet_id               = "${oci_core_subnet.test_subnet.id}"
  database_edition        = "ENTERPRISE_EDITION"
  disk_redundancy         = "NORMAL"
  shape                   = "BM.DenseIO2.52"
  cpu_core_count          = "2"
  ssh_public_keys         = ["${var.ssh_public_key}"]
  domain                  = "${oci_core_subnet.test_subnet.subnet_domain_name}"
  hostname                = "myOracleDB"
  data_storage_size_in_gb = "256"
  license_model           = "LICENSE_INCLUDED"
  node_count              = "1"
  display_name            = "TFExampleDbSystem1"

  db_home {
    db_version   = "12.1.0.2"
    display_name = "TFExampleDbHome1"

    database {
      admin_password = "BEstrO0ng_#11"
      db_name        = "tfDbName"
    }
  }
}

data "oci_database_db_homes" "t" {
  compartment_id = "${var.compartment_ocid}"
  db_system_id   = "${oci_database_db_system.test_db_system.id}"

  filter {
    name   = "display_name"
    values = ["TFExampleDbHome1"]
  }
}

resource "oci_database_db_system" "test_db_system2" {
  availability_domain     = "${oci_core_subnet.test_subnet.availability_domain}"
  compartment_id          = "${var.compartment_ocid}"
  subnet_id               = "${oci_core_subnet.test_subnet.id}"
  database_edition        = "ENTERPRISE_EDITION"
  disk_redundancy         = "NORMAL"
  shape                   = "BM.DenseIO2.52"
  cpu_core_count          = "2"
  ssh_public_keys         = ["${var.ssh_public_key}"]
  domain                  = "${oci_core_subnet.test_subnet.subnet_domain_name}"
  hostname                = "myOracleDB"
  data_storage_size_in_gb = "256"
  license_model           = "LICENSE_INCLUDED"
  node_count              = "1"
  display_name            = "TFExampleDbSystem2"

  db_home {
    db_version   = "12.1.0.2"
    display_name = "dbHome1"

    database {
      admin_password = "BEstrO0ng_#11"
      db_name        = "db2"
    }
  }
}

data "oci_database_databases" "db" {
  compartment_id = "${var.compartment_ocid}"
  db_home_id     = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}

resource "oci_database_data_guard_association" "test_data_guard_association" {
  #Required
  creation_type                    = "ExistingDbSystem"
  database_admin_password          = "BEstrO0ng_#11"
  database_id                      = "${data.oci_database_databases.db.databases.0.id}"
  protection_mode                  = "MAXIMUM_PERFORMANCE"
  transport_type                   = "ASYNC"
  delete_standby_db_home_on_delete = "true"

  #required for ExistingDbSystem creation_type
  peer_db_system_id = "${oci_database_db_system.test_db_system2.id}"
}
