package workflows_argoprojio


type WorkflowSpecVolumesProjected struct {
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	Sources *[]*WorkflowSpecVolumesProjectedSources `field:"optional" json:"sources" yaml:"sources"`
}

