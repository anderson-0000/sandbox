package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsHttpHeaders struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowTemplateSpecTemplateDefaultsHttpHeadersValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

