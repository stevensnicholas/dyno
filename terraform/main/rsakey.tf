resource "tls_private_key" "rsa_prikey" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_ssm_parameter" "private_key" {
  name      = "${var.deployment_id}-prikey"
  type      = "String"
  value     = tls_private_key.res_prikey.private_key_pem
  overwrite = true
}