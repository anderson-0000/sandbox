package workflows_argoprojio


type WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGc struct {
	PodMetadata *WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecTemplateDefaultsResourceManifestFromArtifactArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

