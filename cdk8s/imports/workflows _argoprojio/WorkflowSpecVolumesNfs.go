package workflows _argoprojio


type WorkflowSpecVolumesNfs struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	Server *string `field:"required" json:"server" yaml:"server"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

