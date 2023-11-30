package workflows_argoprojio


type WorkflowSpecVolumesProjectedSourcesDownwardApiItems struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	FieldRef *WorkflowSpecVolumesProjectedSourcesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	ResourceFieldRef *WorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

