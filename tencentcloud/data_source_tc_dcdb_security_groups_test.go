package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTencentCloudDcdbSecurityGroupsDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataSourceDcdbSecurityGroups_basic, defaultDcdbInstanceId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_dcdb_security_groups.security_groups"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_dcdb_security_groups.security_groups", "list.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_dcdb_security_groups.security_groups", "list.0.create_time"),
					resource.TestCheckResourceAttr("data.tencentcloud_dcdb_security_groups.security_groups", "list.0.security_group_id", defaultDcdbSGId),
					resource.TestCheckResourceAttr("data.tencentcloud_dcdb_security_groups.security_groups", "list.0.security_group_name", defaultDcdbSGName),
					resource.TestCheckResourceAttrSet("data.tencentcloud_dcdb_security_groups.security_groups", "list.0.inbound.#"),
				),
			},
		},
	})
}

const testAccDataSourceDcdbSecurityGroups_basic = `

data "tencentcloud_dcdb_security_groups" "security_groups" {
  instance_id = "%s"
  }

`
