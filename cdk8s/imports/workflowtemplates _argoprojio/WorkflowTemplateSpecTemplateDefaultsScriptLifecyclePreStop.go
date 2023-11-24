package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplateDefaultsScriptLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

