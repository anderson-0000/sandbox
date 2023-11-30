package workflows_argoprojio


type WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItems struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	FieldRef *WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	ResourceFieldRef *WorkflowSpecTemplatesVolumesProjectedSourcesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

