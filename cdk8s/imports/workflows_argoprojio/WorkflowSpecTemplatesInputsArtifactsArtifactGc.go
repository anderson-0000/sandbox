package workflows_argoprojio


type WorkflowSpecTemplatesInputsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplatesInputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

