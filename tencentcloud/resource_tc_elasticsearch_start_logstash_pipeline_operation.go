/*
Provides a resource to start a elasticsearch logstash pipeline

Example Usage

```hcl
resource "tencentcloud_elasticsearch_start_logstash_pipeline_operation" "start_logstash_pipeline_operation" {
  instance_id = "ls-xxxxxx"
  pipeline_id = "xxxxxx"
}
```
*/
package tencentcloud

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	elasticsearch "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es/v20180416"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudElasticsearchStartLogstashPipelineOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudElasticsearchStartLogstashPipelineOperationCreate,
		Read:   resourceTencentCloudElasticsearchStartLogstashPipelineOperationRead,
		Delete: resourceTencentCloudElasticsearchStartLogstashPipelineOperationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},

			"pipeline_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Pipeline id.",
			},
		},
	}
}

func resourceTencentCloudElasticsearchStartLogstashPipelineOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_elasticsearch_start_logstash_pipeline_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = elasticsearch.NewStartLogstashPipelinesRequest()
		instanceId string
		pipelineId string
	)
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(instanceId)
	}

	if v, ok := d.GetOk("pipeline_id"); ok {
		pipelineId = v.(string)
		request.PipelineIds = []*string{&pipelineId}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseEsClient().StartLogstashPipelines(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate elasticsearch startLogstashPipelineOperation failed, reason:%+v", logId, err)
		return err
	}

	service := ElasticsearchService{client: meta.(*TencentCloudClient).apiV3Conn}

	conf := BuildStateChangeConf([]string{}, []string{"2"}, 3*readRetryTimeout, time.Second, service.ElasticsearchLogstashPipelineStateRefreshFunc(instanceId, pipelineId, []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	d.SetId(instanceId + FILED_SP + pipelineId)

	return resourceTencentCloudElasticsearchStartLogstashPipelineOperationRead(d, meta)
}

func resourceTencentCloudElasticsearchStartLogstashPipelineOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_elasticsearch_start_logstash_pipeline_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudElasticsearchStartLogstashPipelineOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_elasticsearch_start_logstash_pipeline_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
