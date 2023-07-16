package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/k8s"
)

func NewArgoWorkflows(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
		Chart:     jsii.String("argo/argo-workflows"), //helm repo add argo https://argoproj.github.io/argo-helm
		Values: &map[string]interface{}{
			"nameOverride": id,
		},
	})

	k8s.NewKubeIngress(chart, jsii.String(id+"-ingress"), &k8s.KubeIngressProps{
		Metadata: &k8s.ObjectMeta{
			Name:      jsii.String(id),
			Namespace: jsii.String(id),
		},
		Spec: &k8s.IngressSpec{
			IngressClassName: jsii.String("nginx"),
			Rules: &[]*k8s.IngressRule{
				{
					Host: jsii.String(id + ".localhost"),
					Http: &k8s.HttpIngressRuleValue{
						Paths: &[]*k8s.HttpIngressPath{
							{
								Path:     jsii.String("/"),
								PathType: jsii.String("Prefix"),
								Backend: &k8s.IngressBackend{
									Service: &k8s.IngressServiceBackend{
										Name: jsii.String("argo-workflows-c8b79b5a-server"),
										Port: &k8s.ServiceBackendPort{
											Number: jsii.Number(2746),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	return chart
}
