package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptLifecyclePreStop struct {
	Exec *WorkflowSpecTemplateDefaultsScriptLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsScriptLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

