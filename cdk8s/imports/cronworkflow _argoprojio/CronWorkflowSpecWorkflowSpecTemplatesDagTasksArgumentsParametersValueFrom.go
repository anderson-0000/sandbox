package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParametersValueFrom struct {
	ConfigMapKeyRef *CronWorkflowSpecWorkflowSpecTemplatesDagTasksArgumentsParametersValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	Default *string `field:"optional" json:"default" yaml:"default"`
	Event *string `field:"optional" json:"event" yaml:"event"`
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	JqFilter *string `field:"optional" json:"jqFilter" yaml:"jqFilter"`
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Supplied interface{} `field:"optional" json:"supplied" yaml:"supplied"`
}

