package main

import (
	"encoding/json"
	"os"
	"testing"

	alibaba "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/alibaba/utils/alibaba"
	_ "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/alibaba/utils/env"
)

// //------ Cluster Management
// *CreateCluster(clusterReqInfo ClusterInfo) (ClusterInfo, error)
// *ListCluster() ([]*ClusterInfo, error)
// *GetCluster(clusterIID IID) (ClusterInfo, error)
// *DeleteCluster(clusterIID IID) (bool, error)

// //------ NodeGroup Management
// *AddNodeGroup(clusterIID IID, nodeGroupReqInfo NodeGroupInfo) (NodeGroupInfo, error)
// *ListNodeGroup(clusterIID IID) ([]*NodeGroupInfo, error)
// *GetNodeGroup(clusterIID IID, nodeGroupIID IID) (NodeGroupInfo, error)
// -SetNodeGroupAutoScaling(clusterIID IID, nodeGroupIID IID, on bool) (bool, error)
// -ChangeNodeGroupScaling(clusterIID IID, nodeGroupIID IID, DesiredNodeSize int, MinNodeSize int, MaxNodeSize int) (NodeGroupInfo, error)
// *RemoveNodeGroup(clusterIID IID, nodeGroupIID IID) (bool, error)

// //------ Upgrade K8S
// -UpgradeCluster(clusterIID IID, newVersion string) (ClusterInfo, error)

var access_key string
var access_secret string
var region_id string

func setup() {
	println("setup")
	access_key = os.Getenv("ACCESS_KEY")
	access_secret = os.Getenv("ACCESS_SECRET")
	region_id = os.Getenv("REGION_ID")
}

func shutdown() {
	println("shutdown")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestCreateClusterOnly(t *testing.T) {
	// nodecount = 0
	t.Log("클러스터만 생성하기")

	body := `{
		"name": "cluster_2",
		"region_id": "cn-beijing",
		"cluster_type": "ManagedKubernetes",
		"kubernetes_version": "1.22.10-aliyun.1",
		"vpcid": "vpc-2zek5slojo5bh621ftnrg",
		"container_cidr": "172.24.0.0/16",
		"service_cidr": "172.23.0.0/16",
		"num_of_nodes": 0,
		"master_vswitch_ids": [
			"vsw-2ze0qpwcio7r5bx3nqbp1"
		]
	}`

	// body := `{
	// 	"name": "cluster_2",
	// 	"region_id": "cn-beijing",
	// 	"cluster_type": "ManagedKubernetes",
	// 	"kubernetes_version": "1.22.10-aliyun.1",
	// 	"vpcid": "vpc-2zek5slojo5bh621ftnrg",
	// 	"container_cidr": "172.24.0.0/16",
	// 	"service_cidr": "172.23.0.0/16",
	// 	"key_pair": "kp1",
	// 	"login_password": "$etri2ETRI",
	// 	"master_vswitch_ids": [
	// 		"vsw-2ze0qpwcio7r5bx3nqbp1"
	// 	],
	// 	"master_instance_types": [
	// 		"ecs.g7ne.xlarge,ecs.c7.xlarge"
	// 	],
	// 	"master_system_disk_category": "cloud_essd",
	// 	"master_system_disk_size": 120,
	// 	"num_of_nodes": 0,
	// 	"vswitch_ids": [
	// 		"vsw-2ze0qpwcio7r5bx3nqbp1"
	// 	],
	// 	"worker_vswitch_ids": [
	// 		"vsw-2ze0qpwcio7r5bx3nqbp1"
	// 	],
	// 	"worker_instance_types": [
	// 		"ecs.g7ne.xlarge,ecs.c7.xlarge"
	// 	],
	// 	"worker_system_disk_category": "cloud_essd",
	// 	"worker_system_disk_size": 120,
	// 	"worker_data_disks": [
	// 		{
	// 			"category": "cloud_essd",
	// 			"size": "120"
	// 		}
	// 	]
	// }`

	// "master_count": 3,

	/*
		must contain three (large, lowercase letters, numbers and special symbols).
		$etri2ETRI
	*/

	/*
	   Message: {"code":"ZoneNotSupported","message":
	   "The current zone  does not support creating SLB, please try other zones,
	   request id: 2C47B8CA-E920-5E54-BBC4-47081431E780",
	   "requestId":"E9EE3C8A-71CF-5047-8BA7-15A81C954A10","status":400}
	*/

	result, err := alibaba.CreateCluster(access_key, access_secret, region_id, body)
	if err != nil {
		t.Errorf("Failed to create cluster: %v", err)
	}
	println(result)
}

func TestCreateClusterWithNodeGroup(t *testing.T) {

	t.Log("클러스터 + 노드그룹(1) 생성하기")

	body := `{
		"name": "cluster_0",
		"region_id": "cn-beijing",
		"cluster_type": "ManagedKubernetes",
		"kubernetes_version": "1.22.10-aliyun.1",
		"vpcid": "vpc-2zek5slojo5bh621ftnrg",
		"container_cidr": "172.21.0.0/16",
		"service_cidr": "172.22.0.0/16",
		"key_pair": "kp1",
		"login_password": "$etri2ETRI",
		"master_vswitch_ids": [
			"vsw-2ze0qpwcio7r5bx3nqbp1"
		],
		"master_instance_types": [
			"ecs.g7ne.xlarge,ecs.c7.xlarge"
		],
		"master_system_disk_category": "cloud_essd",
		"master_system_disk_size": 120,
		"num_of_nodes": 0,
		"vswitch_ids": [
			"vsw-2ze0qpwcio7r5bx3nqbp1"
		],
		"worker_vswitch_ids": [
			"vsw-2ze0qpwcio7r5bx3nqbp1"
		],
		"worker_instance_types": [
			"ecs.g7ne.xlarge,ecs.c7.xlarge"
		],
		"worker_system_disk_category": "cloud_essd",
		"worker_system_disk_size": 120,
		"worker_data_disks": [
			{
				"category": "cloud_essd",
				"size": "120"
			}
		]
	}`

	// "master_count": 3,

	/*
		must contain three (large, lowercase letters, numbers and special symbols).
		$etri2ETRI
	*/

	result, err := alibaba.CreateCluster(access_key, access_secret, region_id, body)
	if err != nil {
		t.Errorf("Failed to create cluster: %v", err)
	}
	println(result)
}

func TestGetClusters(t *testing.T) {
	result, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(result)
}

func TestGetCluster(t *testing.T) {
	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)
		println(cluster_id)

		cluster_json_str, err := alibaba.GetCluster(access_key, access_secret, region_id, cluster_id)
		if err != nil {
			t.Errorf("Failed to get cluster: %v", err)
		}
		println(cluster_json_str)

		cluster_json_obj := make(map[string]interface{})
		json.Unmarshal([]byte(cluster_json_str), &cluster_json_obj)
		print(cluster_json_obj["cluster_id"].(string))
	}
}

func TestDeleteCluster(t *testing.T) {
	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)
		println(cluster_id)
		temp, err := alibaba.DeleteCluster(access_key, access_secret, region_id, cluster_id)
		if err != nil {
			t.Errorf("Failed to delete cluster: %v", err)
		}
		println(temp)
	}
}

func TestCreateNodeGroup(t *testing.T) {

	// body  := `{
	// 	"auto_scaling":{
	// 		"enable":true,
	// 		"max_instances":10 ,
	// 		"min_instances":1
	// 	},
	// 	"scaling_group":{
	// 		"desired_size":5
	// 	}
	// }`

	body := `{
		"auto_scaling": {
			"enable": true,
			"max_instances": 5,
			"min_instances": 0
		},
		"kubernetes_config": {
			"runtime": "containerd",
			"runtime_version": "1.5.10"
		},
		"nodepool_info": {
			"name": "nodepoolx"
		},
		"scaling_group": {	
			"instance_charge_type": "PostPaid",
			"instance_types": [
				"ecs.c6.xlarge"
			],
			"key_pair": "kp1",
			"system_disk_category": "cloud_essd",
			"system_disk_size": 70,
			"vswitch_ids": [
				"vsw-2ze0qpwcio7r5bx3nqbp1"
			]
		},
		"management": {
			" enable":true
		}	
	}`

	// desired_size/count setting or modification is not supported for autoscaling-enabled nodepool
	// body := `{
	// 	"auto_scaling": {
	// 		"enable": true,
	// 		"max_instances": 5,
	// 		"min_instances": 0
	// 	},
	// 	"kubernetes_config": {
	// 		"runtime": "containerd",
	// 		"runtime_version": "1.5.10"
	// 	},
	// 	"nodepool_info": {
	// 		"name": "nodepoolx"
	// 	},
	// 	"scaling_group": {
	// 		// "desired_size":1,
	// 		"instance_charge_type": "PostPaid",
	// 		"instance_types": [
	// 			"ecs.c6.xlarge"
	// 		],
	// 		"key_pair": "kp1",
	// 		"system_disk_category": "cloud_essd",
	// 		"system_disk_size": 70,
	// 		"vswitch_ids": [
	// 			"vsw-2ze0qpwcio7r5bx3nqbp1"
	// 		]
	// 	},
	// 	"management": {
	// 		" enable":true
	// 	}
	// }`

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)
		println(cluster_id)

		clusters_json_str, err := alibaba.CreateNodeGroup(access_key, access_secret, region_id, cluster_id, body)
		if err != nil {
			t.Errorf("Failed to create node group: %v", err)
		}
		println(clusters_json_str)
	}

}

// c870636966d134b968a960cd9d978f940
func TestCreateNodeGroup2(t *testing.T) {

	// body := `{
	// 	"nodepool_info": {
	// 		"name": "nodepoolx"
	// 	},
	// 	"auto_scaling": {
	// 		"enable": true,
	// 		"max_instances": 5,
	// 		"min_instances": 0
	// 	},
	// 	"scaling_group": {
	// 		"instance_charge_type": "PostPaid",
	// 		"instance_types": ["ecs.c6.xlarge"],
	// 		"key_pair": "kp1",
	// 		"system_disk_category": "cloud_essd",
	// 		"system_disk_size": 70,
	// 		"vswitch_ids": ["vsw-2ze0qpwcio7r5bx3nqbp1"]
	// 	},
	// 	"management": {
	// 		"enable":true
	// 	}
	// }`

	body := `{
		"nodepool_info": {
			"name": "nodepoolx3"
		},
		"auto_scaling": {
			"enable": true,
			"max_instances": 5,
			"min_instances": 2
		},
		"scaling_group": {
			"instance_types": ["ecs.c6.xlarge"],
			"key_pair": "kp1",
			"system_disk_category": "cloud_essd",
			"system_disk_size": 70,
			"vswitch_ids": ["vsw-2ze0qpwcio7r5bx3nqbp1"]
		},
		"management": {
			"enable":true
		}
	}`

	// "kubernetes_config": {
	// 	"runtime": "containerd",
	// 	"runtime_version": "1.5.10"
	// },

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)
		println(cluster_id)

		clusters_json_str, err := alibaba.CreateNodeGroup(access_key, access_secret, region_id, cluster_id, body)
		if err != nil {
			t.Errorf("Failed to create node group: %v", err)
		}
		println(clusters_json_str)
	}

}

func TestListNodeGroup(t *testing.T) {

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)

		nodepools_json_str, err := alibaba.ListNodeGroup(access_key, access_secret, region_id, cluster_id)
		if err != nil {
			t.Errorf("Failed to list node group: %v", err)
		}
		println(nodepools_json_str)
		nodepools_json_obj := make(map[string]interface{})
		json.Unmarshal([]byte(nodepools_json_str), &nodepools_json_obj)
		nodepools := nodepools_json_obj["nodepools"].([]interface{})
		for _, v := range nodepools {
			node_group_id := v.(map[string]interface{})["nodepool_info"].(map[string]interface{})["nodepool_id"].(string)
			println(node_group_id)

			// get node group
			nodepool_json_str, err := alibaba.GetNodeGroup(access_key, access_secret, region_id, cluster_id, node_group_id)
			if err != nil {
				t.Errorf("Failed to get node group: %v", err)
			}
			println(nodepool_json_str)
		}
	}
}

func TestGetNodeGroup(t *testing.T) {

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}

	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)

		nodepools_json_str, err := alibaba.ListNodeGroup(access_key, access_secret, region_id, cluster_id)
		if err != nil {
			t.Errorf("Failed to list node group: %v", err)
		}

		println(nodepools_json_str)
		nodepools_json_obj := make(map[string]interface{})
		json.Unmarshal([]byte(nodepools_json_str), &nodepools_json_obj)
		nodepools := nodepools_json_obj["nodepools"].([]interface{})
		for _, v := range nodepools {
			node_group_id := v.(map[string]interface{})["nodepool_info"].(map[string]interface{})["nodepool_id"].(string)
			println(node_group_id)

			// get node group
			nodepool_json_str, err := alibaba.GetNodeGroup(access_key, access_secret, region_id, cluster_id, node_group_id)
			if err != nil {
				t.Errorf("Failed to get node group: %v", err)
			}
			println(nodepool_json_str)
		}
	}
}

func TestSetNodeGroupAutoScaling(t *testing.T) {
	//  https://next.api.alibabacloud.com/api/CS/2015-12-15/ModifyClusterNodePool?sdkStyle=old&params={}
	// modify (set auto scaling) on/off
	// body := `{"auto_scaling":{"enable":false}}`
	// body := `{"auto_scaling":{"enable":true}}`
	// body := `{"auto_scaling":{"max_instances":5,"min_instances":0},"scaling_group":{"desired_size":2}}`

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)

		nodepools_json_str, err := alibaba.ListNodeGroup(access_key, access_secret, region_id, cluster_id)
		if err != nil {
			t.Errorf("Failed to list node group: %v", err)
		}
		println(nodepools_json_str)
		nodepools_json_obj := make(map[string]interface{})
		json.Unmarshal([]byte(nodepools_json_str), &nodepools_json_obj)
		nodepools := nodepools_json_obj["nodepools"].([]interface{})
		for _, v := range nodepools {
			node_group_id := v.(map[string]interface{})["nodepool_info"].(map[string]interface{})["nodepool_id"].(string)
			println(node_group_id)

			body := `{"auto_scaling":{"enable":false}}`
			res, err := alibaba.ModifyNodeGroup(access_key, access_secret, region_id, cluster_id, node_group_id, body)
			if err != nil {
				t.Errorf("Failed to modify node group: %v", err)
			}
			println(res)

			body = `{"auto_scaling":{"enable":true}}`
			res, err = alibaba.ModifyNodeGroup(access_key, access_secret, region_id, cluster_id, node_group_id, body)
			if err != nil {
				t.Errorf("Failed to modify node group: %v", err)
			}
			println(res)
			// "{\"code\":\"ErrDefaultNodePool\",\"message\":\" Nodepool is default, cannot enable autoscaling\"}\n"

			// default node pool: cannot enable autoscaling
			// custom, managed node pool: can enable autoscaling
			// https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/manage-node-pools
			// body  := `{"auto_scaling":{"enable":true,"max_instances":10 ,"min_instances":1}, "scaling_group":{"desired_size":5},"management":{" enable":true}}`

		}
	}
}

func TestChangeNodeGroupScaling(t *testing.T) {
	// modify (set auto scaling) on/off count
	// body := `{"auto_scaling":{"enable":false}}`
	// body := `{"auto_scaling":{"enable":true}}`
	// body := `{"auto_scaling":{"max_instances":5,"min_instances":0},"scaling_group":{"desired_size":2}}`

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)

		nodepools_json_str, err := alibaba.ListNodeGroup(access_key, access_secret, region_id, cluster_id)
		if err != nil {
			t.Errorf("Failed to list node group: %v", err)
		}
		println(nodepools_json_str)
		nodepools_json_obj := make(map[string]interface{})
		json.Unmarshal([]byte(nodepools_json_str), &nodepools_json_obj)
		nodepools := nodepools_json_obj["nodepools"].([]interface{})
		for _, v := range nodepools {
			node_group_id := v.(map[string]interface{})["nodepool_info"].(map[string]interface{})["nodepool_id"].(string)
			println(node_group_id)

			body := `{"auto_scaling":{"max_instances":10,"min_instances":0},"scaling_group":{"desired_size":2}}`
			res, err := alibaba.ModifyNodeGroup(access_key, access_secret, region_id, cluster_id, node_group_id, body)
			if err != nil {
				t.Errorf("Failed to modify node group: %v", err)
			}
			println(res)

			body = `{"auto_scaling":{"max_instances":3,"min_instances":1},"scaling_group":{"desired_size":1}}`
			res, err = alibaba.ModifyNodeGroup(access_key, access_secret, region_id, cluster_id, node_group_id, body)
			if err != nil {
				t.Errorf("Failed to modify node group: %v", err)
			}
			println(res)
		}
	}
}

func TestDeleteNodeGroup(t *testing.T) {

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)

		nodepools_json_str, err := alibaba.ListNodeGroup(access_key, access_secret, region_id, cluster_id)
		if err != nil {
			t.Errorf("Failed to list node group: %v", err)
		}
		println(nodepools_json_str)
		nodepools_json_obj := make(map[string]interface{})
		json.Unmarshal([]byte(nodepools_json_str), &nodepools_json_obj)
		nodepools := nodepools_json_obj["nodepools"].([]interface{})
		for _, v := range nodepools {
			node_group_id := v.(map[string]interface{})["nodepool_info"].(map[string]interface{})["nodepool_id"].(string)
			println(node_group_id)

			name := v.(map[string]interface{})["nodepool_info"].(map[string]interface{})["name"].(string)
			println(name)

			if name == "nodepoolx" {
				temp, err := alibaba.DeleteNodeGroup(access_key, access_secret, region_id, cluster_id, node_group_id)
				if err != nil {
					t.Errorf("Failed to delete node group: %v", err)
				}
				println(temp)
			}
		}
	}
}

func TestUpgradeCluster(t *testing.T) {

	// POST /api/v2/clusters/c82e6987e2961451182edacd74faf****/upgrade HTTP/1.1
	// Content-Type:application/json
	// {
	//   "component_name" : "k8s",
	//   "next_version" : "1.16.9-aliyun.1",
	//   "version" : "1.14.8-aliyun.1"
	// }

	//https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/kubernetes-1-22-release-notes#concept-2170736
	// 1.22.3-aliyun.1
	// {
	//   "next_version" : "1.22.3-aliyun.1"
	// }

	clusters_json_str, err := alibaba.GetClusters(access_key, access_secret, region_id)
	if err != nil {
		t.Errorf("Failed to get clusters: %v", err)
	}
	println(clusters_json_str)

	clusters_json_obj := make(map[string]interface{})
	json.Unmarshal([]byte(clusters_json_str), &clusters_json_obj)

	clusters := clusters_json_obj["clusters"].([]interface{})
	for _, v := range clusters {
		cluster_id := v.(map[string]interface{})["cluster_id"].(string)

		version := `{"next_version" : "1.22.3-aliyun.1"}`
		res, err := alibaba.UpgradeCluster(access_key, access_secret, region_id, cluster_id, version)
		if err != nil {
			t.Errorf("Failed to upgrade cluster: %v", err)
		}
		println(res)
	}
}
