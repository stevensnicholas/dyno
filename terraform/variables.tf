variable "deployment_id" {
  type        = string
  description = "main-api"
}

variable "restler_image_tag" {
  type        = string
  description = "restler image tag to pull, if this is the first run then you need to supply an image from another ecr"
}
