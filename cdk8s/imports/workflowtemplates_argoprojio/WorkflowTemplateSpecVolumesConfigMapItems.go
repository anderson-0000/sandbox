package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumesConfigMapItems struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Path *string `field:"required" json:"path" yaml:"path"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
}

