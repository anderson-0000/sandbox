package workflows _argoprojio


type WorkflowSpecWorkflowMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	LabelsFrom *map[string]*WorkflowSpecWorkflowMetadataLabelsFrom `field:"optional" json:"labelsFrom" yaml:"labelsFrom"`
}

