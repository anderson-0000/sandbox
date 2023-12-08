package main

//https://qiita.com/showchan33/items/dc57030f8794ca4a7c31 を参考に作る
import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewGrafana(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
		Chart:     jsii.String("bitnami/grafana"), //helm search repo grafana
		Values: &map[string]interface{}{
			"datasources": map[string]interface{}{
				"datasources.yaml": map[string]interface{}{
					"apiVersion": 1,
					"datasources": map[string]interface{}{
						"name":  "Prometheus",
						"type":   "prometheus",
						"access": "proxy",
						"url":    "http://prometheus-c899c4b6-server.prometheus.svc.cluster.local",
					},
				},
			},
			"dashboardProviders": map[string]interface{}{
				"dashboardproviders.yaml": map[string]interface{}{
					"apiVersion": 1,
					"providers": map[string]interface{}{
						"name":            "default",
						"orgId":           1,
						"folder":          "",
						"type":            "file",
						"disableDeletion": false,
						"editable":        true,
						"options": map[string]interface{}{
							"path": "/var/lib/grafana/dashboards/default",
						},
					},
				},
			},
			"dashboards": map[string]interface{}{
				"default": map[string]interface{}{
					"node-exporter": map[string]interface{}{
						"datasource": "Prometheus",
						"url":        "https://grafana.com/api/dashboards/1860/revisions/32/download", // https://grafana.com/grafana/dashboards/1860-node-exporter-full/
					},
				},
			},
			"ingress": map[string]interface{}{
				"enabled": true,
				"hosts": []string{
					"grafanaaaa.remotehost",
				},
				"ingressClassName": "nginx",
			},
		},
	})
	return chart
}
