package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStop struct {
	Exec *CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	HttpGet *CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	TcpSocket *CronWorkflowSpecWorkflowSpecTemplatesScriptLifecyclePreStopTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

