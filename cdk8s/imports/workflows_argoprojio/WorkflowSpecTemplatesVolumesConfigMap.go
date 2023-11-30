package workflows_argoprojio


type WorkflowSpecTemplatesVolumesConfigMap struct {
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	Items *[]*WorkflowSpecTemplatesVolumesConfigMapItems `field:"optional" json:"items" yaml:"items"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

