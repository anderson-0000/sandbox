package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesHttpHeaders struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplatesHttpHeadersValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

