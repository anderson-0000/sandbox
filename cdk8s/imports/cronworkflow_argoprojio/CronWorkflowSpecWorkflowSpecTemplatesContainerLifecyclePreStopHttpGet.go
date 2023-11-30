package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGet struct {
	Port CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*CronWorkflowSpecWorkflowSpecTemplatesContainerLifecyclePreStopHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

