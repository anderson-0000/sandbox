package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewIngressNginx(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
		Chart:     jsii.String("ingress-nginx/ingress-nginx"), //helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	})

	return chart
}
