package workflows_argoprojio


type WorkflowSpecTemplatesVolumesEmptyDir struct {
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	SizeLimit WorkflowSpecTemplatesVolumesEmptyDirSizeLimit `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

