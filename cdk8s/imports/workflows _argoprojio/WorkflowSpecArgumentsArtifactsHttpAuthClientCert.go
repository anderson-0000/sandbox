package workflows _argoprojio


type WorkflowSpecArgumentsArtifactsHttpAuthClientCert struct {
	ClientCertSecret *WorkflowSpecArgumentsArtifactsHttpAuthClientCertClientCertSecret `field:"optional" json:"clientCertSecret" yaml:"clientCertSecret"`
	ClientKeySecret *WorkflowSpecArgumentsArtifactsHttpAuthClientCertClientKeySecret `field:"optional" json:"clientKeySecret" yaml:"clientKeySecret"`
}

