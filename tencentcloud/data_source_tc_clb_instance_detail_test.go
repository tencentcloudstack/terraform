package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudClbInstanceDetailDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstanceDetailDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.tencentcloud_clb_instance_detail.instance_detail")),
			},
		},
	})
}

const testAccClbInstanceDetailDataSource = `

data "tencentcloud_clb_instance_detail" "instance_detail" {
  fields = 
  target_type = ""
  filters {
		name = ""
		values = 

  }
  }

`
