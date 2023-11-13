/*
Provides a resource to create a tcr manage_replication

Example Usage

```hcl
resource "tencentcloud_tcr_manage_replication" "manage_replication" {
  source_registry_id = "tcr-xxx"
  destination_registry_id = "tcr-xxx"
  rule {
		name = "test"
		dest_namespace = "test"
		override = false
		filters {
			type = "tag"
			value = ""
		}

  }
  description = "this is the tcr rule"
  destination_region_id = 1
  peer_replication_option {
		peer_registry_uin = "113498"
		peer_registry_token = "xxx"
		enable_peer_replication = true

  }
}
```

Import

tcr manage_replication can be imported using the id, e.g.

```
terraform import tencentcloud_tcr_manage_replication.manage_replication manage_replication_id
```
*/
package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tcr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcr/v20190924"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"log"
	"strings"
)

func resourceTencentCloudTcrManageReplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudTcrManageReplicationCreate,
		Read:   resourceTencentCloudTcrManageReplicationRead,
		Delete: resourceTencentCloudTcrManageReplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"source_registry_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Copy source instance Id.",
			},

			"destination_registry_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Copy destination instance Id.",
			},

			"rule": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Synchronization rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Synchronization rule names.",
						},
						"dest_namespace": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Target namespace.",
						},
						"override": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to cover.",
						},
						"filters": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Sync filters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Type (name, tag, and resource).",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Empty by default.",
									},
								},
							},
						},
					},
				},
			},

			"description": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Rule description.",
			},

			"destination_region_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "The region ID of the target instance, such as Guangzhou is 1.",
			},

			"peer_replication_option": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Enable synchronization of configuration items across master account instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer_registry_uin": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Uin of the instance to be synchronized.",
						},
						"peer_registry_token": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Access permanent token of the instance to be synchronized.",
						},
						"enable_peer_replication": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to enable cross-master account instance synchronization.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudTcrManageReplicationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_tcr_manage_replication.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request               = tcr.NewManageReplicationRequest()
		response              = tcr.NewManageReplicationResponse()
		sourceRegistryId      string
		destinationRegistryId string
		name                  string
	)
	if v, ok := d.GetOk("source_registry_id"); ok {
		sourceRegistryId = v.(string)
		request.SourceRegistryId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("destination_registry_id"); ok {
		destinationRegistryId = v.(string)
		request.DestinationRegistryId = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "rule"); ok {
		replicationRule := tcr.ReplicationRule{}
		if v, ok := dMap["name"]; ok {
			replicationRule.Name = helper.String(v.(string))
		}
		if v, ok := dMap["dest_namespace"]; ok {
			replicationRule.DestNamespace = helper.String(v.(string))
		}
		if v, ok := dMap["override"]; ok {
			replicationRule.Override = helper.Bool(v.(bool))
		}
		if v, ok := dMap["filters"]; ok {
			for _, item := range v.([]interface{}) {
				filtersMap := item.(map[string]interface{})
				replicationFilter := tcr.ReplicationFilter{}
				if v, ok := filtersMap["type"]; ok {
					replicationFilter.Type = helper.String(v.(string))
				}
				if v, ok := filtersMap["value"]; ok {
					replicationFilter.Value = helper.String(v.(string))
				}
				replicationRule.Filters = append(replicationRule.Filters, &replicationFilter)
			}
		}
		request.Rule = &replicationRule
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, _ := d.GetOk("destination_region_id"); v != nil {
		request.DestinationRegionId = helper.IntUint64(v.(int))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "peer_replication_option"); ok {
		peerReplicationOption := tcr.PeerReplicationOption{}
		if v, ok := dMap["peer_registry_uin"]; ok {
			peerReplicationOption.PeerRegistryUin = helper.String(v.(string))
		}
		if v, ok := dMap["peer_registry_token"]; ok {
			peerReplicationOption.PeerRegistryToken = helper.String(v.(string))
		}
		if v, ok := dMap["enable_peer_replication"]; ok {
			peerReplicationOption.EnablePeerReplication = helper.Bool(v.(bool))
		}
		request.PeerReplicationOption = &peerReplicationOption
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTcrClient().ManageReplication(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate tcr ManageReplication failed, reason:%+v", logId, err)
		return err
	}

	sourceRegistryId = *response.Response.SourceRegistryId
	d.SetId(strings.Join([]string{sourceRegistryId, destinationRegistryId, name}, FILED_SP))

	return resourceTencentCloudTcrManageReplicationRead(d, meta)
}

func resourceTencentCloudTcrManageReplicationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_tcr_manage_replication.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudTcrManageReplicationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_tcr_manage_replication.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
