package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGc struct {
	PodMetadata *CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy CronWorkflowSpecWorkflowSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

