/*
Provides a resource to create a tsf lane_rule

Example Usage

```hcl
resource "tencentcloud_tsf_lane_rule" "lane_rule" {
    rule_name = ""
  remark = ""
  rule_tag_list {
		tag_id = ""
		tag_name = ""
		tag_operator = ""
		tag_value = ""
		lane_rule_id = ""
		create_time = 
		update_time = 

  }
  rule_tag_relationship = ""
  lane_id = ""
          program_id_list = 
}
```

Import

tsf lane_rule can be imported using the id, e.g.

```
terraform import tencentcloud_tsf_lane_rule.lane_rule lane_rule_id
```
*/
package tencentcloud

import (
"context"
"fmt"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
tsf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsf/v20180326"
"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
"log"
)


func resourceTencentCloudTsfLaneRule () * schema.Resource {
return & schema.Resource {
Create : resourceTencentCloudTsfLaneRuleCreate ,
Read : resourceTencentCloudTsfLaneRuleRead ,
Update : resourceTencentCloudTsfLaneRuleUpdate ,
Delete : resourceTencentCloudTsfLaneRuleDelete ,
Importer : & schema.ResourceImporter {
State : schema.ImportStatePassthrough ,
} ,
Schema : map[string] * schema.Schema {
"rule_id": {
  Computed: true,
  Type: schema.TypeString,
  Description: "Rule id.",
},

"rule_name": {
  Required: true,
  Type: schema.TypeString,
  Description: "Lane rule name.",
},

"remark": {
  Required: true,
  Type: schema.TypeString,
  Description: "Lane rule notes.",
},

"rule_tag_list": {
  Required: true,
  Type: schema.TypeList,
  Description: "List of swimlane rule labels.",
Elem: &schema.Resource{
    Schema: map[string]*schema.Schema{
      "tag_id": {
  Type: schema.TypeString,
  Required: true,
  Description: "Label ID.",
},
"tag_name": {
  Type: schema.TypeString,
  Required: true,
  Description: "Label name.",
},
"tag_operator": {
  Type: schema.TypeString,
  Required: true,
  Description: "Label operator.",
},
"tag_value": {
  Type: schema.TypeString,
  Required: true,
  Description: "Tag value.",
},
"lane_rule_id": {
  Type: schema.TypeString,
  Required: true,
  Description: "Lane rule ID.",
},
"create_time": {
  Type: schema.TypeInt,
  Required: true,
  Description: "Creation time.",
},
"update_time": {
  Type: schema.TypeInt,
  Required: true,
  Description: "Update time.",
},

    },
  },
},

"rule_tag_relationship": {
  Required: true,
  Type: schema.TypeString,
  Description: "Lane rule label relationship.",
},

"lane_id": {
  Required: true,
  Type: schema.TypeString,
  Description: "Lane ID.",
},

"priority": {
  Computed: true,
  Type: schema.TypeInt,
  Description: "Priority.",
},

"enable": {
  Computed: true,
  Type: schema.TypeBool,
  Description: "Open state.",
},

"create_time": {
  Computed: true,
  Type: schema.TypeInt,
  Description: "Creation time.",
},

"update_time": {
  Computed: true,
  Type: schema.TypeInt,
  Description: "Update time.",
},

"program_id_list": {
  Optional: true,
  Type: schema.TypeSet,
  Elem: &schema.Schema{
				Type: schema.TypeString,
	},
  Description: "���.",
},

} ,
}
} 

func resourceTencentCloudTsfLaneRuleCreate (d * schema.ResourceData,meta interface{}) error {
defer logElapsed ("resource.tencentcloud_tsf_lane_rule.create") ()
defer inconsistentCheck (d,meta) () 

logId := getLogId (contextNil) 

var (
request = tsf.NewCreateLaneRuleRequest ()
response = tsf.NewCreateLaneRuleResponse ()
ruleId string
)
if v,ok := d . GetOk ("rule_name");ok {
request . RuleName = helper.String (v .(string))
} 

if v,ok := d . GetOk ("remark");ok {
request . Remark = helper.String (v .(string))
} 

if v,ok := d . GetOk ("rule_tag_list");ok {
for _,item := range v .([] interface{}) {
dMap := item .(map[string] interface{})
laneRuleTag := tsf . LaneRuleTag {}
if v,ok := dMap ["tag_id"];ok {
laneRuleTag . TagId = helper.String (v .(string))
}
if v,ok := dMap ["tag_name"];ok {
laneRuleTag . TagName = helper.String (v .(string))
}
if v,ok := dMap ["tag_operator"];ok {
laneRuleTag . TagOperator = helper.String (v .(string))
}
if v,ok := dMap ["tag_value"];ok {
laneRuleTag . TagValue = helper.String (v .(string))
}
if v,ok := dMap ["lane_rule_id"];ok {
laneRuleTag . LaneRuleId = helper.String (v .(string))
}
if v,ok := dMap ["create_time"];ok {
laneRuleTag . CreateTime = helper.IntInt64 (v .(int))
}
if v,ok := dMap ["update_time"];ok {
laneRuleTag . UpdateTime = helper.IntInt64 (v .(int))
}
request . RuleTagList = append(request . RuleTagList,& laneRuleTag)
}
} 

if v,ok := d . GetOk ("rule_tag_relationship");ok {
request . RuleTagRelationship = helper.String (v .(string))
} 

if v,ok := d . GetOk ("lane_id");ok {
request . LaneId = helper.String (v .(string))
} 

if v,ok := d . GetOk ("program_id_list");ok {
programIdListSet := v .(* schema.Set) . List () 
 for i := range programIdListSet {
programIdList := programIdListSet [i] .(string)
request . ProgramIdList = append(request . ProgramIdList,& programIdList)
}
} 

err := resource.Retry (writeRetryTimeout,func () * resource.RetryError {
result,e := meta .(* TencentCloudClient) . apiV3Conn . UseTsfClient () . CreateLaneRule (request)
if e != nil {
return  retryError (e)
} else {
log.Printf ("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",logId,request . GetAction (),request . ToJsonString (),result . ToJsonString ())
}
response = result
return  nil
})
if err != nil {
log.Printf ("[CRITAL]%s create tsf laneRule failed, reason:%+v",logId,err)
return  err
} 

ruleId = * response . Response . ruleId
d . SetId (ruleId) 

return resourceTencentCloudTsfLaneRuleRead (d,meta)
} 

func resourceTencentCloudTsfLaneRuleRead (d * schema.ResourceData,meta interface{}) error {
defer logElapsed ("resource.tencentcloud_tsf_lane_rule.read") ()
defer inconsistentCheck (d,meta) () 

logId := getLogId (contextNil) 

ctx := context.WithValue (context.TODO (),logIdKey,logId) 

service := TsfService {client:meta .(* TencentCloudClient) . apiV3Conn} 

laneRuleId := d . Id ()


laneRule,err := service . DescribeTsfLaneRuleById (ctx,ruleId)
if err != nil {
return err
} 

if laneRule == nil {
d . SetId ("")
log.Printf ("[WARN]%s resource `TsfLaneRule` [%s] not found, please check if it has been deleted.\n",logId,d . Id ())
return nil
} 

if laneRule . RuleId != nil {
_ = d . Set ("rule_id",laneRule . RuleId)
} 

if laneRule . RuleName != nil {
_ = d . Set ("rule_name",laneRule . RuleName)
} 

if laneRule . Remark != nil {
_ = d . Set ("remark",laneRule . Remark)
} 

if laneRule . RuleTagList != nil {
ruleTagListList := [] interface{} {}
for _,ruleTagList := range laneRule . RuleTagList {
ruleTagListMap := map[string] interface{} {} 

if laneRule.RuleTagList . TagId != nil {
ruleTagListMap ["tag_id"] = laneRule.RuleTagList . TagId
} 

if laneRule.RuleTagList . TagName != nil {
ruleTagListMap ["tag_name"] = laneRule.RuleTagList . TagName
} 

if laneRule.RuleTagList . TagOperator != nil {
ruleTagListMap ["tag_operator"] = laneRule.RuleTagList . TagOperator
} 

if laneRule.RuleTagList . TagValue != nil {
ruleTagListMap ["tag_value"] = laneRule.RuleTagList . TagValue
} 

if laneRule.RuleTagList . LaneRuleId != nil {
ruleTagListMap ["lane_rule_id"] = laneRule.RuleTagList . LaneRuleId
} 

if laneRule.RuleTagList . CreateTime != nil {
ruleTagListMap ["create_time"] = laneRule.RuleTagList . CreateTime
} 

if laneRule.RuleTagList . UpdateTime != nil {
ruleTagListMap ["update_time"] = laneRule.RuleTagList . UpdateTime
} 

ruleTagListList = append(ruleTagListList,ruleTagListMap)
} 

_ = d . Set ("rule_tag_list",ruleTagListList) 

} 

if laneRule . RuleTagRelationship != nil {
_ = d . Set ("rule_tag_relationship",laneRule . RuleTagRelationship)
} 

if laneRule . LaneId != nil {
_ = d . Set ("lane_id",laneRule . LaneId)
} 

if laneRule . Priority != nil {
_ = d . Set ("priority",laneRule . Priority)
} 

if laneRule . Enable != nil {
_ = d . Set ("enable",laneRule . Enable)
} 

if laneRule . CreateTime != nil {
_ = d . Set ("create_time",laneRule . CreateTime)
} 

if laneRule . UpdateTime != nil {
_ = d . Set ("update_time",laneRule . UpdateTime)
} 

if laneRule . ProgramIdList != nil {
_ = d . Set ("program_id_list",laneRule . ProgramIdList)
} 

return nil
} 

func resourceTencentCloudTsfLaneRuleUpdate (d * schema.ResourceData,meta interface{}) error {
defer logElapsed ("resource.tencentcloud_tsf_lane_rule.update") ()
defer inconsistentCheck (d,meta) () 

logId := getLogId (contextNil) 

request := tsf.NewModifyLaneRuleRequest () 



laneRuleId := d . Id ()


request . RuleId = & ruleId


immutableArgs := [] string {"rule_id","rule_name","remark","rule_tag_list","rule_tag_relationship","lane_id","priority","enable","create_time","update_time","program_id_list"}


for _,v := range immutableArgs {
if d . HasChange (v) {
return fmt.Errorf ("argument `%s` cannot be changed",v)
}
}


if d . HasChange ("rule_name") {
if v,ok := d . GetOk ("rule_name");ok {
request . RuleName = helper.String (v .(string))
}
} 

if d . HasChange ("remark") {
if v,ok := d . GetOk ("remark");ok {
request . Remark = helper.String (v .(string))
}
} 

if d . HasChange ("rule_tag_list") {
if v,ok := d . GetOk ("rule_tag_list");ok {
for _,item := range v .([] interface{}) {
laneRuleTag := tsf . LaneRuleTag {}
if v,ok := dMap ["tag_id"];ok {
laneRuleTag . TagId = helper.String (v .(string))
}
if v,ok := dMap ["tag_name"];ok {
laneRuleTag . TagName = helper.String (v .(string))
}
if v,ok := dMap ["tag_operator"];ok {
laneRuleTag . TagOperator = helper.String (v .(string))
}
if v,ok := dMap ["tag_value"];ok {
laneRuleTag . TagValue = helper.String (v .(string))
}
if v,ok := dMap ["lane_rule_id"];ok {
laneRuleTag . LaneRuleId = helper.String (v .(string))
}
if v,ok := dMap ["create_time"];ok {
laneRuleTag . CreateTime = helper.IntInt64 (v .(int))
}
if v,ok := dMap ["update_time"];ok {
laneRuleTag . UpdateTime = helper.IntInt64 (v .(int))
}
request . RuleTagList = append(request . RuleTagList,& laneRuleTag)
}
}
} 

if d . HasChange ("rule_tag_relationship") {
if v,ok := d . GetOk ("rule_tag_relationship");ok {
request . RuleTagRelationship = helper.String (v .(string))
}
} 

if d . HasChange ("lane_id") {
if v,ok := d . GetOk ("lane_id");ok {
request . LaneId = helper.String (v .(string))
}
} 

err := resource.Retry (writeRetryTimeout,func () * resource.RetryError {
result,e := meta .(* TencentCloudClient) . apiV3Conn . UseTsfClient () . ModifyLaneRule (request)
if e != nil {
return  retryError (e)
} else {
log.Printf ("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",logId,request . GetAction (),request . ToJsonString (),result . ToJsonString ())
}
return  nil
})
if err != nil {
log.Printf ("[CRITAL]%s update tsf laneRule failed, reason:%+v",logId,err)
return  err
} 

return resourceTencentCloudTsfLaneRuleRead (d,meta)
} 

func resourceTencentCloudTsfLaneRuleDelete (d * schema.ResourceData,meta interface{}) error {
defer logElapsed ("resource.tencentcloud_tsf_lane_rule.delete") ()
defer inconsistentCheck (d,meta) () 

logId := getLogId (contextNil)
ctx := context.WithValue (context.TODO (),logIdKey,logId) 

service := TsfService {client:meta .(* TencentCloudClient) . apiV3Conn}
laneRuleId := d . Id ()


if err := service . DeleteTsfLaneRuleById (ctx,ruleId) ; err != nil {
return err
} 

return nil
} 
