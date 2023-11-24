package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecArtifactGc struct {
	ForceFinalizerRemoval *bool `field:"optional" json:"forceFinalizerRemoval" yaml:"forceFinalizerRemoval"`
	PodMetadata *CronWorkflowSpecWorkflowSpecArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	PodSpecPatch *string `field:"optional" json:"podSpecPatch" yaml:"podSpecPatch"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy CronWorkflowSpecWorkflowSpecArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

