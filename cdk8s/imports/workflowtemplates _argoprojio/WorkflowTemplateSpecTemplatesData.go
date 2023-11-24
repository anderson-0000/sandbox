package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesData struct {
	Source *WorkflowTemplateSpecTemplatesDataSource `field:"required" json:"source" yaml:"source"`
	Transformation *[]*WorkflowTemplateSpecTemplatesDataTransformation `field:"required" json:"transformation" yaml:"transformation"`
}

