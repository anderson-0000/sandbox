resource "azurerm_network_security_group" "from_home_allow_sg" {
  name                = "${var.common["${terraform.workspace}.name"]}-sg"
  location            = var.common["${terraform.workspace}.location"]
  resource_group_name = var.resource_group.name

  security_rule {
    name                       = "allow_from_home_sr"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_ranges    = ["22", "80", "443"]
    source_address_prefix      = "217.178.201.217"
    destination_address_prefix = "*"
  }

  tags = {
    environment = var.common["${terraform.workspace}.name"]
  }
}

resource "azurerm_virtual_network" "vnet" {
  name                = "${var.common["${terraform.workspace}.name"]}_vnet"
  location            = var.common["${terraform.workspace}.location"]
  resource_group_name = var.resource_group.name
  address_space       = [var.network["${terraform.workspace}.vnet_add"]]

  subnet {
    name           = "dmz"
    address_prefix = var.network["${terraform.workspace}.sub_add"]
    security_group = azurerm_network_security_group.from_home_allow_sg.id
  }

  tags = {
    environment = var.common["${terraform.workspace}.name"]
  }
}
