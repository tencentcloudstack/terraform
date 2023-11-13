/*
Provides a resource to create a cdb ro_replication

Example Usage

```hcl
resource "tencentcloud_cdb_ro_replication" "ro_replication" {
  instance_id = ""
}
```

Import

cdb ro_replication can be imported using the id, e.g.

```
terraform import tencentcloud_cdb_ro_replication.ro_replication ro_replication_id
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"log"
	"time"
)

func resourceTencentCloudCdbRoReplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCdbRoReplicationCreate,
		Read:   resourceTencentCloudCdbRoReplicationRead,
		Delete: resourceTencentCloudCdbRoReplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Read-Only instance ID.",
			},
		},
	}
}

func resourceTencentCloudCdbRoReplicationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cdb_ro_replication.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = cdb.NewStartReplicationRequest()
		response   = cdb.NewStartReplicationResponse()
		instanceId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCdbClient().StartReplication(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cdb roReplication failed, reason:%+v", logId, err)
		return err
	}

	instanceId = *response.Response.InstanceId
	d.SetId(instanceId)

	service := CdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	conf := BuildStateChangeConf([]string{}, []string{"SUCCEED"}, 1*readRetryTimeout, time.Second, service.CdbRoReplicationStateRefreshFunc(d.Id(), []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	return resourceTencentCloudCdbRoReplicationRead(d, meta)
}

func resourceTencentCloudCdbRoReplicationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cdb_ro_replication.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	roReplicationId := d.Id()

	roReplication, err := service.DescribeCdbRoReplicationById(ctx, instanceId)
	if err != nil {
		return err
	}

	if roReplication == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CdbRoReplication` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if roReplication.InstanceId != nil {
		_ = d.Set("instance_id", roReplication.InstanceId)
	}

	return nil
}

func resourceTencentCloudCdbRoReplicationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cdb_ro_replication.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CdbService{client: meta.(*TencentCloudClient).apiV3Conn}
	roReplicationId := d.Id()

	if err := service.DeleteCdbRoReplicationById(ctx, instanceId); err != nil {
		return err
	}

	service := CdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	conf := BuildStateChangeConf([]string{}, []string{"SUCCEED"}, 1*readRetryTimeout, time.Second, service.CdbRoReplicationStateRefreshFunc(d.Id(), []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	return nil
}
