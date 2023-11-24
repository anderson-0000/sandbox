package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGet struct {
	Port WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*WorkflowTemplateSpecTemplatesContainerSetContainersReadinessProbeHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

