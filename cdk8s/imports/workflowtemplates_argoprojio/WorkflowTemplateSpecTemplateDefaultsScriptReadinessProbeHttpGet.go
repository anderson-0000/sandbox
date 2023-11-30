package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGet struct {
	Port WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGetPort `field:"required" json:"port" yaml:"port"`
	Host *string `field:"optional" json:"host" yaml:"host"`
	HttpHeaders *[]*WorkflowTemplateSpecTemplateDefaultsScriptReadinessProbeHttpGetHttpHeaders `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

