package workflows_argoprojio


type WorkflowSpecVolumesProjectedSourcesSecret struct {
	Items *[]*WorkflowSpecVolumesProjectedSourcesSecretItems `field:"optional" json:"items" yaml:"items"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

