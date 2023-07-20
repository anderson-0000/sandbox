package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/k8s"
)

func NewPvAndPvc(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	c := cdk8s.NewChart(scope, jsii.String(id), &cprops) // idがファイル名になる

	storage := 1.0
	storageByte := 1073741824.0 //1Giのこと

	k8s.NewKubePersistentVolume(c, jsii.String("pv1"), &k8s.KubePersistentVolumeProps{
                Metadata: &k8s.ObjectMeta{
                        Name: jsii.String("pv1"),
                },
                Spec: &k8s.PersistentVolumeSpec{
                        AccessModes: &[]*string{
				jsii.String("ReadWriteOnce"),
			},
                        Capacity: &map[string]k8s.Quantity{
                                "storage": k8s.Quantity_FromNumber(&storageByte),
                        },
                        PersistentVolumeReclaimPolicy: jsii.String("Delete"),
                        HostPath: &k8s.HostPathVolumeSource{
                                Path: jsii.String("/var/lib/k8s-pvs/pv1/pv1"),
                        },
			StorageClassName: jsii.String(""),
                },
        })

	cdk8splus26.NewPersistentVolumeClaim(c, jsii.String("pvc1"), &cdk8splus26.PersistentVolumeClaimProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("pvc1"),
		},
		AccessModes: &[]cdk8splus26.PersistentVolumeAccessMode{
			cdk8splus26.PersistentVolumeAccessMode_READ_WRITE_ONCE,
		},
		Storage: cdk8s.Size_Gibibytes(&storage),
		Volume: cdk8splus26.PersistentVolume_FromPersistentVolumeName(c, jsii.String("pv1-volume-name"), jsii.String("pv1")),
		StorageClassName: jsii.String(""),
	})

	return c
}
