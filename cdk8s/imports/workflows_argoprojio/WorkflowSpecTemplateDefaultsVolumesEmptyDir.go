package workflows_argoprojio


type WorkflowSpecTemplateDefaultsVolumesEmptyDir struct {
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	SizeLimit WorkflowSpecTemplateDefaultsVolumesEmptyDirSizeLimit `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

