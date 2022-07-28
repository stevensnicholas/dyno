resource "tls_private_key" "RSAprikey" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_ssm_parameter" "prikey" {
  name      = "prikey"
  type      = "SecureString"
  value     = tls_private_key.RSAprikey.private_key_pem
  overwrite = true
}