package workflows _argoprojio


type WorkflowSpecArtifactGc struct {
	ForceFinalizerRemoval *bool `field:"optional" json:"forceFinalizerRemoval" yaml:"forceFinalizerRemoval"`
	PodMetadata *WorkflowSpecArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowSpecArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

