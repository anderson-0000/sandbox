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

	namespaces := []string{"testtesttest", "argocd", "argo-workflows", "ubuntu", "dex", "ingress-nginx-controller", "cert-manager"}

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

func NewArgocd(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
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

	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
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

func NewArgoWorkflows(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
		Chart: jsii.String("argo/argo-workflows"), //helm repo add argo https://argoproj.github.io/argo-helm
		Values: &map[string]interface{}{
			"nameOverride": id,
		},
	})

	k8s.NewKubeIngress(chart, jsii.String(id + "-ingress"), &k8s.KubeIngressProps{
	    Metadata: &k8s.ObjectMeta{
	        Name: jsii.String(id),
		Namespace: jsii.String(id),
	    },
	    Spec: &k8s.IngressSpec{
	        IngressClassName: jsii.String("nginx"),
	        Rules: &[]*k8s.IngressRule{
	            {
	                Host: jsii.String(id + ".localhost"),
	                Http: &k8s.HttpIngressRuleValue{
	                    Paths: &[]*k8s.HttpIngressPath{
	                        {
	                            Path:     jsii.String("/"),
	                            PathType: jsii.String("Prefix"),
	                            Backend: &k8s.IngressBackend{
	                                Service: &k8s.IngressServiceBackend{
	                                    Name: jsii.String("argo-workflows-c8b79b5a-server"),
	                                    Port: &k8s.ServiceBackendPort{
	                                        Number: jsii.Number(2746),
	                                    },
	                                },
	                            },
	                        },
	                    },
	                },
	            },
	        },
	    },
	})

	return chart
}

func NewNginx(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String("testtesttest"),
		Chart: jsii.String("bitnami/nginx"), //helm repo add bitnami https://charts.bitnami.com/bitnami
		Values: &map[string]interface{}{
			"service": map[string]interface{}{
				"ports": map[string]interface{}{
					"http":  "8080",
				},
			},
		},
	})

	return chart
}

func NewIngressNginx(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String(id),
		Chart: jsii.String("ingress-nginx/ingress-nginx"), //helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	})

	return chart
}

func NewCertManager(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// define resources here
	cdk8s.NewHelm(chart, jsii.String(id), &cdk8s.HelmProps{
		Namespace: jsii.String("cert-manager"),
		Chart: jsii.String("jetstack/cert-manager"), //helm repo add jetstack https://charts.jetstack.io
	})

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
