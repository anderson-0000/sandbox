package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesContainerSet struct {
	Containers *[]*WorkflowTemplateSpecTemplatesContainerSetContainers `field:"required" json:"containers" yaml:"containers"`
	RetryStrategy *WorkflowTemplateSpecTemplatesContainerSetRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	VolumeMounts *[]*WorkflowTemplateSpecTemplatesContainerSetVolumeMounts `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

