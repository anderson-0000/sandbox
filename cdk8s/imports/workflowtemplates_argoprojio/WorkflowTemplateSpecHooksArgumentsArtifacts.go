package workflowtemplates_argoprojio


type WorkflowTemplateSpecHooksArgumentsArtifacts struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Archive *WorkflowTemplateSpecHooksArgumentsArtifactsArchive `field:"optional" json:"archive" yaml:"archive"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	ArtifactGc *WorkflowTemplateSpecHooksArgumentsArtifactsArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	Artifactory *WorkflowTemplateSpecHooksArgumentsArtifactsArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowTemplateSpecHooksArgumentsArtifactsAzure `field:"optional" json:"azure" yaml:"azure"`
	Deleted *bool `field:"optional" json:"deleted" yaml:"deleted"`
	From *string `field:"optional" json:"from" yaml:"from"`
	FromExpression *string `field:"optional" json:"fromExpression" yaml:"fromExpression"`
	Gcs *WorkflowTemplateSpecHooksArgumentsArtifactsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowTemplateSpecHooksArgumentsArtifactsGit `field:"optional" json:"git" yaml:"git"`
	GlobalName *string `field:"optional" json:"globalName" yaml:"globalName"`
	Hdfs *WorkflowTemplateSpecHooksArgumentsArtifactsHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowTemplateSpecHooksArgumentsArtifactsHttp `field:"optional" json:"http" yaml:"http"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	Oss *WorkflowTemplateSpecHooksArgumentsArtifactsOss `field:"optional" json:"oss" yaml:"oss"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Raw *WorkflowTemplateSpecHooksArgumentsArtifactsRaw `field:"optional" json:"raw" yaml:"raw"`
	RecurseMode *bool `field:"optional" json:"recurseMode" yaml:"recurseMode"`
	S3 *WorkflowTemplateSpecHooksArgumentsArtifactsS3 `field:"optional" json:"s3" yaml:"s3"`
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

