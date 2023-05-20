package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2"

)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewNameSpace(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops) // idがファイル名になる

	// define resources here
	cdk8splus26.NewNamespace(chart, jsii.String("ns-ando"), &cdk8splus26.NamespaceProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("ando"),
		},
	})

	return chart
}

func NewDeploymentUbuntu(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
        var cprops cdk8s.ChartProps
        if props != nil {
                cprops = props.ChartProps
        }
        chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

        // define resources here
        cdk8splus26.NewDeployment(chart, jsii.String("Deployment"), &cdk8splus26.DeploymentProps{
        	Replicas: jsii.Number(3),
        	Containers: &[]*cdk8splus26.ContainerProps{{
        		Image: jsii.String("ubuntu"),
        	}},
        })

        return chart
}

func NewHelmArgocd(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
        var cprops cdk8s.ChartProps
        if props != nil {
                cprops = props.ChartProps
        }
        chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

        // define resources here
        cdk8s.NewHelm(chart, jsii.String("helm"), &cdk8s.HelmProps{
          	Chart: jsii.String("argo/argo-cd"), //helm repo add argo https://argoproj.github.io/argo-helm が必要
        })

        return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewNameSpace(app, "create-ns-ando", nil)
	NewDeploymentUbuntu(app, "ubuntu", nil)
	NewHelmArgocd(app, "argocd", nil)
	app.Synth()
}
