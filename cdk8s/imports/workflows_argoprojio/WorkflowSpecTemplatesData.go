package workflows_argoprojio


type WorkflowSpecTemplatesData struct {
	Source *WorkflowSpecTemplatesDataSource `field:"required" json:"source" yaml:"source"`
	Transformation *[]*WorkflowSpecTemplatesDataTransformation `field:"required" json:"transformation" yaml:"transformation"`
}

