package workflows _argoprojio


type WorkflowSpecTemplatesVolumesQuobyte struct {
	Registry *string `field:"required" json:"registry" yaml:"registry"`
	Volume *string `field:"required" json:"volume" yaml:"volume"`
	Group *string `field:"optional" json:"group" yaml:"group"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	Tenant *string `field:"optional" json:"tenant" yaml:"tenant"`
	User *string `field:"optional" json:"user" yaml:"user"`
}

