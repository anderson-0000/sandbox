package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

