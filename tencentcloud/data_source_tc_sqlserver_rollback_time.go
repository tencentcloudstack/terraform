/*
Use this data source to query detailed information of sqlserver rollback_time

Example Usage

```hcl
data "tencentcloud_sqlserver_rollback_time" "rollback_time" {
  instance_id = "mssql-j8kv137v"
  d_bs =
  }
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sqlserver "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver/v20180328"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func dataSourceTencentCloudSqlserverRollbackTime() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudSqlserverRollbackTimeRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance ID.",
			},

			"d_bs": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of databases to be queried.",
			},

			"details": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Information of time range available for database rollback.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"d_b_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Database name.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Start time of time range available for rollback.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "End time of time range available for rollback.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudSqlserverRollbackTimeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloud_sqlserver_rollback_time.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["InstanceId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("d_bs"); ok {
		dBsSet := v.(*schema.Set).List()
		paramMap["DBs"] = helper.InterfacesStringsPoint(dBsSet)
	}

	service := SqlserverService{client: meta.(*TencentCloudClient).apiV3Conn}

	var details []*sqlserver.DbRollbackTimeInfo

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeSqlserverRollbackTimeByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		details = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(details))
	tmpList := make([]map[string]interface{}, 0, len(details))

	if details != nil {
		for _, dbRollbackTimeInfo := range details {
			dbRollbackTimeInfoMap := map[string]interface{}{}

			if dbRollbackTimeInfo.DBName != nil {
				dbRollbackTimeInfoMap["d_b_name"] = dbRollbackTimeInfo.DBName
			}

			if dbRollbackTimeInfo.StartTime != nil {
				dbRollbackTimeInfoMap["start_time"] = dbRollbackTimeInfo.StartTime
			}

			if dbRollbackTimeInfo.EndTime != nil {
				dbRollbackTimeInfoMap["end_time"] = dbRollbackTimeInfo.EndTime
			}

			ids = append(ids, *dbRollbackTimeInfo.InstanceId)
			tmpList = append(tmpList, dbRollbackTimeInfoMap)
		}

		_ = d.Set("details", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
