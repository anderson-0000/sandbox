variable "common" {
  default = {
    "dev.name"     = "dev"
    "dev.location" = "Japan East"
    "prd.name"     = "prod"
    "prd.location" = "Japan East"
  }
}

variable "network" {
  default = {
    "dev.vnet_add" = "10.10.0.0/16"
    "dev.sub_add"  = "10.10.1.0/24"
    "prd.vnet_add" = "10.11.0.0/16"
    "prd.sub_add"  = "10.11.1.0/24"
  }
}

variable "kubernetes" {
  default = {
    ssh_public_key = "~/.ssh/id_rsa.pub"
    version  = "1.25.5"
    "dev.vnet_add" = "10.10.0.0/16"
    "dev.sub_add"  = "10.10.1.0/24"
    "prd.vnet_add" = "10.11.0.0/16"
    "prd.sub_add"  = "10.11.1.0/24"
  }
}
