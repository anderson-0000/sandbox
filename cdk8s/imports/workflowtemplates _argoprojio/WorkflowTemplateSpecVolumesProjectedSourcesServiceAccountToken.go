package workflowtemplates _argoprojio


type WorkflowTemplateSpecVolumesProjectedSourcesServiceAccountToken struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	ExpirationSeconds *float64 `field:"optional" json:"expirationSeconds" yaml:"expirationSeconds"`
}

