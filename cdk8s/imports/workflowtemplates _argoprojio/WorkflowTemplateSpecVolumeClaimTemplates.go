package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumeClaimTemplates struct {
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	Spec *WorkflowTemplateSpecVolumeClaimTemplatesSpec `field:"optional" json:"spec" yaml:"spec"`
	Status *WorkflowTemplateSpecVolumeClaimTemplatesStatus `field:"optional" json:"status" yaml:"status"`
}

