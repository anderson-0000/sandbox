package workflows _argoprojio


type WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGc struct {
	PodMetadata *WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplatesResourceManifestFromArtifactArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

