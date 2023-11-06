package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -test.run TestAccTencentCloudMonitorTmpRegionsDataSource_basic -v
func TestAccTencentCloudMonitorTmpRegionsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMonitorTmpRegionsDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_monitor_tmp_regions.tmp_regions"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.0.area"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.0.region"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.0.region_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.0.region_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.0.region_pay_mode"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.0.region_short_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_monitor_tmp_regions.tmp_regions", "region_set.0.region_state"),
				),
			},
		},
	})
}

const testAccMonitorTmpRegionsDataSource = `

data "tencentcloud_monitor_tmp_regions" "tmp_regions" {
  pay_mode = 1
}

`
