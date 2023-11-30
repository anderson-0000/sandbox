package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerSet struct {
	Containers *[]*WorkflowSpecTemplateDefaultsContainerSetContainers `field:"required" json:"containers" yaml:"containers"`
	RetryStrategy *WorkflowSpecTemplateDefaultsContainerSetRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	VolumeMounts *[]*WorkflowSpecTemplateDefaultsContainerSetVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

