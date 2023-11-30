package workflows_argoprojio


type WorkflowSpecTemplatesMemoize struct {
	Cache *WorkflowSpecTemplatesMemoizeCache `field:"required" json:"cache" yaml:"cache"`
	Key *string `field:"required" json:"key" yaml:"key"`
	MaxAge *string `field:"required" json:"maxAge" yaml:"maxAge"`
}

