package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudSslUpdateCertificateRecordRollbackResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSslUpdateCertificateRecordRollback,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_ssl_update_certificate_record_rollback.update_certificate_record_rollback", "id")),
			},
			{
				ResourceName:      "tencentcloud_ssl_update_certificate_record_rollback.update_certificate_record_rollback",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccSslUpdateCertificateRecordRollback = `

resource "tencentcloud_ssl_update_certificate_record_rollback" "update_certificate_record_rollback" {
  deploy_record_id = 
}

`
