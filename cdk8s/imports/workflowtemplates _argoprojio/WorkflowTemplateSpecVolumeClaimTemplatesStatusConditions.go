package workflowtemplates _argoprojio

import (
	"time"
)

type WorkflowTemplateSpecVolumeClaimTemplatesStatusConditions struct {
	Status *string `field:"required" json:"status" yaml:"status"`
	Type *string `field:"required" json:"type" yaml:"type"`
	LastProbeTime *time.Time `field:"optional" json:"lastProbeTime" yaml:"lastProbeTime"`
	LastTransitionTime *time.Time `field:"optional" json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Message *string `field:"optional" json:"message" yaml:"message"`
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

