package workflows _argoprojio


type WorkflowSpecVolumesEmptyDir struct {
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	SizeLimit WorkflowSpecVolumesEmptyDirSizeLimit `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

