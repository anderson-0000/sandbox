package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStart struct {
	Exec *WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

