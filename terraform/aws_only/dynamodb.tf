resource "aws_dynamodb_table" "dyno_table" {
  name = "${var.deployment_id}-fuzz-results"
  hash_key       = "clientID"
  attribute {
    name = "clientID"
    type = "S"
  }
}