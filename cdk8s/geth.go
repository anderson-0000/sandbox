package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2"
)

func NewDeploymentGeth(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	// localnet
	replicas := float64(1.0)

	cdk8splus26.NewDeployment(chart, jsii.String(id), &cdk8splus26.DeploymentProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String(id),
		},
		Replicas: &replicas,
		Containers: &[]*cdk8splus26.ContainerProps{
			{
				Name:  jsii.String(id),
				Image: jsii.String("ethereum/client-go:latest"),
				Args: &[]*string{
					jsii.String("--ipcdisable"),
					jsii.String("--rpc"),
				},
			},
		},
	})

	return chart
}
