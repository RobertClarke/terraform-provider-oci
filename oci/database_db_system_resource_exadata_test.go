// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/oracle/oci-go-sdk/database"
)

// TestAccResourceDatabaseDBSystem_Exadata tests DBsystems using Exadata
func TestResourceDatabaseDBSystemExaData(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemExaData")
	defer httpreplay.SaveScenario()

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "DBSystem_Exadata") {
		t.Skip("Skipping suppressed DBSystem_Exadata")
	}

	provider := testAccProvider
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: ResourceDatabaseBaseConfig + `
				data "oci_identity_availability_domain" "ad" {
  					compartment_id 		= "${var.compartment_id}"
  					ad_number      		= 3
				}
				resource "oci_core_subnet" "exadata_subnet" {
					availability_domain = "${data.oci_identity_availability_domain.ad.name}"
					cidr_block          = "10.1.22.0/24"
					display_name        = "ExadataSubnet"
					compartment_id      = "${var.compartment_id}"
					vcn_id              = "${oci_core_virtual_network.t.id}"
					route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
					dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
					security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}", "${oci_core_security_list.exadata_shapes_security_list.id}"]
					dns_label           = "subnetexadata1"
				}
				resource "oci_core_subnet" "exadata_backup_subnet" {
					availability_domain = "${data.oci_identity_availability_domain.ad.name}"
					cidr_block          = "10.1.23.0/24"
					display_name        = "ExadataBackupSubnet"
					compartment_id      = "${var.compartment_id}"
					vcn_id              = "${oci_core_virtual_network.t.id}"
					route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
					dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
					security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
					dns_label           = "subnetexadata2"
				}

				resource "oci_core_security_list" "exadata_shapes_security_list" {
					compartment_id = "${var.compartment_id}"
					vcn_id         = "${oci_core_virtual_network.t.id}"
					display_name   = "ExadataSecurityList"

					ingress_security_rules {
						source    = "10.1.22.0/24"
						protocol  = "6"
					}

					ingress_security_rules {
						source    = "10.1.22.0/24"
						protocol  = "1"
					}

					egress_security_rules {
						destination = "10.1.22.0/24"
						protocol    = "6"
					}

					egress_security_rules {
						destination = "10.1.22.0/24"
						protocol    = "1"
					}
				}

				resource "oci_database_db_system" "t" {
					availability_domain = "${data.oci_identity_availability_domain.ad.name}"
					compartment_id = "${var.compartment_id}"
					subnet_id = "${oci_core_subnet.exadata_subnet.id}"
					backup_subnet_id = "${oci_core_subnet.exadata_backup_subnet.id}"
					database_edition = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
					disk_redundancy = "NORMAL"
					shape = "Exadata.Quarter1.84"
					cpu_core_count = "22"
					ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
					domain = "${oci_core_subnet.exadata_subnet.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
					hostname = "myOracleDB"
					data_storage_size_in_gb = "256"
					license_model = "LICENSE_INCLUDED"
					node_count = "1"
					time_zone = "US/Pacific"
					db_home {
						db_version = "12.1.0.2"
						database {
							admin_password = "BEstrO0ng_#11"
							db_name = "aTFdb"
						}
					}
				}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// DB System Resource tests
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "time_created"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "database_edition", "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "disk_redundancy", "NORMAL"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "shape", "Exadata.Quarter1.84"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "cpu_core_count", "22"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.#", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "ssh_public_keys.0", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),
					resource.TestMatchResourceAttr(ResourceDatabaseResourceName, "display_name", regexp.MustCompile(`dbsystem\d+`)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "domain", "subnetexadata1.tfvcn.oraclevcn.com"),
					resource.TestCheckResourceAttrSet(ResourceDatabaseResourceName, "hostname"), // see comment in SetData fn as to why this is removed
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "data_storage_size_in_gb", "256"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "license_model", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "node_count", "1"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.db_version", "12.1.0.2"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "db_home.0.database.0.db_name", "aTFdb"),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "state", string(database.DatabaseLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(ResourceDatabaseResourceName, "time_zone", "US/Pacific"),
				),
			},
		},
	})
}
