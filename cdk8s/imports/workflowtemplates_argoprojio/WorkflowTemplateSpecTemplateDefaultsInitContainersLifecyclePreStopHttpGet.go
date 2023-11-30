package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGet struct {
	Port WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*WorkflowTemplateSpecTemplateDefaultsInitContainersLifecyclePreStopHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

