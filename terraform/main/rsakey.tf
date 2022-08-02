resource "tls_private_key" "rsa_prikey" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_ssm_parameter" "private_key" {
  name      = "prikey"
  type      = "String"
  value     = tls_private_key.RSAprikey.private_key_pem
  overwrite = true
}