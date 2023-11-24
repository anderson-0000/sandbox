package workflows _argoprojio


type WorkflowSpecVolumeClaimTemplates struct {
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	Spec *WorkflowSpecVolumeClaimTemplatesSpec `field:"optional" json:"spec" yaml:"spec"`
	Status *WorkflowSpecVolumeClaimTemplatesStatus `field:"optional" json:"status" yaml:"status"`
}

