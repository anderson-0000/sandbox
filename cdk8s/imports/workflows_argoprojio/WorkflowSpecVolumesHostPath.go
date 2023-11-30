package workflows_argoprojio


type WorkflowSpecVolumesHostPath struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	Type *string `field:"optional" json:"type" yaml:"type"`
}

