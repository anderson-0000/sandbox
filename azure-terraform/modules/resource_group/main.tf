resource "azurerm_resource_group" "rg" {
  name     = "${var.common["${terraform.workspace}.name"]}-rg"
  location = var.common["${terraform.workspace}.location"]
}
