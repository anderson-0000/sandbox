package workflows _argoprojio


type WorkflowSpecVolumesEphemeralVolumeClaimTemplate struct {
	Spec *WorkflowSpecVolumesEphemeralVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

