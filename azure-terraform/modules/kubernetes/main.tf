resource "azurerm_kubernetes_cluster" "this" {
  location            = var.common["${terraform.workspace}.location"]
  name                = "${var.common["${terraform.workspace}.name"]}-cluster"
  resource_group_name = azurerm_resource_group.name
  kubernetes_version  = var.kubernetes.version
  sku_tier            = "Free"
  enable_auto_scaling = True
  tags = { 
    environment = var.common["${terraform.workspace}.name"]
  }

  default_node_pool {
    name       = "${var.common["${terraform.workspace}.name"]}-default"
    vm_size    = "Standard_D2_v2"
    node_count = var.agent_count
    vnet_subnet_id  = data.azurerm_subnet.kubesubnet.id
  }
  linux_profile {
    admin_username = "admin"

    ssh_key {
      key_data = file(var.kubernetes.ssh_public_key)
    }
  }
  network_profile {
    network_plugin    = "azure"
    dns_service_ip     = var.aks_dns_service_ip
    docker_bridge_cidr = var.aks_docker_bridge_cidr
    service_cidr       = var.aks_service_cidr     /16
  }
}
