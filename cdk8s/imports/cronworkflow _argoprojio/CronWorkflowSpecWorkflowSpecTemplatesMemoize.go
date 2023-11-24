package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesMemoize struct {
	Cache *CronWorkflowSpecWorkflowSpecTemplatesMemoizeCache `field:"required" json:"cache" yaml:"cache"`
	Key *string `field:"required" json:"key" yaml:"key"`
	MaxAge *string `field:"required" json:"maxAge" yaml:"maxAge"`
}

