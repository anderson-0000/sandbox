package workflows_argoprojio


type WorkflowSpecTemplatesSidecarsStartupProbeGrpc struct {
	Port *float64 `field:"required" json:"port" yaml:"port"`
	Service *string `field:"optional" json:"service" yaml:"service"`
}

