package workflows_argoprojio


type WorkflowSpecTemplateDefaultsContainerStartupProbeGrpc struct {
	Port *float64 `field:"required" json:"port" yaml:"port"`
	Service *string `field:"optional" json:"service" yaml:"service"`
}

