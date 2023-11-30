package workflows_argoprojio


type WorkflowSpecTemplatesOutputsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplatesOutputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

