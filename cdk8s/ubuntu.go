package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/k8s"
)

func NewDeploymentUbuntu(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	c := cdk8s.NewChart(scope, jsii.String(id), &cprops) // idがファイル名になる

	k8s.NewKubeDeployment(c, jsii.String("ubuntu"), &k8s.KubeDeploymentProps{
		Metadata: &k8s.ObjectMeta{
			Name:      jsii.String("ubuntu"),
			Namespace: jsii.String("ubuntu"),
		},
		Spec: &k8s.DeploymentSpec{
			MinReadySeconds:         jsii.Number(0),
			ProgressDeadlineSeconds: jsii.Number(600),
			Replicas:                jsii.Number(1),
			Selector: &k8s.LabelSelector{
				MatchLabels: &map[string]*string{
					"cdk8s.io/metadata.addr": jsii.String("ubuntu-deployment"),
				},
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: &map[string]*string{
						"cdk8s.io/metadata.addr": jsii.String("ubuntu-deployment"),
					},
				},
				Spec: &k8s.PodSpec{
					AutomountServiceAccountToken: jsii.Bool(false),
					DnsPolicy:                    jsii.String("ClusterFirst"),
					SetHostnameAsFqdn:            jsii.Bool(false),
					SecurityContext: &k8s.PodSecurityContext{
						FsGroupChangePolicy: jsii.String("Always"),
						RunAsNonRoot:        jsii.Bool(false),
					},
					RestartPolicy: jsii.String("Always"),
					HostNetwork:   jsii.Bool(false),
					Containers: &[]*k8s.Container{{
						Command: &[]*string{
							jsii.String("/bin/bash"),
						},
						ImagePullPolicy: jsii.String("Always"),
						Name:            jsii.String("main"),
						Tty:             jsii.Bool(true),
						Image:           jsii.String("ubuntu"),
						SecurityContext: &k8s.SecurityContext{
							AllowPrivilegeEscalation: jsii.Bool(false),
							Privileged:               jsii.Bool(false),
							ReadOnlyRootFilesystem:   jsii.Bool(true),
							RunAsNonRoot:             jsii.Bool(false),
						},
						Stdin: jsii.Bool(true),
					}},
				},
			},
		},
	})

	return c
}
