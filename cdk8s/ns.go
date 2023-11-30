package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2"
)

func NewNameSpace(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops) // idがファイル名になる

	// define resources here

	namespaces := []string{
		"testtesttest",
		"argocd",
		"argo-workflows",
		"ubuntu",
		"dex",
		"ingress-nginx-controller",
		"cert-manager",
		"geth",
		"ci",
		"prometheus",
		"grafana",
	}

	for _, ns := range namespaces {
		cdk8splus26.NewNamespace(chart, jsii.String("ns-"+ns), &cdk8splus26.NamespaceProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name: jsii.String(ns),
			},
		})
	}

	return chart
}
