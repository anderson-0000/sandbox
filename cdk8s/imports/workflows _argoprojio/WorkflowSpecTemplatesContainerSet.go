package workflows _argoprojio


type WorkflowSpecTemplatesContainerSet struct {
	Containers *[]*WorkflowSpecTemplatesContainerSetContainers `field:"required" json:"containers" yaml:"containers"`
	RetryStrategy *WorkflowSpecTemplatesContainerSetRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	VolumeMounts *[]*WorkflowSpecTemplatesContainerSetVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

