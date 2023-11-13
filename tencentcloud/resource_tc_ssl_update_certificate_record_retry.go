/*
Provides a resource to create a ssl update_certificate_record_retry

Example Usage

```hcl
resource "tencentcloud_ssl_update_certificate_record_retry" "update_certificate_record_retry" {
  deploy_record_id =
  deploy_record_detail_id =
}
```

Import

ssl update_certificate_record_retry can be imported using the id, e.g.

```
terraform import tencentcloud_ssl_update_certificate_record_retry.update_certificate_record_retry update_certificate_record_retry_id
```
*/
package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"log"
	"strings"
)

func resourceTencentCloudSslUpdateCertificateRecordRetry() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudSslUpdateCertificateRecordRetryCreate,
		Read:   resourceTencentCloudSslUpdateCertificateRecordRetryRead,
		Delete: resourceTencentCloudSslUpdateCertificateRecordRetryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"deploy_record_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Deployment record ID to be retried.",
			},

			"deploy_record_detail_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Deployment record details ID to be retried.",
			},
		},
	}
}

func resourceTencentCloudSslUpdateCertificateRecordRetryCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_ssl_update_certificate_record_retry.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request              = ssl.NewUpdateCertificateRecordRetryRequest()
		response             = ssl.NewUpdateCertificateRecordRetryResponse()
		deployRecordId       int
		deployRecordDetailId int
	)
	if v, _ := d.GetOk("deploy_record_id"); v != nil {
		deployRecordId = v.(int64)
		request.DeployRecordId = helper.IntInt64(v.(int))
	}

	if v, _ := d.GetOk("deploy_record_detail_id"); v != nil {
		deployRecordDetailId = v.(int64)
		request.DeployRecordDetailId = helper.IntInt64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseSslClient().UpdateCertificateRecordRetry(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate ssl updateCertificateRecordRetry failed, reason:%+v", logId, err)
		return err
	}

	deployRecordId = *response.Response.DeployRecordId
	d.SetId(strings.Join([]string{helper.Int64ToStr(deployRecordId), helper.Int64ToStr(deployRecordDetailId)}, FILED_SP))

	return resourceTencentCloudSslUpdateCertificateRecordRetryRead(d, meta)
}

func resourceTencentCloudSslUpdateCertificateRecordRetryRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_ssl_update_certificate_record_retry.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudSslUpdateCertificateRecordRetryDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_ssl_update_certificate_record_retry.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
