package workflows _argoprojio


type WorkflowSpecTemplateDefaultsMemoize struct {
	Cache *WorkflowSpecTemplateDefaultsMemoizeCache `field:"required" json:"cache" yaml:"cache"`
	Key *string `field:"required" json:"key" yaml:"key"`
	MaxAge *string `field:"required" json:"maxAge" yaml:"maxAge"`
}

