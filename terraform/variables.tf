variable "deployment_id" {
  type        = string
  description = "main-api"
}

variable "restler_image_tag" {
  type        = string
  default = "117712065617.dkr.ecr.ap-southeast-2.amazonaws.com/main_dyno_image_repository:9270505710dab7540a8340efc24fef311b75552e"
  description = <<EOT
  restler image tag to pull, if this is the first run then you need to supply an image from another ecr
  For example to get the current latest pushed image on main (production) run the following comand
  echo 117712065617.dkr.ecr.ap-southeast-2.amazonaws.com/main_dyno_image_repository:`aws ecr describe-images --repository-name main_dyno_image_repository --query 'sort_by(imageDetails,& imagePushedAt)[-1].imageTags[0]' | tr -d '"'`
  EOT
}
