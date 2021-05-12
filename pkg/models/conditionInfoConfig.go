/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"encoding/json"
	"log"
)

type ConditionInfoConfig struct {
	Name           string          `json:"name"`
	ConditionInfos []ConditionInfo `json:"conditionInfos"`
}

type ConditionInfo struct {
	Type     string `json:"type"`
	Expected string `json:"expected"`
}

func defaultServingConditionConfiguration() []byte {
	configurationJSON := `
[
	{
		"name": "ksvc",
		"conditionInfos": [
			{
				"type": "ConfigurationsReady",
				"expected":"True"
			},
			{
				"type": "RoutesReady",
				"expected":"True"
			},
			{
				"type": "Ready",
				"expected":"True"
			}
		]
	},
	{
		"name": "configuration",
		"conditionInfos": [
		]
	},
	{
		"name": "revision",
		"conditionInfos": [
			{
				"type": "ContainerHealthy",
				"expected":"True"
			},
			{
				"type": "ResourcesAvailable",
				"expected":"True"
			},
			{
				"type": "Ready",
				"expected":"True"
			},
			{
				"type": "Active",
				"expected":"True,False"
			}
		]
	},
	{
		"name": "deployment",
		"conditionInfos": [
			{
				"type": "Progressing",
				"expected":"True"
			},
			{
				"type": "Available",
				"expected":"True"
			}
		]
	},
	{
		"name": "replicaset",
		"conditionInfos": [
		]
	},
	{
		"name": "pod",
		"conditionInfos": [
			{
				"type": "PodScheduled",
				"expected":"True"
			},
			{
				"type": "Initialized",
				"expected":"True"
			},
			{
				"type": "ContainersReady",
				"expected":"True"
			},
			{
				"type": "Ready",
				"expected":"True"
			}
		]
	},
	{
		"name": "kpa",
		"conditionInfos": [
			{
				"type": "ScaleTargetInitialized",
				"expected":"True"
			},
			{
				"type": "SKSReady"
			},
			{
				"type": "Ready"
			},
			{
				"type": "Active"
			}
		]
	},
	{
		"name": "metric",
		"conditionInfos": [
		]
	},
	{
		"name": "sks",
		"conditionInfos": [
			{
				"type": "ActivatorEndpointsPopulated",
				"expected":"True"
			},
			{
				"type": "EndpointsPopulated"
			},
			{
				"type": "Ready"
			}
		]
	},
	{
		"name": "route",
		"conditionInfos": [
			{
				"type": "AllTrafficAssigned",
				"expected":"True"
			},
			{
				"type": "CertificateProvisioned",
				"expected":"True"
			},
			{
				"type": "IngressReady",
				"expected":"True"
			},
			{
				"type": "Ready",
				"expected":"True"
			}
		]
	},
	{
		"name": "kingress",
		"conditionInfos": [
			{
				"type": "NetworkConfigured",
				"expected":"True"
			},
			{
				"type": "LoadBalancerReady",
				"expected":"True"
			},
			{
				"type": "Ready",
				"expected":"True"
			}			
		]
	}	
]`

	return []byte(configurationJSON)
}

func LoadServingConditionInfoConfiguration() map[string][]ConditionInfo {

	var configurations []ConditionInfoConfig
	if err := json.Unmarshal(defaultServingConditionConfiguration(), &configurations); err != nil {
		log.Fatal(err)
	}

	conditionMaps := make(map[string][]ConditionInfo)
	for _, item := range configurations {
		if _, ok := conditionMaps[item.Name]; !ok {
			conditionMaps[item.Name] = item.ConditionInfos
		}
	}

	return conditionMaps
}
