package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumeClaimTemplates struct {
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	Spec *CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesSpec `field:"optional" json:"spec" yaml:"spec"`
	Status *CronWorkflowSpecWorkflowSpecVolumeClaimTemplatesStatus `field:"optional" json:"status" yaml:"status"`
}

