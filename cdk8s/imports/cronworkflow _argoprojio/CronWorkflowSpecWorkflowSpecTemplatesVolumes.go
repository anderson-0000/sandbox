package cronworkflow _argoprojio


type CronWorkflowSpecWorkflowSpecTemplatesVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *CronWorkflowSpecWorkflowSpecTemplatesVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *CronWorkflowSpecWorkflowSpecTemplatesVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *CronWorkflowSpecWorkflowSpecTemplatesVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *CronWorkflowSpecWorkflowSpecTemplatesVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *CronWorkflowSpecWorkflowSpecTemplatesVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *CronWorkflowSpecWorkflowSpecTemplatesVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *CronWorkflowSpecWorkflowSpecTemplatesVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *CronWorkflowSpecWorkflowSpecTemplatesVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *CronWorkflowSpecWorkflowSpecTemplatesVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *CronWorkflowSpecWorkflowSpecTemplatesVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *CronWorkflowSpecWorkflowSpecTemplatesVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *CronWorkflowSpecWorkflowSpecTemplatesVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *CronWorkflowSpecWorkflowSpecTemplatesVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *CronWorkflowSpecWorkflowSpecTemplatesVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *CronWorkflowSpecWorkflowSpecTemplatesVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *CronWorkflowSpecWorkflowSpecTemplatesVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *CronWorkflowSpecWorkflowSpecTemplatesVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *CronWorkflowSpecWorkflowSpecTemplatesVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *CronWorkflowSpecWorkflowSpecTemplatesVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *CronWorkflowSpecWorkflowSpecTemplatesVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *CronWorkflowSpecWorkflowSpecTemplatesVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *CronWorkflowSpecWorkflowSpecTemplatesVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *CronWorkflowSpecWorkflowSpecTemplatesVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *CronWorkflowSpecWorkflowSpecTemplatesVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *CronWorkflowSpecWorkflowSpecTemplatesVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *CronWorkflowSpecWorkflowSpecTemplatesVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *CronWorkflowSpecWorkflowSpecTemplatesVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *CronWorkflowSpecWorkflowSpecTemplatesVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *CronWorkflowSpecWorkflowSpecTemplatesVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

