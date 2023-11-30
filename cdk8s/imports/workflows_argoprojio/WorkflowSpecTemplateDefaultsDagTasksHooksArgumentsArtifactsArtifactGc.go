package workflows_argoprojio


type WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplateDefaultsDagTasksHooksArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

