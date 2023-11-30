package workflows_argoprojio


type WorkflowSpecArgumentsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

