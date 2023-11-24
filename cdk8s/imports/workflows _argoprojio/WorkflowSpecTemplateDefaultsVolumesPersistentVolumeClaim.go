package workflows _argoprojio


type WorkflowSpecTemplateDefaultsVolumesPersistentVolumeClaim struct {
	ClaimName *string `field:"required" json:"claimName" yaml:"claimName"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

