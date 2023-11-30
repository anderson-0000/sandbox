package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsData struct {
	Source *CronWorkflowSpecWorkflowSpecTemplateDefaultsDataSource `field:"required" json:"source" yaml:"source"`
	Transformation *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsDataTransformation `field:"required" json:"transformation" yaml:"transformation"`
}

