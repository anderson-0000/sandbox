package workflowtemplates _argoprojio


type WorkflowTemplateSpecArtifactGc struct {
	ForceFinalizerRemoval *bool `field:"optional" json:"forceFinalizerRemoval" yaml:"forceFinalizerRemoval"`
	PodMetadata *WorkflowTemplateSpecArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowTemplateSpecArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

