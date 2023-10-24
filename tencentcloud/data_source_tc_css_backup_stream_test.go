package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudCssBackupStreamDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCssBackupStreamDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.tencentcloud_css_backup_stream.backup_stream")),
			},
		},
	})
}

const testAccCssBackupStreamDataSource = `

data "tencentcloud_css_backup_stream" "backup_stream" {
  stream_name = ""
}

`
