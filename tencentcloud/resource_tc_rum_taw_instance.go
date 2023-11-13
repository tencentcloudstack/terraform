/*
Provides a resource to create a rum taw_instance

Example Usage

```hcl
resource "tencentcloud_rum_taw_instance" "taw_instance" {
  area_id = &lt;nil&gt;
  charge_type = &lt;nil&gt;
  data_retention_days = &lt;nil&gt;
  instance_name = &lt;nil&gt;
  tags {
		key = &lt;nil&gt;
		value = &lt;nil&gt;

  }
  instance_desc = &lt;nil&gt;
  count_num = &lt;nil&gt;
  period_retain = &lt;nil&gt;
  buying_channel = &lt;nil&gt;
            tags = {
    "createdBy" = "terraform"
  }
}
```

Import

rum taw_instance can be imported using the id, e.g.

```
terraform import tencentcloud_rum_taw_instance.taw_instance taw_instance_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	rum "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rum/v20210622"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"log"
)

func resourceTencentCloudRumTawInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudRumTawInstanceCreate,
		Read:   resourceTencentCloudRumTawInstanceRead,
		Update: resourceTencentCloudRumTawInstanceUpdate,
		Delete: resourceTencentCloudRumTawInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"area_id": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Region ID (at least greater than 0).",
			},

			"charge_type": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Billing type (1: Pay-as-you-go).",
			},

			"data_retention_days": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Data retention period (at least greater than 0).",
			},

			"instance_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance name (up to 255 bytes).",
			},

			"tags": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Tag list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag key.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag value.",
						},
					},
				},
			},

			"instance_desc": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Instance description (up to 1,024 bytes).",
			},

			"count_num": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Number of data entries reported per day.",
			},

			"period_retain": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Billing for data storage.",
			},

			"buying_channel": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Instance purchase channel. Valid value: cdn.",
			},

			"instance_status": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Instance status (`1` = creating, `2` = running, `3` = exception, `4` = restarting, `5` = stopping, `6` = stopped, `7` = deleted).",
			},

			"cluster_id": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Cluster ID.",
			},

			"charge_status": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Billing status (`1` = in use, `2` = expired, `3` = destroyed, `4` = assigning, `5` = failed).",
			},

			"updated_at": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Update time.",
			},

			"created_at": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Create time.",
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudRumTawInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_rum_taw_instance.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = rum.NewCreateTawInstanceRequest()
		response   = rum.NewCreateTawInstanceResponse()
		instanceId string
	)
	if v, ok := d.GetOkExists("area_id"); ok {
		request.AreaId = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("charge_type"); ok {
		request.ChargeType = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("data_retention_days"); ok {
		request.DataRetentionDays = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("tags"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			tag := rum.Tag{}
			if v, ok := dMap["key"]; ok {
				tag.Key = helper.String(v.(string))
			}
			if v, ok := dMap["value"]; ok {
				tag.Value = helper.String(v.(string))
			}
			request.Tags = append(request.Tags, &tag)
		}
	}

	if v, ok := d.GetOk("instance_desc"); ok {
		request.InstanceDesc = helper.String(v.(string))
	}

	if v, ok := d.GetOk("count_num"); ok {
		request.CountNum = helper.String(v.(string))
	}

	if v, ok := d.GetOk("period_retain"); ok {
		request.PeriodRetain = helper.String(v.(string))
	}

	if v, ok := d.GetOk("buying_channel"); ok {
		request.BuyingChannel = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseRumClient().CreateTawInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create rum tawInstance failed, reason:%+v", logId, err)
		return err
	}

	instanceId = *response.Response.InstanceId
	d.SetId(instanceId)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::rum:%s:uin/:tawInstance/%s", region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudRumTawInstanceRead(d, meta)
}

func resourceTencentCloudRumTawInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_rum_taw_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := RumService{client: meta.(*TencentCloudClient).apiV3Conn}

	tawInstanceId := d.Id()

	tawInstance, err := service.DescribeRumTawInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if tawInstance == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `RumTawInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if tawInstance.AreaId != nil {
		_ = d.Set("area_id", tawInstance.AreaId)
	}

	if tawInstance.ChargeType != nil {
		_ = d.Set("charge_type", tawInstance.ChargeType)
	}

	if tawInstance.DataRetentionDays != nil {
		_ = d.Set("data_retention_days", tawInstance.DataRetentionDays)
	}

	if tawInstance.InstanceName != nil {
		_ = d.Set("instance_name", tawInstance.InstanceName)
	}

	if tawInstance.Tags != nil {
		tagsList := []interface{}{}
		for _, tags := range tawInstance.Tags {
			tagsMap := map[string]interface{}{}

			if tawInstance.Tags.Key != nil {
				tagsMap["key"] = tawInstance.Tags.Key
			}

			if tawInstance.Tags.Value != nil {
				tagsMap["value"] = tawInstance.Tags.Value
			}

			tagsList = append(tagsList, tagsMap)
		}

		_ = d.Set("tags", tagsList)

	}

	if tawInstance.InstanceDesc != nil {
		_ = d.Set("instance_desc", tawInstance.InstanceDesc)
	}

	if tawInstance.CountNum != nil {
		_ = d.Set("count_num", tawInstance.CountNum)
	}

	if tawInstance.PeriodRetain != nil {
		_ = d.Set("period_retain", tawInstance.PeriodRetain)
	}

	if tawInstance.BuyingChannel != nil {
		_ = d.Set("buying_channel", tawInstance.BuyingChannel)
	}

	if tawInstance.InstanceStatus != nil {
		_ = d.Set("instance_status", tawInstance.InstanceStatus)
	}

	if tawInstance.ClusterId != nil {
		_ = d.Set("cluster_id", tawInstance.ClusterId)
	}

	if tawInstance.ChargeStatus != nil {
		_ = d.Set("charge_status", tawInstance.ChargeStatus)
	}

	if tawInstance.UpdatedAt != nil {
		_ = d.Set("updated_at", tawInstance.UpdatedAt)
	}

	if tawInstance.CreatedAt != nil {
		_ = d.Set("created_at", tawInstance.CreatedAt)
	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "rum", "tawInstance", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudRumTawInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_rum_taw_instance.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := rum.NewModifyInstanceRequest()

	tawInstanceId := d.Id()

	request.InstanceId = &instanceId

	immutableArgs := []string{"area_id", "charge_type", "data_retention_days", "instance_name", "tags", "instance_desc", "count_num", "period_retain", "buying_channel", "instance_status", "cluster_id", "charge_status", "updated_at", "created_at"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseRumClient().ModifyInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update rum tawInstance failed, reason:%+v", logId, err)
		return err
	}

	if d.HasChange("tags") {
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("rum", "tawInstance", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudRumTawInstanceRead(d, meta)
}

func resourceTencentCloudRumTawInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_rum_taw_instance.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := RumService{client: meta.(*TencentCloudClient).apiV3Conn}
	tawInstanceId := d.Id()

	if err := service.DeleteRumTawInstanceById(ctx, instanceId); err != nil {
		return err
	}

	return nil
}
