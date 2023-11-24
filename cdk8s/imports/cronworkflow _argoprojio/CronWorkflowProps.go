package cronworkflow _argoprojio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CronWorkflowProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	Spec *CronWorkflowSpec `field:"required" json:"spec" yaml:"spec"`
}

