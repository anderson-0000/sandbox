package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGc struct {
	PodMetadata *CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy CronWorkflowSpecWorkflowSpecTemplatesInputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

