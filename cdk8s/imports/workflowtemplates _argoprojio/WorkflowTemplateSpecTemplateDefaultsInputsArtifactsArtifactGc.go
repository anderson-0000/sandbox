package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGc struct {
	PodMetadata *WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcPodMetadata `field:"optional" json:"podMetadata" yaml:"podMetadata"`
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	Strategy WorkflowTemplateSpecTemplateDefaultsInputsArtifactsArtifactGcStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

