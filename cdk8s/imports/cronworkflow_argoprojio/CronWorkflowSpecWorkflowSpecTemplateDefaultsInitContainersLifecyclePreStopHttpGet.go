package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet struct {
	Port CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

