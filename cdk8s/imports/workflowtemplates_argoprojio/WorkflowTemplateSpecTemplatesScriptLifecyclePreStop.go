package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesScriptLifecyclePreStop struct {
	Exec *WorkflowTemplateSpecTemplatesScriptLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *WorkflowTemplateSpecTemplatesScriptLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *WorkflowTemplateSpecTemplatesScriptLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

