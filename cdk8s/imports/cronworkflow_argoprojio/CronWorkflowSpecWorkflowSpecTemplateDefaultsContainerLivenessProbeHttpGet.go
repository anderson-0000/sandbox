package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGet struct {
	Port CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*CronWorkflowSpecWorkflowSpecTemplateDefaultsContainerLivenessProbeHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

