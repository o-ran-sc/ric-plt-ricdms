package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetConfigMap() []*corev1.ConfigMap {

	configMap1 := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		Data: map[string]string{
			"schema.json": "{\n" +
				"    \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n" +
				"    \"$id\": \"http://o-ran-sc.org/xapp_root.json\",\n" +
				"    \"type\": \"object\",\n" +
				"    \"title\": \"The xApp Root Schema\",\n" +
				"    \"properties\": {\n" +
				"        \"name\": {\n" +
				"            \"$id\": \"#/properties/name\",\n" +
				"            \"type\": \"string\",\n" +
				"            \"title\": \"The xApp Name\",\n" +
				"            \"default\": \"xapp\",\n" +
				"            \"examples\": [\n" +
				"                \"example_xapp\"\n" +
				"            ]\n" +
				"        },\n" +
				"        \"version\": {\n" +
				"            \"$id\": \"#/properties/version\",\n" +
				"            \"type\": \"string\",\n" +
				"            \"title\": \"The xApp version\",\n" +
				"            \"default\": \"1.0.0\",\n" +
				"            \"examples\": [\n" +
				"                \"1.0.0\"\n" +
				"            ],\n" +
				"            \"pattern\": \"^(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)(?:-((?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\\\.(?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\\\+([0-9a-zA-Z-]+(?:\\\\.[0-9a-zA-Z-]+)*))?$\"\n" +
				"        },\n" +
				"        \"annotations\": {\n" +
				"            \"$id\": \"#/properties/annotation\",\n" +
				"            \"type\": \"object\",\n" +
				"            \"title\": \"The k8s pod annotation\",\n" +
				"            \"additionalProperties\": {\n" +
				"                \"anyOf\": [\n" +
				"                    {\n" +
				"                        \"type\": \"string\"\n" +
				"                    },\n" +
				"                    {\n" +
				"                        \"type\": \"array\",\n" +
				"                        \"items\": {\n" +
				"                            \"type\": \"object\"\n" +
				"                        }\n" +
				"                    }\n" +
				"                ]\n" +
				"            }\n" +
				"        },\n" +
				"        \"containers\": {\n" +
				"            \"$id\": \"#/properties/containers\",\n" +
				"            \"type\": \"array\",\n" +
				"            \"title\": \"The Container Schema\",\n" +
				"            \"items\": {\n" +
				"                \"$id\": \"#/properties/containers/items\",\n" +
				"                \"type\": \"object\",\n" +
				"                \"title\": \"The Container Items Schema\",\n" +
				"                \"required\": [\n" +
				"                    \"name\",\n" +
				"                    \"image\"\n" +
				"                ],\n" +
				"                \"properties\": {\n" +
				"                    \"name\": {\n" +
				"                        \"$id\": \"#/properties/containers/items/properties/name\",\n" +
				"                        \"type\": \"string\",\n" +
				"                        \"title\": \"The xApp Container Name\",\n" +
				"                        \"default\": \"xapp\",\n" +
				"                        \"examples\": [\n" +
				"                            \"xapp\"\n" +
				"                        ]\n" +
				"                    },\n" +
				"                    \"image\": {\n" +
				"                        \"$id\": \"#/properties/containers/items/properties/image\",\n" +
				"                        \"type\": \"object\",\n" +
				"                        \"title\": \"The Container Image\",\n" +
				"                        \"required\": [\n" +
				"                            \"registry\",\n" +
				"                            \"name\",\n" +
				"                            \"tag\"\n" +
				"                        ],\n" +
				"                        \"properties\": {\n" +
				"                            \"registry\": {\n" +
				"                                \"$id\": \"#/properties/containers/items/properties/image/properties/registry\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The xApp Image Registry\",\n" +
				"                                \"default\": \"nexus3.o-ran-sc.org:10002\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"nexus3.o-ran-sc.org:10002\"\n" +
				"                                ],\n" +
				"                                \"pattern\": \"^([A-Za-z0-9\\\\.-]{1,}\\\\.[A-Za-z]{1,}|((?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)))(?:\\\\:\\\\d+)?$\"\n" +
				"                            },\n" +
				"                            \"name\": {\n" +
				"                                \"$id\": \"#/properties/containers/items/properties/image/properties/name\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The xApp Image Name\",\n" +
				"                                \"default\": \"xapp\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"xapp\"\n" +
				"                                ]\n" +
				"                            },\n" +
				"                            \"tag\": {\n" +
				"                                \"$id\": \"#/properties/containers/items/properties/image/properties/tag\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The xApp Image Tag\",\n" +
				"                                \"default\": \"latest\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"latest\"\n" +
				"                                ]\n" +
				"                            }\n" +
				"                        }\n" +
				"                    },\n" +
				"                    \"command\": {\n" +
				"                        \"$id\": \"#/properties/containers/items/properties/command\",\n" +
				"                        \"type\": \"array\",\n" +
				"                        \"items\": [\n" +
				"                            {\n" +
				"                                \"$id\": \"#/properties/containers/items/properties/command/item\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The Command Item\",\n" +
				"                                \"default\": \"/bin/sh\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"/bin/sh\"\n" +
				"                                ]\n" +
				"                            }\n" +
				"                        ]\n" +
				"                    },\n" +
				"                    \"args\": {\n" +
				"                        \"$id\": \"#/properties/containers/items/properties/args\",\n" +
				"                        \"type\": \"array\",\n" +
				"                        \"items\": [\n" +
				"                            {\n" +
				"                                \"$id\": \"#/properties/containers/items/properties/args/item\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The Command Arguement Item\",\n" +
				"                                \"default\": \"-c\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"-c\"\n" +
				"                                ]\n" +
				"                            }\n" +
				"                        ]\n" +
				"                    }\n" +
				"                }\n" +
				"            }\n" +
				"        },\n" +
				"        \"livenessProbe\": {\n" +
				"            \"$id\": \"#/properties/livenessprobe\",\n" +
				"            \"type\": \"object\",\n" +
				"            \"title\": \"The Liveness Probe Definition\",\n" +
				"            \"properties\": {\n" +
				"                \"exec\": {\n" +
				"                    \"$id\": \"#/properties/livenessprobe/exec\",\n" +
				"                    \"type\": \"object\",\n" +
				"                    \"title\": \"Script of Liveness Probe\",\n" +
				"                    \"properties\": {\n" +
				"                        \"command\": {\n" +
				"                            \"$id\": \"#/properties/livenessprobe/exec/command\",\n" +
				"                            \"type\": \"array\",\n" +
				"                            \"items\": [\n" +
				"                                {\n" +
				"                                    \"$id\": \"#/properties/livenessprobe/exec/command/item\",\n" +
				"                                    \"type\": \"string\",\n" +
				"                                    \"title\": \"The Command Item\",\n" +
				"                                    \"default\": \"/bin/sh\",\n" +
				"                                    \"examples\": [\n" +
				"                                        \"/bin/sh\"\n" +
				"                                    ]\n" +
				"                                }\n" +
				"                            ]\n" +
				"                        }\n" +
				"                    },\n" +
				"                    \"required\": [\n" +
				"                        \"command\"\n" +
				"                    ]\n" +
				"                },\n" +
				"                \"httpGet\": {\n" +
				"                    \"$id\": \"#/properties/livenessprobe/httpget\",\n" +
				"                    \"type\": \"object\",\n" +
				"                    \"title\": \"Http of Liveness Probe\",\n" +
				"                    \"properties\": {\n" +
				"                        \"path\": {\n" +
				"                            \"$id\": \"#/properties/livenessprobe/httpget/path\",\n" +
				"                            \"type\": \"string\",\n" +
				"                            \"title\": \"The Path of Http Liveness Probe\",\n" +
				"                            \"default\": \"/health\",\n" +
				"                            \"examples\": [\n" +
				"                                \"/health\"\n" +
				"                            ]\n" +
				"                        },\n" +
				"                        \"port\": {\n" +
				"                            \"$id\": \"#/properties/livenessprobe/httpget/port\",\n" +
				"                            \"type\": \"integer\",\n" +
				"                            \"title\": \"The Port of Http Liveness Probe\",\n" +
				"                            \"default\": 80,\n" +
				"                            \"examples\": [\n" +
				"                                80\n" +
				"                            ]\n" +
				"                        }\n" +
				"                    },\n" +
				"                    \"required\": [\n" +
				"                        \"path\",\n" +
				"                        \"port\"\n" +
				"                    ]\n" +
				"                },\n" +
				"                \"initialDelaySeconds\": {\n" +
				"                    \"$id\": \"#/properties/livenessprobe/initialdelayseconds\",\n" +
				"                    \"type\": \"integer\",\n" +
				"                    \"title\": \"Initial Delay of Liveness Probe\",\n" +
				"                    \"default\": 5,\n" +
				"                    \"examples\": [\n" +
				"                        5\n" +
				"                    ]\n" +
				"                },\n" +
				"                \"periodSeconds\": {\n" +
				"                    \"$id\": \"#/properties/livenessprobe/periodseconds\",\n" +
				"                    \"type\": \"integer\",\n" +
				"                    \"title\": \"Period of Liveness Probe\",\n" +
				"                    \"default\": 15,\n" +
				"                    \"examples\": [\n" +
				"                        15\n" +
				"                    ]\n" +
				"                }\n" +
				"            },\n" +
				"            \"oneOf\": [\n" +
				"                {\n" +
				"                    \"$id\": \"#/properties/livenessprobe/oneof/exec\",\n" +
				"                    \"required\": [\n" +
				"                        \"exec\",\n" +
				"                        \"initialDelaySeconds\",\n" +
				"                        \"periodSeconds\"\n" +
				"                    ]\n" +
				"                },\n" +
				"                {\n" +
				"                    \"$id\": \"#/properties/livenessprobe/oneof/httpget\",\n" +
				"                    \"required\": [\n" +
				"                        \"httpGet\",\n" +
				"                        \"initialDelaySeconds\",\n" +
				"                        \"periodSeconds\"\n" +
				"                    ]\n" +
				"                }\n" +
				"            ]\n" +
				"        },\n" +
				"        \"readinessProbe\": {\n" +
				"            \"$id\": \"#/properties/readinessprobe\",\n" +
				"            \"type\": \"object\",\n" +
				"            \"title\": \"The Readiness Probe Definition\",\n" +
				"            \"properties\": {\n" +
				"                \"exec\": {\n" +
				"                    \"$id\": \"#/properties/readinessprobe/exec\",\n" +
				"                    \"type\": \"object\",\n" +
				"                    \"title\": \"Script of Readiness Probe\",\n" +
				"                    \"properties\": {\n" +
				"                        \"command\": {\n" +
				"                            \"$id\": \"#/properties/readinessprobe/exec/command\",\n" +
				"                            \"type\": \"array\",\n" +
				"                            \"items\": [\n" +
				"                                {\n" +
				"                                    \"type\": \"string\"\n" +
				"                                }\n" +
				"                            ]\n" +
				"                        }\n" +
				"                    },\n" +
				"                    \"required\": [\n" +
				"                        \"command\"\n" +
				"                    ]\n" +
				"                },\n" +
				"                \"httpGet\": {\n" +
				"                    \"$id\": \"#/properties/readinessprobe/httpget\",\n" +
				"                    \"type\": \"object\",\n" +
				"                    \"title\": \"Http of Readiness Probe\",\n" +
				"                    \"properties\": {\n" +
				"                        \"path\": {\n" +
				"                            \"$id\": \"#/properties/readinessprobe/httpget/path\",\n" +
				"                            \"type\": \"string\",\n" +
				"                            \"title\": \"The Path of Http Readiness Probe\",\n" +
				"                            \"default\": \"/health\",\n" +
				"                            \"examples\": [\n" +
				"                                \"/health\"\n" +
				"                            ]\n" +
				"                        },\n" +
				"                        \"port\": {\n" +
				"                            \"$id\": \"#/properties/readinessprobe/httpget/port\",\n" +
				"                            \"type\": \"integer\",\n" +
				"                            \"title\": \"The Port of Http Readiness Probe\",\n" +
				"                            \"default\": 80,\n" +
				"                            \"examples\": [\n" +
				"                                80\n" +
				"                            ]\n" +
				"                        }\n" +
				"                    },\n" +
				"                    \"required\": [\n" +
				"                        \"path\",\n" +
				"                        \"port\"\n" +
				"                    ]\n" +
				"                },\n" +
				"                \"initialDelaySeconds\": {\n" +
				"                    \"$id\": \"#/properties/readinessprobe/initialdelayseconds\",\n" +
				"                    \"type\": \"integer\",\n" +
				"                    \"title\": \"Initial Delay of Readiness Probe\",\n" +
				"                    \"default\": 5,\n" +
				"                    \"examples\": [\n" +
				"                        5\n" +
				"                    ]\n" +
				"                },\n" +
				"                \"periodSeconds\": {\n" +
				"                    \"$id\": \"#/properties/readinessprobe/periodseconds\",\n" +
				"                    \"type\": \"integer\",\n" +
				"                    \"title\": \"Period of Readiness Probe\",\n" +
				"                    \"default\": 15,\n" +
				"                    \"examples\": [\n" +
				"                        15\n" +
				"                    ]\n" +
				"                }\n" +
				"            },\n" +
				"            \"oneOf\": [\n" +
				"                {\n" +
				"                    \"$id\": \"#/properties/readinessprobe/oneof/exec\",\n" +
				"                    \"required\": [\n" +
				"                        \"exec\",\n" +
				"                        \"initialDelaySeconds\",\n" +
				"                        \"periodSeconds\"\n" +
				"                    ]\n" +
				"                },\n" +
				"                {\n" +
				"                    \"$id\": \"#/properties/readinessprobe/oneof/httpget\",\n" +
				"                    \"required\": [\n" +
				"                        \"httpGet\",\n" +
				"                        \"initialDelaySeconds\",\n" +
				"                        \"periodSeconds\"\n" +
				"                    ]\n" +
				"                }\n" +
				"            ]\n" +
				"        },\n" +
				"        \"messaging\": {\n" +
				"            \"type\": \"object\",\n" +
				"            \"$id\": \"#/properties/messaging\",\n" +
				"            \"title\": \"The Messaging Schema\",\n" +
				"            \"properties\": {\n" +
				"                \"ports\": {\n" +
				"                    \"$id\": \"#/properties/messaging/ports\",\n" +
				"                    \"type\": \"array\",\n" +
				"                    \"title\": \"The Ports for Messaging\",\n" +
				"                    \"items\": {\n" +
				"                        \"$id\": \"#/properties/messaging/ports/items\",\n" +
				"                        \"type\": \"object\",\n" +
				"                        \"title\": \"The Item of Port\",\n" +
				"                        \"required\": [\n" +
				"                            \"name\",\n" +
				"                            \"container\",\n" +
				"                            \"port\"\n" +
				"                        ],\n" +
				"                        \"dependencies\": {\n" +
				"                            \"txMessages\": [\n" +
				"                                \"rxMessages\",\n" +
				"                                \"policies\"\n" +
				"                            ],\n" +
				"                            \"rxMessages\": [\n" +
				"                                \"txMessages\",\n" +
				"                                \"policies\"\n" +
				"                            ],\n" +
				"                            \"policies\": [\n" +
				"                                \"rxMessages\",\n" +
				"                                \"txMessages\"\n" +
				"                            ]\n" +
				"                        },\n" +
				"                        \"properties\": {\n" +
				"                            \"name\": {\n" +
				"                                \"$id\": \"#/properties/messaging/ports/items/name\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The Name of the Port\",\n" +
				"                                \"default\": \"App\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"App\"\n" +
				"                                ]\n" +
				"                            },\n" +
				"                            \"container\": {\n" +
				"                                \"$id\": \"#/properties/messaging/ports/items/container\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The Container of the Port\",\n" +
				"                                \"default\": \"xapp\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"xapp\"\n" +
				"                                ]\n" +
				"                            },\n" +
				"                            \"port\": {\n" +
				"                                \"$id\": \"#/properties/messaging/ports/items/port\",\n" +
				"                                \"type\": \"integer\",\n" +
				"                                \"title\": \"The Port Number\",\n" +
				"                                \"default\": 8080,\n" +
				"                                \"examples\": [\n" +
				"                                    8080\n" +
				"                                ]\n" +
				"                            },\n" +
				"                            \"description\": {\n" +
				"                                \"$id\": \"#/properties/messaging/ports/items/description\",\n" +
				"                                \"type\": \"string\",\n" +
				"                                \"title\": \"The description for the port\",\n" +
				"                                \"default\": \"port description\",\n" +
				"                                \"examples\": [\n" +
				"                                    \"port description\"\n" +
				"                                ]\n" +
				"                            },\n" +
				"                            \"txMessages\": {\n" +
				"                                \"$id\": \"#/properties/messaging/ports/items/txmessages\",\n" +
				"                                \"type\": \"array\",\n" +
				"                                \"title\": \"The txMessage Types\",\n" +
				"                                \"items\": {\n" +
				"                                    \"$id\": \"#/properties/messaging/ports/items//txmessages/item\",\n" +
				"                                    \"type\": \"string\",\n" +
				"                                    \"title\": \"The txMessage Types Item\",\n" +
				"                                    \"default\": \"RIC_SUB\",\n" +
				"                                    \"examples\": [\n" +
				"                                        \"RIC_SUB\"\n" +
				"                                    ]\n" +
				"                                }\n" +
				"                            },\n" +
				"                            \"rxMessages\": {\n" +
				"                                \"$id\": \"#/properties/messaging/ports/items/rxmessages\",\n" +
				"                                \"type\": \"array\",\n" +
				"                                \"title\": \"The rxMessage Types\",\n" +
				"                                \"items\": {\n" +
				"                                    \"$id\": \"#/properties/messaging/ports/items/rxmessages/item\",\n" +
				"                                    \"type\": \"string\",\n" +
				"                                    \"title\": \"The rxMessage Types Item\",\n" +
				"                                    \"default\": \"RIC_SUB\",\n" +
				"                                    \"examples\": [\n" +
				"                                        \"RIC_SUB\"\n" +
				"                                    ]\n" +
				"                                }\n" +
				"                            },\n" +
				"                            \"policies\": {\n" +
				"                                \"$id\": \"#/properties/messaging/ports/items/policies\",\n" +
				"                                \"type\": \"array\",\n" +
				"                                \"title\": \"The Policies Types\",\n" +
				"                                \"items\": {\n" +
				"                                    \"$id\": \"#/properties/messaging/ports/items/policies/item\",\n" +
				"                                    \"type\": \"integer\",\n" +
				"                                    \"title\": \"The Policy Types Item\",\n" +
				"                                    \"default\": 1,\n" +
				"                                    \"examples\": [\n" +
				"                                        1\n" +
				"                                    ]\n" +
				"                                }\n" +
				"                            }\n" +
				"                        }\n" +
				"                    }\n" +
				"                }\n" +
				"            },\n" +
				"            \"required\": [\n" +
				"                \"ports\"\n" +
				"            ]\n" +
				"        },\n" +
				"        \"metrics\": {\n" +
				"            \"type\": \"object\",\n" +
				"            \"$id\": \"#/properties/metrics\",\n" +
				"            \"title\": \"The Metrics Schema\",\n" +
				"            \"items\": {\n" +
				"                \"$id\": \"#/properties/metrics/items\",\n" +
				"                \"type\": \"object\",\n" +
				"                \"title\": \"The Metrics Items Schema\",\n" +
				"                \"required\": [\n" +
				"                    \"objectName\",\n" +
				"                    \"objectInstance\",\n" +
				"                    \"name\",\n" +
				"                    \"type\",\n" +
				"                    \"description\"\n" +
				"                ],\n" +
				"                \"properties\": {\n" +
				"                    \"objectName\": {\n" +
				"                        \"$id\": \"#/properties/metrics/items/objectname\",\n" +
				"                        \"type\": \"string\",\n" +
				"                        \"title\": \"The Object Name\"\n" +
				"                    },\n" +
				"                    \"objectInstance\": {\n" +
				"                        \"$id\": \"#/properties/metrics/items/objectinstance\",\n" +
				"                        \"type\": \"string\",\n" +
				"                        \"title\": \"The Object Instance\"\n" +
				"                    },\n" +
				"                    \"name\": {\n" +
				"                        \"$id\": \"#/properties/metrics/items/name\",\n" +
				"                        \"type\": \"string\",\n" +
				"                        \"title\": \"The Object Name\"\n" +
				"                    },\n" +
				"                    \"type\": {\n" +
				"                        \"$id\": \"#/properties/metrics/items/type\",\n" +
				"                        \"type\": \"string\",\n" +
				"                        \"title\": \"The Object Type\"\n" +
				"                    },\n" +
				"                    \"description\": {\n" +
				"                        \"$id\": \"#/properties/metrics/items/description\",\n" +
				"                        \"type\": \"string\",\n" +
				"                        \"title\": \"The Object Description\"\n" +
				"                    }\n" +
				"                }\n" +
				"            }\n" +
				"        },\n" +
				"        \"controls\": {\n" +
				"            \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n" +
				"            \"$id\": \"#/controls\",\n" +
				"            \"type\": \"object\",\n" +
				"            \"title\": \"Controls Section Schema\",\n" +
				"            \"required\": [],\n" +
				"            \"properties\": {}\n" +
				"        }\n" +
				"    }\n" +
				"}\n" +
				"",
			"config-file.json": "{\n" +
				"    \"name\": \"hw-go\",\n" +
				"    \"version\": \"1.0.0\",\n" +
				"    \"containers\": [{\"image\":{\"name\":\"o-ran-sc/ric-app-hw-go\",\"registry\":\"nexus3.o-ran-sc.org:10004\",\"tag\":\"1.1.1\"},\"name\":\"hw-go\"}],\n" +
				"    \"livenessProbe\": {\n" +
				"        \"httpGet\": {\n" +
				"            \"path\": \"ric/v1/health/alive\",\n" +
				"            \"port\": 8080\n" +
				"        },\n" +
				"        \"initialDelaySeconds\": 5,\n" +
				"        \"periodSeconds\": 15\n" +
				"    },\n" +
				"    \"readinessProbe\": {\n" +
				"        \"httpGet\": {\n" +
				"            \"path\": \"ric/v1/health/ready\",\n" +
				"            \"port\": 8080\n" +
				"        },\n" +
				"        \"initialDelaySeconds\": 5,\n" +
				"        \"periodSeconds\": 15\n" +
				"    },\n" +
				"    \"messaging\": {\n" +
				"        \"ports\": [{\"container\":\"hw-go\",\"description\":\"http service\",\"name\":\"http\",\"port\":8080},{\"container\":\"hw-go\",\"description\":\"rmr route port for hw-go xapp\",\"name\":\"rmrroute\",\"port\":4561},{\"container\":\"hw-go\",\"description\":\"rmr data port for hw-go\",\"mtypes\":[{\"id\":55555,\"name\":\"TESTNAME1\"},{\"id\":55556,\"name\":\"TESTNAME2\"}],\"name\":\"rmrdata\",\"policies\":[1],\"port\":4560,\"rxMessages\":[\"RIC_SUB_RESP\",\"RIC_SUB_FAILURE\",\"RIC_SUB_DEL_RESP\",\"RIC_INDICATION\"],\"txMessages\":[\"RIC_SUB_REQ\",\"RIC_SUB_DEL_REQ\",\"RIC_SGNB_ADDITION_REQ\",\"RIC_SGNB_ADDITION_ACK\"]}]\n" +
				"    },\n" +
				"    \"rmr\": {\n" +
				"        \"protPort\": \"tcp:4560\",\n" +
				"        \"maxSize\": 2072,\n" +
				"        \"numWorkers\": 1,\n" +
				"        \"txMessages\": [\"RIC_SUB_REQ\",\"A1_POLICY_RESP\",\"A1_POLICY_QUERY\",\"RIC_HEALTH_CHECK_RESP\"],\n" +
				"        \"rxMessages\": [\"RIC_SUB_RESP\",\"A1_POLICY_REQ\",\"RIC_HEALTH_CHECK_REQ\"],\n" +
				"        \"policies\": [1]\n" +
				"    },\n" +
				"    \"controls\": {\n" +
				"        \"fileStrorage\": false,\n" +
				"        \"logger\": {\n" +
				"            \"level\": 3\n" +
				"        }\n" +
				"    },\n" +
				"    \"db\": {\n" +
				"        \"waitForSdl\": false\n" +
				"    }\n" +
				"}\n" +
				"",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-ricplt-hw-go-appconfig",
		},
	}

	configMap2 := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-ricplt-hw-go-appenv",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: map[string]string{
			"XAPP_DESCRIPTOR_PATH":     "/opt/ric/config",
			"DBAAS_PORT_6379_TCP_ADDR": "service-ricplt-dbaas-tcp.ricplt.svc.cluster.local",
			"RMR_SRC_ID":               "service-ricplt-hw-go-rmr.ricplt",
			"DBAAS_SERVICE_PORT":       "6379",
			"RMR_RTG_SVC":              "4561",
			"SERVICE_METRICSDB_HOST":   "service-metricsdb.ricplt.svc.cluster.local",
			"SERVICE_METRICSDB_PORT":   "8086",
			"DBAAS_PORT_6379_TCP_PORT": "6379",
			"DBAAS_SERVICE_HOST":       "service-ricplt-dbaas-tcp.ricplt.svc.cluster.local",
		},
	}

	return []*corev1.ConfigMap{configMap1, configMap2}
}
