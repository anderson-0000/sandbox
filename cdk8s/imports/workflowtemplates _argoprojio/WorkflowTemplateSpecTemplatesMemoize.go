package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesMemoize struct {
	Cache *WorkflowTemplateSpecTemplatesMemoizeCache `field:"required" json:"cache" yaml:"cache"`
	Key *string `field:"required" json:"key" yaml:"key"`
	MaxAge *string `field:"required" json:"maxAge" yaml:"maxAge"`
}

