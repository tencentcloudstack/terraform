// Code generated by iacg; DO NOT EDIT.
package tke

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tkev20180525 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func ResourceTencentCloudKubernetesClusterMasterAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudKubernetesClusterMasterAttachmentCreate,
		Read:   resourceTencentCloudKubernetesClusterMasterAttachmentRead,
		Delete: resourceTencentCloudKubernetesClusterMasterAttachmentDelete,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the cluster.",
			},

			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the CVM instance, this cvm will reinstall the system.",
			},

			"node_role": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node role, values: MASTER_ETCD, WORKER. MASTER_ETCD needs to be specified only when creating an INDEPENDENT_CLUSTER independent cluster. The number of MASTER_ETCD nodes is 3-7, and it is recommended to have an odd number. The minimum configuration for MASTER_ETCD is 4C8G.",
			},

			"enhanced_security_service": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "To specify whether to enable cloud security service. Default is TRUE.",
			},

			"enhanced_monitor_service": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "To specify whether to enable cloud monitor service. Default is TRUE.",
			},

			"enhanced_automation_service": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Activate TencentCloud Automation Tools (TAT) service. If this parameter is not specified, the public image will default to enabling the Cloud Automation Assistant service, while other images will default to not enabling the Cloud Automation Assistant service.",
			},

			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Sensitive:    true,
				Description:  "Password to access, should be set if `key_ids` not set.",
				ValidateFunc: tccommon.ValidateAsConfigPassword,
			},

			"key_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"security_group_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "The security group to which the instance belongs. This parameter can be obtained by calling the sgId field in the return value of DescribeSecureGroups. If this parameter is not specified, the default security group will be bound.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "When reinstalling the system, you can specify the HostName of the instance to be modified (this parameter must be passed when the cluster is in HostName mode, and the rule name should be consistent with the HostName of the CVM instance creation interface except that uppercase characters are not supported).",
			},

			"desired_pod_numbers": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "When the node belongs to the podCIDR size customization mode, the maximum number of pods running on the node can be specified.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"extra_args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "Custom parameters for cluster master component.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kube_api_server": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Description: "Kube apiserver custom parameters. The parameter format is [\"k1=v1\", \"k1=v2\"].",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"kube_controller_manager": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Description: "Kube controller manager custom parameters.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"kube_scheduler": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Description: "kube scheduler custom parameters.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"etcd": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Description: "etcd custom parameters. Only supports independent clusters.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"master_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "Advanced Node Settings. commonly used to attach existing instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mount_target": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Mount target. Default is not mounting.",
						},
						"docker_graph_path": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Docker graph path. Default is `/var/lib/docker`.",
						},
						"user_script": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "User script encoded in base64, which will be executed after the k8s component runs. The user needs to ensure the script's reentrant and retry logic. The script and its generated log files can be viewed in the node path /data/ccs_userscript/. If the node needs to be initialized before joining the schedule, it can be used in conjunction with the `unschedulable` parameter. After the final initialization of the userScript is completed, add the command \"kubectl uncordon nodename --kubeconfig=/root/.kube/config\" to add the node to the schedule.",
						},
						"unschedulable": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "Set whether the joined nodes participate in scheduling, with a default value of 0, indicating participation in scheduling; Non 0 means not participating in scheduling.",
						},
						"labels": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: "Node label list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "Name of map.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "Value of map.",
									},
								},
							},
						},
						"data_disk": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "Configurations of data disk.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disk_type": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Types of disk. Valid value: `LOCAL_BASIC`, `LOCAL_SSD`, `CLOUD_BASIC`, `CLOUD_PREMIUM`, `CLOUD_SSD`, `CLOUD_HSSD`, `CLOUD_TSSD` and `CLOUD_BSSD`.",
									},
									"file_system": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "File system, e.g. `ext3/ext4/xfs`.",
									},
									"disk_size": {
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Description: "Volume of disk in GB. Default is `0`.",
									},
									"auto_format_and_mount": {
										Type:        schema.TypeBool,
										Optional:    true,
										ForceNew:    true,
										Description: "Indicate whether to auto format and mount or not. Default is `false`.",
									},
									"mount_target": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Mount target.",
									},
									"disk_partition": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The name of the device or partition to mount. NOTE: this argument doesn't support setting in node pool, or will leads to mount error.",
									},
								},
							},
						},
						"extra_args": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "Custom parameter information related to the node. This is a white-list parameter.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kubelet": {
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Description: "Kubelet custom parameter. The parameter format is [\"k1=v1\", \"k1=v2\"].",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"desired_pod_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "Indicate to set desired pod number in node. valid when the cluster is podCIDR.",
						},
						"gpu_args": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "GPU driver parameters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mig_enable": {
										Type:        schema.TypeBool,
										Optional:    true,
										ForceNew:    true,
										Description: "Whether to enable MIG.",
									},
									"driver": {
										Type:         schema.TypeMap,
										Optional:     true,
										ForceNew:     true,
										Description:  "GPU driver version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
										ValidateFunc: tccommon.ValidateTkeGpuDriverVersion,
									},
									"cuda": {
										Type:         schema.TypeMap,
										Optional:     true,
										ForceNew:     true,
										Description:  "CUDA  version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
										ValidateFunc: tccommon.ValidateTkeGpuDriverVersion,
									},
									"cudnn": {
										Type:         schema.TypeMap,
										Optional:     true,
										ForceNew:     true,
										Description:  "cuDNN version. Format like: `{ version: String, name: String, doc_name: String, dev_name: String }`. `version`: cuDNN version; `name`: cuDNN name; `doc_name`: Doc name of cuDNN; `dev_name`: Dev name of cuDNN.",
										ValidateFunc: tccommon.ValidateTkeGpuDriverVersion,
									},
									"custom_driver": {
										Type:        schema.TypeMap,
										Optional:    true,
										ForceNew:    true,
										Description: "Custom GPU driver. Format like: `{address: String}`. `address`: URL of custom GPU driver address.",
									},
								},
							},
						},
						"taints": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: "Node taint.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Key of the taint.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Value of the taint.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Effect of the taint.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudKubernetesClusterMasterAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_kubernetes_cluster_master_attachment.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	var (
		clusterId  string
		instanceId string
		nodeRole   string
	)
	var (
		request  = tkev20180525.NewScaleOutClusterMasterRequest()
		response = tkev20180525.NewScaleOutClusterMasterResponse()
	)

	if v, ok := d.GetOk("cluster_id"); ok {
		clusterId = v.(string)
	}
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
	}
	if v, ok := d.GetOk("node_role"); ok {
		nodeRole = v.(string)
	}

	request.ClusterId = helper.String(clusterId)

	if err := resourceTencentCloudKubernetesClusterMasterAttachmentCreatePostFillRequest0(ctx, request); err != nil {
		return err
	}

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTkeV20180525Client().ScaleOutClusterMasterWithContext(ctx, request)
		if e != nil {
			return resourceTencentCloudKubernetesClusterMasterAttachmentCreateRequestOnError0(ctx, request, e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create kubernetes cluster master attachment failed, reason:%+v", logId, err)
		return err
	}

	_ = response

	if err := resourceTencentCloudKubernetesClusterMasterAttachmentCreatePostHandleResponse0(ctx, response); err != nil {
		return err
	}

	d.SetId(strings.Join([]string{clusterId, instanceId, nodeRole}, tccommon.FILED_SP))

	return resourceTencentCloudKubernetesClusterMasterAttachmentRead(d, meta)
}

func resourceTencentCloudKubernetesClusterMasterAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_kubernetes_cluster_master_attachment.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := TkeService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	instanceId := idSplit[1]
	nodeRole := idSplit[2]

	_ = d.Set("cluster_id", clusterId)

	_ = d.Set("instance_id", instanceId)

	_ = d.Set("node_role", nodeRole)

	respData, err := service.DescribeKubernetesClusterMasterAttachmentById(ctx, clusterId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `kubernetes_cluster_master_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	respData1, err := service.DescribeKubernetesClusterMasterAttachmentById1(ctx, instanceId)
	if err != nil {
		return err
	}

	if respData1 == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `kubernetes_cluster_master_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	var respData2 *tkev20180525.DescribeClusterInstancesResponseParams
	reqErr2 := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeKubernetesClusterMasterAttachmentById2(ctx, clusterId, instanceId, nodeRole)
		if e != nil {
			return resourceTencentCloudKubernetesClusterMasterAttachmentReadRequestOnError2(ctx, result, e)
		}
		if err := resourceTencentCloudKubernetesClusterMasterAttachmentReadRequestOnSuccess2(ctx, result); err != nil {
			return err
		}
		respData2 = result
		return nil
	})
	if reqErr2 != nil {
		log.Printf("[CRITAL]%s read kubernetes cluster master attachment failed, reason:%+v", logId, reqErr2)
		return reqErr2
	}

	if respData2 == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `kubernetes_cluster_master_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	return nil
}

func resourceTencentCloudKubernetesClusterMasterAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_kubernetes_cluster_master_attachment.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	idSplit := strings.Split(d.Id(), tccommon.FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	instanceId := idSplit[1]
	nodeRole := idSplit[2]

	var (
		request  = tkev20180525.NewScaleInClusterMasterRequest()
		response = tkev20180525.NewScaleInClusterMasterResponse()
	)

	request.ClusterId = helper.String(clusterId)

	if err := resourceTencentCloudKubernetesClusterMasterAttachmentDeletePostFillRequest0(ctx, request); err != nil {
		return err
	}

	err := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
		result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UseTkeV20180525Client().ScaleInClusterMasterWithContext(ctx, request)
		if e != nil {
			return resourceTencentCloudKubernetesClusterMasterAttachmentDeleteRequestOnError0(ctx, e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete kubernetes cluster master attachment failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	_ = instanceId
	_ = nodeRole
	return nil
}
