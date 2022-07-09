variable "deployment_id" {
  type        = string
  description = "main-api"
  default = "main"
}

variable "is_local" {
  type        = bool
  description = "Set true if local"
  default = false
}
