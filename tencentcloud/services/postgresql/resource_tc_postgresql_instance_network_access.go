package postgresql

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	postgresqlv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres/v20170312"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func ResourceTencentCloudPostgresqlInstanceNetworkAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudPostgresqlInstanceNetworkAccessCreate,
		Read:   resourceTencentCloudPostgresqlInstanceNetworkAccessRead,
		Delete: resourceTencentCloudPostgresqlInstanceNetworkAccessDelete,
		Importer: &schema.ResourceImporter{
			StateContext: networkAccessCustomResourceImporter,
		},
		Schema: map[string]*schema.Schema{
			"db_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance ID in the format of postgres-6bwgamo3.",
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unified VPC ID.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subnet ID.",
			},

			"is_assign_vip": {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Whether to manually assign the VIP. Valid values: `true` (manually assign), `false` (automatically assign).",
			},

			"vip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Target VIP.",
			},
		},
	}
}

func resourceTencentCloudPostgresqlInstanceNetworkAccessCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_postgresql_instance_network_access.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	var (
		logId        = tccommon.GetLogId(tccommon.ContextNil)
		ctx          = tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)
		request      = postgresqlv20170312.NewCreateDBInstanceNetworkAccessRequest()
		response     = postgresqlv20170312.NewCreateDBInstanceNetworkAccessResponse()
		dbInsntaceId string
		vpcId        string
		subnetId     string
		vip          string
	)

	if v, ok := d.GetOk("db_instance_id"); ok {
		request.DBInstanceId = helper.String(v.(string))
		dbInsntaceId = v.(string)
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
		vpcId = v.(string)
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
		subnetId = v.(string)
	}

	if v, ok := d.GetOkExists("is_assign_vip"); ok {
		request.IsAssignVip = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("vip"); ok {
		request.Vip = helper.String(v.(string))
		vip = v.(string)
	}

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UsePostgresqlV20170312Client().CreateDBInstanceNetworkAccessWithContext(ctx, request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Create postgresql instance network access failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create postgresql instance network access failed, reason:%+v", logId, err)
		return err
	}

	if response.Response.FlowId == nil {
		return fmt.Errorf("FlowId is nil.")
	}

	// wait & get vip
	flowRequest := postgresqlv20170312.NewDescribeTasksRequest()
	flowRequest.TaskId = response.Response.FlowId
	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UsePostgresqlV20170312Client().DescribeTasksWithContext(ctx, request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Create postgresql instance network access failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create postgresql instance network access failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{dbInsntaceId, vpcId, subnetId, vip}, tccommon.FILED_SP))

	return resourceTencentCloudPostgresqlInstanceNetworkAccessRead(d, meta)
}

func resourceTencentCloudPostgresqlInstanceNetworkAccessRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_postgresql_instance_network_access.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := PostgresqlService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	dbInsntaceId := idSplit[0]
	vpcId := idSplit[1]
	subnetId := idSplit[2]
	vip := idSplit[3]

	_ = d.Set("db_instance_id", dbInsntaceId)

	_ = d.Set("vpc_id", vpcId)

	_ = d.Set("subnet_id", subnetId)

	respData, err := service.DescribePostgresqlInstanceNetworkAccessById(ctx, dbInsntaceId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `postgresql_instance_network_access` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if err := resourceTencentCloudPostgresqlInstanceNetworkAccessReadPreHandleResponse0(ctx, respData); err != nil {
		return err
	}

	if respData.VpcId != nil {
		_ = d.Set("vpc_id", respData.VpcId)
	}

	if respData.SubnetId != nil {
		_ = d.Set("subnet_id", respData.SubnetId)
	}

	if respData.DBInstanceId != nil {
		_ = d.Set("db_instance_id", respData.DBInstanceId)
	}

	_ = vpcId
	_ = subnetId
	_ = vip
	return nil
}

func resourceTencentCloudPostgresqlInstanceNetworkAccessDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_postgresql_instance_network_access.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	dbInsntaceId := idSplit[0]
	vpcId := idSplit[1]
	subnetId := idSplit[2]
	vip := idSplit[3]

	var (
		request  = postgresqlv20170312.NewDeleteDBInstanceNetworkAccessRequest()
		response = postgresqlv20170312.NewDeleteDBInstanceNetworkAccessResponse()
	)

	request.DBInstanceId = helper.String(dbInsntaceId)

	request.VpcId = helper.String(vpcId)

	request.SubnetId = helper.String(subnetId)

	request.Vip = helper.String(vip)

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UsePostgresqlV20170312Client().DeleteDBInstanceNetworkAccessWithContext(ctx, request)
		if e != nil {
			return tccommon.RetryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete postgresql instance network access failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	if _, err := (&resource.StateChangeConf{
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
		Pending:    []string{},
		Refresh:    resourcePostgresqlInstanceNetworkAccessDeleteStateRefreshFunc_0_0(ctx, dbInsntaceId, vpcId, subnetId, vip),
		Target:     []string{"Running"},
		Timeout:    180 * time.Second,
	}).WaitForStateContext(ctx); err != nil {
		return err
	}
	return nil
}

func resourcePostgresqlInstanceNetworkAccessCreateStateRefreshFunc_0_0(ctx context.Context, dbInsntaceId string, vpcId string, subnetId string, vip string) resource.StateRefreshFunc {
	var req *postgresqlv20170312.DescribeDBInstanceAttributeRequest
	return func() (interface{}, string, error) {
		meta := tccommon.ProviderMetaFromContext(ctx)
		if meta == nil {
			return nil, "", fmt.Errorf("resource data can not be nil")
		}
		if req == nil {
			d := tccommon.ResourceDataFromContext(ctx)
			if d == nil {
				return nil, "", fmt.Errorf("resource data can not be nil")
			}
			_ = d
			req = postgresqlv20170312.NewDescribeDBInstanceAttributeRequest()
			req.DBInstanceId = helper.String(dbInsntaceId)

		}
		resp, err := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UsePostgresqlV20170312Client().DescribeDBInstanceAttributeWithContext(ctx, req)
		if err != nil {
			return nil, "", err
		}
		if resp == nil || resp.Response == nil {
			return nil, "", nil
		}
		state := fmt.Sprintf("%v", *resp.Response.DBInstance.DBInstanceStatus)
		return resp.Response, state, nil
	}
}

func resourcePostgresqlInstanceNetworkAccessDeleteStateRefreshFunc_0_0(ctx context.Context, dbInsntaceId string, vpcId string, subnetId string, vip string) resource.StateRefreshFunc {
	var req *postgresqlv20170312.DescribeDBInstanceAttributeRequest
	return func() (interface{}, string, error) {
		meta := tccommon.ProviderMetaFromContext(ctx)
		if meta == nil {
			return nil, "", fmt.Errorf("resource data can not be nil")
		}
		if req == nil {
			d := tccommon.ResourceDataFromContext(ctx)
			if d == nil {
				return nil, "", fmt.Errorf("resource data can not be nil")
			}
			_ = d
			req = postgresqlv20170312.NewDescribeDBInstanceAttributeRequest()
			req.DBInstanceId = helper.String(dbInsntaceId)

		}
		resp, err := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UsePostgresqlV20170312Client().DescribeDBInstanceAttributeWithContext(ctx, req)
		if err != nil {
			return nil, "", err
		}
		if resp == nil || resp.Response == nil {
			return nil, "", nil
		}
		state := fmt.Sprintf("%v", *resp.Response.DBInstance.DBInstanceStatus)
		return resp.Response, state, nil
	}
}
