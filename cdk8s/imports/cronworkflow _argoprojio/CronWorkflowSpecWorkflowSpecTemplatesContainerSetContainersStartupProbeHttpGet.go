package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGet struct {
	Port CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*CronWorkflowSpecWorkflowSpecTemplatesContainerSetContainersStartupProbeHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

