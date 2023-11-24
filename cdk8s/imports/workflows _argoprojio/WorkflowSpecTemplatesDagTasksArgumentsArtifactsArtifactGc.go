package workflows _argoprojio


type WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplatesDagTasksArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

