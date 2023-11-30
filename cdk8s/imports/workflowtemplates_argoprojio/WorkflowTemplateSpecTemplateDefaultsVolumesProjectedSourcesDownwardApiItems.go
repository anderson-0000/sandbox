package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItems struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	FieldRef *WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	ResourceFieldRef *WorkflowTemplateSpecTemplateDefaultsVolumesProjectedSourcesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

