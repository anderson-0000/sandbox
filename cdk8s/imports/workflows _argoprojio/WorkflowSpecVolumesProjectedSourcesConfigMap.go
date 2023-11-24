package workflows _argoprojio


type WorkflowSpecVolumesProjectedSourcesConfigMap struct {
	Items *[]*WorkflowSpecVolumesProjectedSourcesConfigMapItems `field:"optional" json:"items" yaml:"items"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

