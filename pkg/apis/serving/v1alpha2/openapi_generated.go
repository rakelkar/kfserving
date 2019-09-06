// +build !ignore_autogenerated

/*
Copyright 2019 kubeflow.org.

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

// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	openapispec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.AlibiExplainerSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "AlibiExplainerSpec defines the arguments for configuring an Alibi Explanation Server",
					Properties: map[string]openapispec.Schema{
						"type": {
							SchemaProps: openapispec.SchemaProps{
								Description: "The type of Alibi explainer",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"storageUri": {
							SchemaProps: openapispec.SchemaProps{
								Description: "The location of a trained explanation model",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to latest Alibi Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
						"config": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Inline custom parameter settings for explainer",
								Type:        []string{"object"},
								AdditionalProperties: &openapispec.SchemaOrBool{
									Schema: &openapispec.Schema{
										SchemaProps: openapispec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
					Required: []string{"type"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "CustomSpec provides a hook for arbitrary container configuration.",
					Properties: map[string]openapispec.Schema{
						"container": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("k8s.io/api/core/v1.Container"),
							},
						},
					},
					Required: []string{"container"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.Container"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.DeploymentSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "DeploymentSpec defines the configuration for a given KFService service component",
					Properties: map[string]openapispec.Schema{
						"serviceAccountName": {
							SchemaProps: openapispec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Properties: map[string]openapispec.Schema{
						"predictor": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Predictor defines the model serving spec",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorSpec"),
							},
						},
						"explainer": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Explainer defines the model explanation service spec explainer service calls to transformer or predictor service",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerSpec"),
							},
						},
						"transformer": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Transformer defines the transformer service spec for pre/post processing transformer service calls to predictor service",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TransformerSpec"),
							},
						},
					},
					Required: []string{"predictor"},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TransformerSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ExplainerSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "ExplainerSpec defines the arguments for a model explanation server",
					Properties: map[string]openapispec.Schema{
						"alibi": {
							SchemaProps: openapispec.SchemaProps{
								Description: "The following fields follow a \"1-of\" semantic. Users must specify exactly one openapispec.",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.AlibiExplainerSpec"),
							},
						},
						"custom": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"),
							},
						},
						"serviceAccountName": {
							SchemaProps: openapispec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.AlibiExplainerSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Properties: map[string]openapispec.Schema{
						"image": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
					},
					Required: []string{"image"},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworksConfig": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Properties: map[string]openapispec.Schema{
						"tensorflow": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig"),
							},
						},
						"tensorrt": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig"),
							},
						},
						"xgboost": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig"),
							},
						},
						"sklearn": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig"),
							},
						},
						"pytorch": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig"),
							},
						},
						"onnx": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.FrameworkConfig"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFService": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "KFService is the Schema for the services API",
					Properties: map[string]openapispec.Schema{
						"kind": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"apiVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"metadata": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
							},
						},
						"spec": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceSpec"),
							},
						},
						"status": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceStatus"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceList": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "KFServiceList contains a list of Service",
					Properties: map[string]openapispec.Schema{
						"kind": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"apiVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"metadata": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
							},
						},
						"items": {
							SchemaProps: openapispec.SchemaProps{
								Type: []string{"array"},
								Items: &openapispec.SchemaOrArray{
									Schema: &openapispec.Schema{
										SchemaProps: openapispec.SchemaProps{
											Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFService"),
										},
									},
								},
							},
						},
					},
					Required: []string{"items"},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFService", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "KFServiceSpec defines the desired state of KFService",
					Properties: map[string]openapispec.Schema{
						"default": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Default defines default KFService endpoints",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec"),
							},
						},
						"canary": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Canary defines an alternate endpoints to route a percentage of traffic.",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec"),
							},
						},
						"canaryTrafficPercent": {
							SchemaProps: openapispec.SchemaProps{
								Description: "CanaryTrafficPercent defines the percentage of traffic going to canary KFService endpoints",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
					Required: []string{"default"},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.EndpointSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.KFServiceStatus": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "KFServiceStatus defines the observed state of KFService",
					Properties: map[string]openapispec.Schema{
						"observedGeneration": {
							SchemaProps: openapispec.SchemaProps{
								Description: "ObservedGeneration is the 'Generation' of the Service that was last processed by the controller.",
								Type:        []string{"integer"},
								Format:      "int64",
							},
						},
						"conditions": {
							VendorExtensible: openapispec.VendorExtensible{
								Extensions: openapispec.Extensions{
									"x-kubernetes-patch-merge-key": "type",
									"x-kubernetes-patch-strategy":  "merge",
								},
							},
							SchemaProps: openapispec.SchemaProps{
								Description: "Conditions the latest available observations of a resource's current state.",
								Type:        []string{"array"},
								Items: &openapispec.SchemaOrArray{
									Schema: &openapispec.Schema{
										SchemaProps: openapispec.SchemaProps{
											Ref: ref("knative.dev/pkg/apis.Condition"),
										},
									},
								},
							},
						},
						"url": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"default": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec"),
							},
						},
						"canary": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec", "knative.dev/pkg/apis.Condition"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ONNXSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "ONNXSpec defines arguments for configuring ONNX model serving.",
					Properties: map[string]openapispec.Schema{
						"storageUri": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"runtimeVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to latest ONNX Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PredictorSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "PredictorSpec defines the configuration to route traffic to a predictor.",
					Properties: map[string]openapispec.Schema{
						"custom": {
							SchemaProps: openapispec.SchemaProps{
								Description: "The following fields follow a \"1-of\" semantic. Users must specify exactly one openapispec.",
								Ref:         ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"),
							},
						},
						"tensorflow": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorflowSpec"),
							},
						},
						"tensorrt": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorRTSpec"),
							},
						},
						"xgboost": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.XGBoostSpec"),
							},
						},
						"sklearn": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.SKLearnSpec"),
							},
						},
						"onnx": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ONNXSpec"),
							},
						},
						"pytorch": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PyTorchSpec"),
							},
						},
						"serviceAccountName": {
							SchemaProps: openapispec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.ONNXSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PyTorchSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.SKLearnSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorRTSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorflowSpec", "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.XGBoostSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.PyTorchSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "PyTorchSpec defines arguments for configuring PyTorch model serving.",
					Properties: map[string]openapispec.Schema{
						"storageUri": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"modelClassName": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults PyTorch model class name to 'PyTorchModel'",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"runtimeVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to latest PyTorch Version",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.SKLearnSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "SKLearnSpec defines arguments for configuring SKLearn model serving.",
					Properties: map[string]openapispec.Schema{
						"storageUri": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"runtimeVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to latest SKLearn Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.StatusConfigurationSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "StatusConfigurationSpec describes the state of the configuration receiving traffic.",
					Properties: map[string]openapispec.Schema{
						"name": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"replicas": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"integer"},
								Format: "int32",
							},
						},
						"traffic": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"integer"},
								Format: "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorRTSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "TensorRTSpec defines arguments for configuring TensorRT model serving.",
					Properties: map[string]openapispec.Schema{
						"storageUri": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"runtimeVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to latest TensorRT Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TensorflowSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "TensorflowSpec defines arguments for configuring Tensorflow model serving.",
					Properties: map[string]openapispec.Schema{
						"storageUri": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"runtimeVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to latest TF Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.TransformerSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "TransformerSpec defines transformer service for pre/post processing",
					Properties: map[string]openapispec.Schema{
						"custom": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"),
							},
						},
						"serviceAccountName": {
							SchemaProps: openapispec.SchemaProps{
								Description: "ServiceAccountName is the name of the ServiceAccount to use to run the service",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"minReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Minimum number of replicas, pods won't scale down to 0 in case of no traffic",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
						"maxReplicas": {
							SchemaProps: openapispec.SchemaProps{
								Description: "This is the up bound for autoscaler to scale to",
								Type:        []string{"integer"},
								Format:      "int32",
							},
						},
					},
				},
			},
			Dependencies: []string{
				"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.CustomSpec"},
		},
		"github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2.XGBoostSpec": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "XGBoostSpec defines arguments for configuring XGBoost model serving.",
					Properties: map[string]openapispec.Schema{
						"storageUri": {
							SchemaProps: openapispec.SchemaProps{
								Type:   []string{"string"},
								Format: "",
							},
						},
						"runtimeVersion": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to latest XGBoost Version.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"resources": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Defaults to requests and limits of 1CPU, 2Gb MEM.",
								Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
							},
						},
					},
					Required: []string{"storageUri"},
				},
			},
			Dependencies: []string{
				"k8s.io/api/core/v1.ResourceRequirements"},
		},
		"knative.dev/pkg/apis.Condition": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "Conditions defines a readiness condition for a Knative resource. See: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#typical-status-properties",
					Properties: map[string]openapispec.Schema{
						"type": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Type of condition.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"status": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Status of the condition, one of True, False, Unknown.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"severity": {
							SchemaProps: openapispec.SchemaProps{
								Description: "Severity with which to treat failures of this type of condition. When this is not specified, it defaults to Error.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"lastTransitionTime": {
							SchemaProps: openapispec.SchemaProps{
								Description: "LastTransitionTime is the last time the condition transitioned from one status to another. We use VolatileTime in place of metav1.Time to exclude this from creating equality.Semantic differences (all other things held constant).",
								Ref:         ref("knative.dev/pkg/apis.VolatileTime"),
							},
						},
						"reason": {
							SchemaProps: openapispec.SchemaProps{
								Description: "The reason for the condition's last transition.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"message": {
							SchemaProps: openapispec.SchemaProps{
								Description: "A human readable message indicating details about the transition.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
					},
					Required: []string{"type", "status"},
				},
			},
			Dependencies: []string{
				"knative.dev/pkg/apis.VolatileTime"},
		},
		"knative.dev/pkg/apis.VolatileTime": {
			Schema: openapispec.Schema{
				SchemaProps: openapispec.SchemaProps{
					Description: "VolatileTime wraps metav1.Time",
					Properties: map[string]openapispec.Schema{
						"Inner": {
							SchemaProps: openapispec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
							},
						},
					},
					Required: []string{"Inner"},
				},
			},
			Dependencies: []string{
				"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
		},
	}
}
