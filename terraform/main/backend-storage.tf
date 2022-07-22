resource "aws_s3_bucket" "backend_data_store" {
  bucket = "${var.deployment_id}-fuzzing-results-comp9447"
}

resource "aws_s3_bucket_versioning" "backend_data_store" {
  bucket = aws_s3_bucket.backend_data_store.id
  versioning_configuration {
    status = "Enabled"
  }
}
resource "aws_s3_bucket_public_access_block" "backend_data_store" {
  bucket              = aws_s3_bucket.backend_data_store.id
  block_public_acls   = true
  block_public_policy = true
}