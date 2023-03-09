module "resource_group" {
  source = "../modules/resource_group"
  common = var.common
}

module "network" {
  source         = "../modules/network"
  common         = var.common
  network        = var.network
  resource_group = module.resource_group.this
}
