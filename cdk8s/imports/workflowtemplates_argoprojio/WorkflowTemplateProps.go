package workflowtemplates_argoprojio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type WorkflowTemplateProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	Spec *WorkflowTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

