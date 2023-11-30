package workflows_argoprojio


type WorkflowSpecTemplatesHttpHeaders struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *WorkflowSpecTemplatesHttpHeadersValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

