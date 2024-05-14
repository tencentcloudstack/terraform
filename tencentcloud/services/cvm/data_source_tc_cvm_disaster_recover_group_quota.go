// Code generated by iacg; DO NOT EDIT.
package cvm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
)

func DataSourceTencentCloudCvmDisasterRecoverGroupQuota() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCvmDisasterRecoverGroupQuotaRead,
		Schema: map[string]*schema.Schema{
			"group_quota": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of placement groups that can be created.",
			},

			"current_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of placement groups that have been created by the current user.",
			},

			"cvm_in_host_group_quota": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Quota on instances in a physical-machine-type disaster recovery group.",
			},

			"cvm_in_sw_group_quota": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Quota on instances in a switch-type disaster recovery group.",
			},

			"cvm_in_rack_group_quota": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Quota on instances in a rack-type disaster recovery group.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCvmDisasterRecoverGroupQuotaRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("data_source.tencentcloud_cvm_disaster_recover_group_quota.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(nil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := CvmService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	paramMap := make(map[string]interface{})
	var respData *cvm.DescribeDisasterRecoverGroupQuotaResponseParams
	err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCvmDisasterRecoverGroupQuotaByFilter(ctx, paramMap)
		if e != nil {
			return tccommon.RetryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	if respData.GroupQuota != nil {
		_ = d.Set("group_quota", respData.GroupQuota)
	}

	if respData.CurrentNum != nil {
		_ = d.Set("current_num", respData.CurrentNum)
	}

	if respData.CvmInHostGroupQuota != nil {
		_ = d.Set("cvm_in_host_group_quota", respData.CvmInHostGroupQuota)
	}

	if respData.CvmInSwGroupQuota != nil {
		_ = d.Set("cvm_in_sw_group_quota", respData.CvmInSwGroupQuota)
	}

	if respData.CvmInRackGroupQuota != nil {
		_ = d.Set("cvm_in_rack_group_quota", respData.CvmInRackGroupQuota)
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := tccommon.WriteToFile(output.(string), dataSourceTencentCloudCvmDisasterRecoverGroupQuotaReadOutputContent(ctx)); e != nil {
			return e
		}
	}

	return nil
}
