package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumesEmptyDir struct {
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	SizeLimit WorkflowTemplateSpecVolumesEmptyDirSizeLimit `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

