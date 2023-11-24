package workflows _argoprojio


type WorkflowSpecTemplateDefaultsData struct {
	Source *WorkflowSpecTemplateDefaultsDataSource `field:"required" json:"source" yaml:"source"`
	Transformation *[]*WorkflowSpecTemplateDefaultsDataTransformation `field:"required" json:"transformation" yaml:"transformation"`
}

