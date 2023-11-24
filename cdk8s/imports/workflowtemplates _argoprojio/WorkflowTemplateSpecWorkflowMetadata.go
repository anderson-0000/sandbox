package workflowtemplates _argoprojio


type WorkflowTemplateSpecWorkflowMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	LabelsFrom *map[string]*WorkflowTemplateSpecWorkflowMetadataLabelsFrom `field:"optional" json:"labelsFrom" yaml:"labelsFrom"`
}

