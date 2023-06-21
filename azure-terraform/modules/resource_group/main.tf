resource "azurerm_resource_group" "this" {
  name     = "${var.common["${terraform.workspace}.name"]}-rg"
  location = var.common["${terraform.workspace}.location"]
}
