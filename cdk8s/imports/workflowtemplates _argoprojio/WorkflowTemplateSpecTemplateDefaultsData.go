package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsData struct {
	Source *WorkflowTemplateSpecTemplateDefaultsDataSource `field:"required" json:"source" yaml:"source"`
	Transformation *[]*WorkflowTemplateSpecTemplateDefaultsDataTransformation `field:"required" json:"transformation" yaml:"transformation"`
}

