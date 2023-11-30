package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItems struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	FieldRef *CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	ResourceFieldRef *CronWorkflowSpecWorkflowSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

