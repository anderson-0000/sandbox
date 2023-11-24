package workflows _argoprojio


type WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplateDefaultsDagTasksArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

