package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGet struct {
	Port WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*WorkflowTemplateSpecTemplateDefaultsContainerLifecyclePostStartHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

