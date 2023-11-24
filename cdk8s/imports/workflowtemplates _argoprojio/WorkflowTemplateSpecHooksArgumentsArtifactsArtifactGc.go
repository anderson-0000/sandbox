package workflowtemplates _argoprojio


type WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGc struct {
	PodMetadata *WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

