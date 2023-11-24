package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItems struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	FieldRef *CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

