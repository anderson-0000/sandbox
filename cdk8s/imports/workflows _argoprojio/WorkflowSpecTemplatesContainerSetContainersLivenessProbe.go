package workflows _argoprojio


type WorkflowSpecTemplatesContainerSetContainersLivenessProbe struct {
	Exec *WorkflowSpecTemplatesContainerSetContainersLivenessProbeExec `field:"optional" json:"exec" yaml:"exec"`
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	Grpc *WorkflowSpecTemplatesContainerSetContainersLivenessProbeGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	HttpGet *WorkflowSpecTemplatesContainerSetContainersLivenessProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	TcpSocket *WorkflowSpecTemplatesContainerSetContainersLivenessProbeTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	TerminationGracePeriodSeconds *float64 `field:"optional" json:"terminationGracePeriodSeconds" yaml:"terminationGracePeriodSeconds"`
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

