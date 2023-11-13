package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudWafDomainListDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccWafDomainListDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.tencentcloud_waf_domain_list.domain_list")),
			},
		},
	})
}

const testAccWafDomainListDataSource = `

data "tencentcloud_waf_domain_list" "domain_list" {
  filters {
		name = ""
		values = 
		exact_match = 

  }
  }

`
