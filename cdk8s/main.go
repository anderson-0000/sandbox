package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/k8s"
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

	namespaces := []string{"ando", "argocd", "argo-workflows", "ubuntu"}

	for _, ns := range namespaces {
		cdk8splus26.NewNamespace(chart, jsii.String("ns-" + ns), &cdk8splus26.NamespaceProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name: jsii.String(ns),
			},
		})
	}

	return chart
}

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
			MinReadySeconds:        jsii.Number(0),
			ProgressDeadlineSeconds: jsii.Number(600),
			Replicas:               jsii.Number(1),
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

func NewRoleBinding(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	role := cdk8splus26.NewRole(chart, jsii.String("read-only"), &cdk8splus26.RoleProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("read-only"),
		},
	})

	cdk8splus26.NewRoleBinding(chart, jsii.String("sa-read"), &cdk8splus26.RoleBindingProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("read"),
		},
		Role: role,
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
	dexConfig := `connectors:
	- type: github
		id: github
		name: GitHub
		config:
			clientID: Iv1.441a977aa92f84fb
			clientSecret: $dex.github.clientSecret # Alternatively $<some_K8S_secret>:dex.github.clientSecret
			orgs:
			- name: your-github-org`

	cdk8s.NewHelm(chart, jsii.String("helm"), &cdk8s.HelmProps{
		Chart: jsii.String("argo/argo-cd"), //helm repo add argo https://argoproj.github.io/argo-helm が必要
		Values: &map[string]interface{}{
			"configs": map[string]interface{}{
				"cm": map[string]interface{}{
					"create":                       true,
					"application.instanceLabelKey": "argocd.argoproj.io/instance",
					"dex.config":                   dexConfig,
				},
			},
		},
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewNameSpace(app, "ns", nil)
	NewDeploymentUbuntu(app, "ubuntu", nil)
	NewRoleBinding(app, "rolebinding", nil)
	NewHelmArgocd(app, "argocd", nil)
	app.Synth()
}
