// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	CpeRequiredOnlyResource = CpeResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Required, Create, cpeRepresentation)

	cpeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"filter":         RepresentationGroup{Required, cpeDataSourceFilterRepresentation}}
	cpeDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_cpe.test_cpe.id}`}},
	}

	cpeRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"ip_address":     Representation{repType: Required, create: `189.44.2.135`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `MyCpe`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	CpeResourceDependencies = DefinedTagsDependencies
)

func TestCoreCpeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCpeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_cpe.test_cpe"
	datasourceName := "data.oci_core_cpes.test_cpes"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCpeDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CpeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Required, Create, cpeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "189.44.2.135"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CpeResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CpeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Optional, Create, cpeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyCpe"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "189.44.2.135"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CpeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Optional, Update, cpeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "189.44.2.135"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cpes", "test_cpes", Optional, Update, cpeDataSourceRepresentation) +
					compartmentIdVariableStr + CpeResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cpe", "test_cpe", Optional, Update, cpeRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "cpes.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "cpes.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "cpes.0.ip_address", "189.44.2.135"),
					resource.TestCheckResourceAttrSet(datasourceName, "cpes.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreCpeDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cpe" {
			noResourceFound = false
			request := oci_core.GetCpeRequest{}

			tmp := rs.Primary.ID
			request.CpeId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

			_, err := client.GetCpe(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreCpe") {
		resource.AddTestSweepers("CoreCpe", &resource.Sweeper{
			Name:         "CoreCpe",
			Dependencies: DependencyGraph["cpe"],
			F:            sweepCoreCpeResource,
		})
	}
}

func sweepCoreCpeResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient
	cpeIds, err := getCpeIds(compartment)
	if err != nil {
		return err
	}
	for _, cpeId := range cpeIds {
		if ok := SweeperDefaultResourceId[cpeId]; !ok {
			deleteCpeRequest := oci_core.DeleteCpeRequest{}

			deleteCpeRequest.CpeId = &cpeId

			deleteCpeRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteCpe(context.Background(), deleteCpeRequest)
			if error != nil {
				fmt.Printf("Error deleting Cpe %s %s, It is possible that the resource is already deleted. Please verify manually \n", cpeId, error)
				continue
			}
		}
	}
	return nil
}

func getCpeIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "CpeId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient

	listCpesRequest := oci_core.ListCpesRequest{}
	listCpesRequest.CompartmentId = &compartmentId
	listCpesResponse, err := virtualNetworkClient.ListCpes(context.Background(), listCpesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Cpe list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cpe := range listCpesResponse.Items {
		id := *cpe.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "CpeId", id)
	}
	return resourceIds, nil
}
