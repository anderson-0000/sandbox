package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewPrometheus(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
		Chart:     jsii.String("prometheus-community/prometheus"), //helm search repo prometheus
		Values: &map[string]interface{}{
			"server": map[string]interface{}{
				"ingress": map[string]interface{}{
					"enabled": true,
					"hosts": []string{
						"prometheus.remotehost",
					},
					"ingressClassName": "nginx",
				},
				"persistentVolume": map[string]bool{
					"enabled": false,
				},
			},
			"alertmanager": map[string]interface{}{
				"persistence": map[string]bool{
					"enabled": false,
				},
			},
		},
	})
	return chart
}
