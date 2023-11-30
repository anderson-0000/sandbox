package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesAzureFile struct {
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

