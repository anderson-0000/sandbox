package workflowtemplates _argoprojio


type WorkflowTemplateSpecArgumentsArtifactsArtifactGc struct {
	PodMetadata *WorkflowTemplateSpecArgumentsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowTemplateSpecArgumentsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

