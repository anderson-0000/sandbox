package workflowtemplates_argoprojio


type WorkflowTemplateSpecVolumes struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	AwsElasticBlockStore *WorkflowTemplateSpecVolumesAwsElasticBlockStore `field:"optional" json:"awsElasticBlockStore" yaml:"awsElasticBlockStore"`
	AzureDisk *WorkflowTemplateSpecVolumesAzureDisk `field:"optional" json:"azureDisk" yaml:"azureDisk"`
	AzureFile *WorkflowTemplateSpecVolumesAzureFile `field:"optional" json:"azureFile" yaml:"azureFile"`
	Cephfs *WorkflowTemplateSpecVolumesCephfs `field:"optional" json:"cephfs" yaml:"cephfs"`
	Cinder *WorkflowTemplateSpecVolumesCinder `field:"optional" json:"cinder" yaml:"cinder"`
	ConfigMap *WorkflowTemplateSpecVolumesConfigMap `field:"optional" json:"configMap" yaml:"configMap"`
	Csi *WorkflowTemplateSpecVolumesCsi `field:"optional" json:"csi" yaml:"csi"`
	DownwardApi *WorkflowTemplateSpecVolumesDownwardApi `field:"optional" json:"downwardApi" yaml:"downwardApi"`
	EmptyDir *WorkflowTemplateSpecVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	Ephemeral *WorkflowTemplateSpecVolumesEphemeral `field:"optional" json:"ephemeral" yaml:"ephemeral"`
	Fc *WorkflowTemplateSpecVolumesFc `field:"optional" json:"fc" yaml:"fc"`
	FlexVolume *WorkflowTemplateSpecVolumesFlexVolume `field:"optional" json:"flexVolume" yaml:"flexVolume"`
	Flocker *WorkflowTemplateSpecVolumesFlocker `field:"optional" json:"flocker" yaml:"flocker"`
	GcePersistentDisk *WorkflowTemplateSpecVolumesGcePersistentDisk `field:"optional" json:"gcePersistentDisk" yaml:"gcePersistentDisk"`
	GitRepo *WorkflowTemplateSpecVolumesGitRepo `field:"optional" json:"gitRepo" yaml:"gitRepo"`
	Glusterfs *WorkflowTemplateSpecVolumesGlusterfs `field:"optional" json:"glusterfs" yaml:"glusterfs"`
	HostPath *WorkflowTemplateSpecVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	Iscsi *WorkflowTemplateSpecVolumesIscsi `field:"optional" json:"iscsi" yaml:"iscsi"`
	Nfs *WorkflowTemplateSpecVolumesNfs `field:"optional" json:"nfs" yaml:"nfs"`
	PersistentVolumeClaim *WorkflowTemplateSpecVolumesPersistentVolumeClaim `field:"optional" json:"persistentVolumeClaim" yaml:"persistentVolumeClaim"`
	PhotonPersistentDisk *WorkflowTemplateSpecVolumesPhotonPersistentDisk `field:"optional" json:"photonPersistentDisk" yaml:"photonPersistentDisk"`
	PortworxVolume *WorkflowTemplateSpecVolumesPortworxVolume `field:"optional" json:"portworxVolume" yaml:"portworxVolume"`
	Projected *WorkflowTemplateSpecVolumesProjected `field:"optional" json:"projected" yaml:"projected"`
	Quobyte *WorkflowTemplateSpecVolumesQuobyte `field:"optional" json:"quobyte" yaml:"quobyte"`
	Rbd *WorkflowTemplateSpecVolumesRbd `field:"optional" json:"rbd" yaml:"rbd"`
	ScaleIo *WorkflowTemplateSpecVolumesScaleIo `field:"optional" json:"scaleIo" yaml:"scaleIo"`
	Secret *WorkflowTemplateSpecVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
	Storageos *WorkflowTemplateSpecVolumesStorageos `field:"optional" json:"storageos" yaml:"storageos"`
	VsphereVolume *WorkflowTemplateSpecVolumesVsphereVolume `field:"optional" json:"vsphereVolume" yaml:"vsphereVolume"`
}

