package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifacts struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Archive *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArchive `field:"optional" json:"archive" yaml:"archive"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	ArtifactGc *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	Artifactory *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsAzure `field:"optional" json:"azure" yaml:"azure"`
	Deleted *bool `field:"optional" json:"deleted" yaml:"deleted"`
	From *string `field:"optional" json:"from" yaml:"from"`
	FromExpression *string `field:"optional" json:"fromExpression" yaml:"fromExpression"`
	Gcs *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsGit `field:"optional" json:"git" yaml:"git"`
	GlobalName *string `field:"optional" json:"globalName" yaml:"globalName"`
	Hdfs *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsHttp `field:"optional" json:"http" yaml:"http"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	Oss *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsOss `field:"optional" json:"oss" yaml:"oss"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Raw *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsRaw `field:"optional" json:"raw" yaml:"raw"`
	RecurseMode *bool `field:"optional" json:"recurseMode" yaml:"recurseMode"`
	S3 *WorkflowTemplateSpecTemplatesDagTasksHooksArgumentsArtifactsS3 `field:"optional" json:"s3" yaml:"s3"`
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

