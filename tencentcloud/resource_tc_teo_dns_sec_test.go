package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudTeoDnsSecResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTeoDnsSec,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_teo_dns_sec.dns_sec", "id")),
			},
			{
				ResourceName:      "tencentcloud_teo_dns_sec.dns_sec",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccTeoDnsSec = `

resource "tencentcloud_teo_dns_sec" "dns_sec" {
  zone_id = &lt;nil&gt;
  status = &lt;nil&gt;
    }

`
