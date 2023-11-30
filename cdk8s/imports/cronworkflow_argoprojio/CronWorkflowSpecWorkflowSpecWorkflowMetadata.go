package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecWorkflowMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	LabelsFrom *map[string]*CronWorkflowSpecWorkflowSpecWorkflowMetadataLabelsFrom `field:"optional" json:"labelsFrom" yaml:"labelsFrom"`
}

