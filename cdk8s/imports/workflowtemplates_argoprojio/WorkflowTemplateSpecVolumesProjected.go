package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesProjected struct {
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	Sources *[]*WorkflowTemplateSpecVolumesProjectedSources `field:"optional" json:"sources" yaml:"sources"`
}

