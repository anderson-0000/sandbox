package cronworkflow_argoprojio


type CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

