package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewArgocd(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	dexConfig := `connectors:
	- type: github
		id: github
		name: GitHub
		config:
			clientID: Iv1.441a977aa92f84fb
			clientSecret: $dex.github.clientSecret # Alternatively $<some_K8S_secret>:dex.github.clientSecret
			orgs:
			- name: your-github-org`

	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
		Chart:     jsii.String("argo/argo-cd"), //helm repo add argo https://argoproj.github.io/argo-helm が必要
		Values: &map[string]interface{}{
			"configs": map[string]interface{}{
				"cm": map[string]interface{}{
					"create":                       true,
					"application.instanceLabelKey": "argocd.argoproj.io/instance",
					"dex.config":                   dexConfig,
				},
				"params": map[string]interface{}{
					"server.insecure": "true",
				},
			},
			"server": map[string]interface{}{
				"ingress": map[string]interface{}{
					"enabled": "true",
					"hosts": []string{
						"argocd.remotehost",
					},
					"ingressClassName": "nginx",
				},
				"ingressGrpc": map[string]interface{}{
					"enabled": "true",
					"hosts": []string{
						"grpc.argocd.remotehost",
					},
				},
			},
		},
	})

	return chart
}
