package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudTsfAddClusterInstancesResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfAddClusterInstances,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_tsf_add_cluster_instances.add_cluster_instances", "id")),
			},
			{
				ResourceName:      "tencentcloud_tsf_add_cluster_instances.add_cluster_instances",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccTsfAddClusterInstances = `

resource "tencentcloud_tsf_add_cluster_instances" "add_cluster_instances" {
  cluster_id = "cluster-123456"
  instance_id_list = 
  os_name = "Ubuntu 20.04"
  image_id = "img-123456"
  password = "MyP@ssw0rd"
  key_id = "key-123456"
  sg_id = "sg-123456"
  instance_import_mode = "R"
  os_customize_type = "my_customize"
  feature_id_list = 
  instance_advanced_settings {
		mount_target = "/mnt/data"
		docker_graph_path = "/var/lib/docker"

  }
  security_group_ids = 
}

`
