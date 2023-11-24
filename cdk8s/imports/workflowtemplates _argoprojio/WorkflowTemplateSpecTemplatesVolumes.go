package workflowtemplates _argoprojio


type WorkflowTemplateSpecTemplatesVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *WorkflowTemplateSpecTemplatesVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *WorkflowTemplateSpecTemplatesVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *WorkflowTemplateSpecTemplatesVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *WorkflowTemplateSpecTemplatesVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *WorkflowTemplateSpecTemplatesVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *WorkflowTemplateSpecTemplatesVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *WorkflowTemplateSpecTemplatesVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *WorkflowTemplateSpecTemplatesVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *WorkflowTemplateSpecTemplatesVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *WorkflowTemplateSpecTemplatesVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *WorkflowTemplateSpecTemplatesVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *WorkflowTemplateSpecTemplatesVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *WorkflowTemplateSpecTemplatesVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *WorkflowTemplateSpecTemplatesVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *WorkflowTemplateSpecTemplatesVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *WorkflowTemplateSpecTemplatesVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *WorkflowTemplateSpecTemplatesVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *WorkflowTemplateSpecTemplatesVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *WorkflowTemplateSpecTemplatesVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *WorkflowTemplateSpecTemplatesVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *WorkflowTemplateSpecTemplatesVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *WorkflowTemplateSpecTemplatesVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *WorkflowTemplateSpecTemplatesVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *WorkflowTemplateSpecTemplatesVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *WorkflowTemplateSpecTemplatesVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *WorkflowTemplateSpecTemplatesVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *WorkflowTemplateSpecTemplatesVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *WorkflowTemplateSpecTemplatesVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *WorkflowTemplateSpecTemplatesVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

