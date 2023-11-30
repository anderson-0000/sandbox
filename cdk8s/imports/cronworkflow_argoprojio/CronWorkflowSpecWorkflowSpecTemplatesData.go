package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesData struct {
	Source *CronWorkflowSpecWorkflowSpecTemplatesDataSource `field:"required" json:"source" yaml:"source"`
	Transformation *[]*CronWorkflowSpecWorkflowSpecTemplatesDataTransformation `field:"required" json:"transformation" yaml:"transformation"`
}

