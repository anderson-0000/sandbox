package workflows _argoprojio


type WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

