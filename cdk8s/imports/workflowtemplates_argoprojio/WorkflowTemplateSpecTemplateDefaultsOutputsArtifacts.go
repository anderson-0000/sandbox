package workflowtemplates_argoprojio


type WorkflowTemplateSpecTemplateDefaultsOutputsArtifacts struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Archive *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArchive `field:"optional" json:"archive" yaml:"archive"`
	ArchiveLogs *bool `field:"optional" json:"archiveLogs" yaml:"archiveLogs"`
	ArtifactGc *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactGc `field:"optional" json:"artifactGc" yaml:"artifactGc"`
	Artifactory *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsArtifactory `field:"optional" json:"artifactory" yaml:"artifactory"`
	Azure *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsAzure `field:"optional" json:"azure" yaml:"azure"`
	Deleted *bool `field:"optional" json:"deleted" yaml:"deleted"`
	From *string `field:"optional" json:"from" yaml:"from"`
	FromExpression *string `field:"optional" json:"fromExpression" yaml:"fromExpression"`
	Gcs *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGcs `field:"optional" json:"gcs" yaml:"gcs"`
	Git *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsGit `field:"optional" json:"git" yaml:"git"`
	GlobalName *string `field:"optional" json:"globalName" yaml:"globalName"`
	Hdfs *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHdfs `field:"optional" json:"hdfs" yaml:"hdfs"`
	Http *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsHttp `field:"optional" json:"http" yaml:"http"`
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
	Oss *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsOss `field:"optional" json:"oss" yaml:"oss"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Raw *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsRaw `field:"optional" json:"raw" yaml:"raw"`
	RecurseMode *bool `field:"optional" json:"recurseMode" yaml:"recurseMode"`
	S3 *WorkflowTemplateSpecTemplateDefaultsOutputsArtifactsS3 `field:"optional" json:"s3" yaml:"s3"`
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
}

