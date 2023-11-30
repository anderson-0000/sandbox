package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItems struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	FieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecTemplateDefaultsVolumesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

