package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestTencentCloudAsRemoveInstancesResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAsRemoveInstances,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_as_remove_instances.remove_instances", "id")),
			},
		},
	})
}

const testAccAsRemoveInstances = `

resource "tencentcloud_as_remove_instances" "remove_instances" {
  auto_scaling_group_id = ""
  instance_ids = ""
}

`
