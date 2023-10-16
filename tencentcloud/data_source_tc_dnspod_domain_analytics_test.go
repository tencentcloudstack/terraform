package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudDnspodDomain_analyticsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspodDomain_analyticsDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.tencentcloud_dnspod_domain_analytics.domain_analytics")),
			},
		},
	})
}

const testAccDnspodDomain_analyticsDataSource = `

data "tencentcloud_dnspod_domain_analytics" "domain_analytics" {
  domain = "dnspod.cn"
  start_date = "2023-10-07"
  end_date = "2023-10-12"
  dns_format = "HOUR"
  # domain_id = 123
  tags = {
    "createdBy" = "terraform"
  }
}

`
