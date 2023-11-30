package workflows_argoprojio


type WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

