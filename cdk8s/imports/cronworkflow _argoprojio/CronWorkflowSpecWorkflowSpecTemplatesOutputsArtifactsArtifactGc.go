package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGc struct {
	PodMetadata *CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy CronWorkflowSpecWorkflowSpecTemplatesOutputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

