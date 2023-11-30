package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecVolumesAzureFile struct {
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

