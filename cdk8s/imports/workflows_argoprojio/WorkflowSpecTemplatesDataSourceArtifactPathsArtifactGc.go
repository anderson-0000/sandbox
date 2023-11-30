package workflows_argoprojio


type WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGc struct {
	PodMetadata *WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplatesDataSourceArtifactPathsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

