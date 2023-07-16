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

func NewRoleBinding(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	roles := []struct {
		name string
		role string
		sa   string
	}{
		{"read-only", "read-only", "sa-read-only"},
		{"admin", "admin", "sa-admin"},
	}

	for _, r := range roles {
		role := cdk8splus26.NewRole(chart, jsii.String(r.name), &cdk8splus26.RoleProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name: jsii.String(r.role),
			},
		})

		cdk8splus26.NewRoleBinding(chart, jsii.String(r.sa), &cdk8splus26.RoleBindingProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name: jsii.String(r.role),
			},
			Role: role,
		})
	}

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewNameSpace(app, "ns", nil)
	NewDeploymentUbuntu(app, "ubuntu", nil)
	NewRoleBinding(app, "rolebinding", nil)
	NewArgocd(app, "argocd", nil)
	NewArgoWorkflows(app, "argo-workflows", nil)
	NewNginx(app, "nginx", nil)
	NewIngressNginx(app, "ingress-nginx-controller", nil)
	NewCertManager(app, "cert-manager", nil)
	app.Synth()
}
