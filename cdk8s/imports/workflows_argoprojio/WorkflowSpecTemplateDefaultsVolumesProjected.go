package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesProjected struct {
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	Sources *[]*WorkflowSpecTemplateDefaultsVolumesProjectedSources `field:"optional" json:"sources" yaml:"sources"`
}

