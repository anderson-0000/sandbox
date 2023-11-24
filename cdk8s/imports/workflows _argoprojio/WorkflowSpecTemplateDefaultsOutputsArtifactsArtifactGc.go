package workflows _argoprojio


type WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplateDefaultsOutputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

