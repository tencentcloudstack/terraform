/*
Provides a resource to create a vpc dc_gateway_attachment

Example Usage

```hcl
resource "tencentcloud_vpc_dc_gateway_attachment" "dc_gateway_attachment" {
  vpc_id = "vpc-111"
  nat_gateway_id = "nat-test123"
  direct_connect_gateway_id = "dcg-test123"
}
```

Import

vpc dc_gateway_attachment can be imported using the id, e.g.

```
terraform import tencentcloud_vpc_dc_gateway_attachment.dc_gateway_attachment dc_gateway_attachment_id
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"log"
	"strings"
)

func resourceTencentCloudVpcDcGatewayAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudVpcDcGatewayAttachmentCreate,
		Read:   resourceTencentCloudVpcDcGatewayAttachmentRead,
		Delete: resourceTencentCloudVpcDcGatewayAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Vpc id.",
			},

			"nat_gateway_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "NatGatewayId.",
			},

			"direct_connect_gateway_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "DirectConnectGatewayId.",
			},
		},
	}
}

func resourceTencentCloudVpcDcGatewayAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_vpc_dc_gateway_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request                = vpc.NewAssociateDirectConnectGatewayNatGatewayRequest()
		response               = vpc.NewAssociateDirectConnectGatewayNatGatewayResponse()
		vpcId                  string
		directConnectGatewayId string
		natGatewayId           string
	)
	if v, ok := d.GetOk("vpc_id"); ok {
		vpcId = v.(string)
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("nat_gateway_id"); ok {
		natGatewayId = v.(string)
		request.NatGatewayId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("direct_connect_gateway_id"); ok {
		directConnectGatewayId = v.(string)
		request.DirectConnectGatewayId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().AssociateDirectConnectGatewayNatGateway(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc dcGatewayAttachment failed, reason:%+v", logId, err)
		return err
	}

	vpcId = *response.Response.VpcId
	d.SetId(strings.Join([]string{vpcId, directConnectGatewayId, natGatewayId}, FILED_SP))

	return resourceTencentCloudVpcDcGatewayAttachmentRead(d, meta)
}

func resourceTencentCloudVpcDcGatewayAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_vpc_dc_gateway_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	natGatewayId := idSplit[2]

	dcGatewayAttachment, err := service.DescribeVpcDcGatewayAttachmentById(ctx, vpcId, directConnectGatewayId, natGatewayId)
	if err != nil {
		return err
	}

	if dcGatewayAttachment == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `VpcDcGatewayAttachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if dcGatewayAttachment.VpcId != nil {
		_ = d.Set("vpc_id", dcGatewayAttachment.VpcId)
	}

	if dcGatewayAttachment.NatGatewayId != nil {
		_ = d.Set("nat_gateway_id", dcGatewayAttachment.NatGatewayId)
	}

	if dcGatewayAttachment.DirectConnectGatewayId != nil {
		_ = d.Set("direct_connect_gateway_id", dcGatewayAttachment.DirectConnectGatewayId)
	}

	return nil
}

func resourceTencentCloudVpcDcGatewayAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_vpc_dc_gateway_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	natGatewayId := idSplit[2]

	if err := service.DeleteVpcDcGatewayAttachmentById(ctx, vpcId, directConnectGatewayId, natGatewayId); err != nil {
		return err
	}

	return nil
}
