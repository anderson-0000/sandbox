package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesScriptEnv struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueFrom *CronWorkflowSpecWorkflowSpecTemplatesScriptEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

