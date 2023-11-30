package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesVolumesGlusterfs struct {
	Endpoints *string `field:"required" json:"endpoints" yaml:"endpoints"`
	Path *string `field:"required" json:"path" yaml:"path"`
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

