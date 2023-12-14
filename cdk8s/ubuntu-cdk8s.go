package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2"
)

func NewDeploymentCdk8sUbuntu(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	c := cdk8s.NewChart(scope, jsii.String(id), &cprops) // idがファイル名になる

	ubuntu := cdk8splus26.NewDeployment(c, jsii.String("ubuntu"), &cdk8splus26.DeploymentProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String("ubuntu-cdk8s"),
			Namespace: jsii.String("ubuntu"),
		},
		Replicas: jsii.Number(1),
		Containers: &[]*cdk8splus26.ContainerProps{
			{
				Image: jsii.String("ubuntu"),
				Command: &[]*string{
					jsii.String("/bin/bash"),
				},
				Name: jsii.String("main"),
				SecurityContext: &cdk8splus26.ContainerSecurityContextProps{
					EnsureNonRoot:          jsii.Bool(false),
					ReadOnlyRootFilesystem: jsii.Bool(false),
				},
			},
		},
	})

	//	ubuntu.ApiObject().AddJsonPatch(cdk8s.JsonPatch_Add(jsii.String("/spec/template/spec/containers/0/stdin"), true))
	//	ubuntu.ApiObject().AddJsonPatch(cdk8s.JsonPatch_Add(jsii.String("/spec/template/spec/containers/0/tty"), true))
	// ↑↓は同じ

	ubuntu.ApiObject().AddJsonPatch(
		cdk8s.JsonPatch_Add(jsii.String("/spec/template/spec/containers/0/stdin"), true),
		cdk8s.JsonPatch_Add(jsii.String("/spec/template/spec/containers/0/tty"), true),
	)
	return c
}

