package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerSet struct {
	Containers *[]*WorkflowTemplateSpecTemplateDefaultsContainerSetContainers `field:"required" json:"containers" yaml:"containers"`
	RetryStrategy *WorkflowTemplateSpecTemplateDefaultsContainerSetRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	VolumeMounts *[]*WorkflowTemplateSpecTemplateDefaultsContainerSetVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

