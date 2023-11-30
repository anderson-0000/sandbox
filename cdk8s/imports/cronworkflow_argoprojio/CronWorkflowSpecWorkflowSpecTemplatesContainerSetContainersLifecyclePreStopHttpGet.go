package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGet struct {
	Port CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersLifecyclePreStopHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

