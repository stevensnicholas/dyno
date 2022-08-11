resource "tls_private_key" "rsa_prikey" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_ssm_parameter" "private_key" {
  name      = "${var.deployment_id}-prikey"
  type      = "String"
  value     = tls_private_key.rsa_prikey.private_key_pem
  overwrite = true
}

resource "aws_ssm_parameter" "public_key" {
  name      = "${var.deployment_id}-pubkey"
  type      = "String"
  value     = tls_private_key.rsa_prikey.public_key_pem
  overwrite = true
}

resource "aws_ssm_parameter" "client_id" {
  name      = "${var.deployment_id}-client_id"
  type      = "String"
  value     = var.client_id
  overwrite = true
}

resource "aws_ssm_parameter" "client_secret" {
  name      = "${var.deployment_id}-client_secret"
  type      = "String"
  value     = var.client_secret
  overwrite = true
}