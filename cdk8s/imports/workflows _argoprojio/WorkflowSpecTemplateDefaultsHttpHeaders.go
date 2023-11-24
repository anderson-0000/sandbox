package workflows _argoprojio


type WorkflowSpecTemplateDefaultsHttpHeaders struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplateDefaultsHttpHeadersValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

