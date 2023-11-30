package workflows_argoprojio


type WorkflowSpecTemplateDefaultsScriptLifecyclePostStart struct {
	Exec *WorkflowSpecTemplateDefaultsScriptLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowSpecTemplateDefaultsScriptLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

