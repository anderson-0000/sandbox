package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGc struct {
	PodMetadata *CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy CronWorkflowSpecWorkflowSpecHooksArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

