package workflows_argoprojio


type WorkflowSpecTemplatesVolumesProjected struct {
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	Sources *[]*WorkflowSpecTemplatesVolumesProjectedSources `field:"optional" json:"sources" yaml:"sources"`
}

