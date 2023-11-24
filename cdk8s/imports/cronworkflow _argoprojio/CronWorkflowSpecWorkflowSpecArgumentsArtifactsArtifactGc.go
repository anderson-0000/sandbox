package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGc struct {
	PodMetadata *CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy CronWorkflowSpecWorkflowSpecArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

