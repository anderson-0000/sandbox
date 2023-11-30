package workflows_argoprojio


type WorkflowSpecVolumesDownwardApiItems struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	FieldRef *WorkflowSpecVolumesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	ResourceFieldRef *WorkflowSpecVolumesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

