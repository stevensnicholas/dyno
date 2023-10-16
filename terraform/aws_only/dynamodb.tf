resource "aws_dynamodb_table" "dyno_table"{
  name  = "Dyno"
  hash_key = "clientId"
  range_key = "fuzzId"
  attribute {
    name = "clientId"
    type = "S"
  }
  attribute {
    name = "fuzzId"
    type = "S"
  }
}