package main

import (
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2"
)

type Config struct {
	IntegerValue  int     `yaml:"integerValue"`
	FloatingValue float64 `yaml:"floatingValue"`
	StringValue   string  `yaml:"stringValue"`
	BooleanValue  bool    `yaml:"booleanValue"`
}

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
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV 環境変数が設定されていません")
	}

	data, err := ioutil.ReadFile(env + ".yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("IntegerValue: %d\n", config.IntegerValue)
	log.Printf("FloatingValue: %f\n", config.FloatingValue)
	log.Printf("StringValue: %s\n", config.StringValue)
	log.Printf("BooleanValue: %t\n", config.BooleanValue)

	app := cdk8s.NewApp(nil)

	NewNameSpace(app, "ns", nil)
	NewDeploymentUbuntu(app, "ubuntu", nil)
	NewDeploymentCdk8sUbuntu(app, "ubuntu-cdk8s", nil)
	NewRoleBinding(app, "rolebinding", nil)
	NewArgocd(app, "argocd", nil)
	NewArgoWorkflows(app, "argo-workflows", nil)
	NewNginx(app, "nginx", nil)
	NewIngressNginx(app, "ingress-nginx-controller", nil)
	NewCertManager(app, "cert-manager", nil)
	NewDeploymentGeth(app, "geth", nil)
	NewPvAndPvc(app, "pv-and-pvc", nil)
	NewPrometheus(app, "prometheus", nil)
	NewGrafana(app, "grafana", nil)
	app.Synth()
}
