# SQS Terraform Setup 

resource "aws_sqs_queue" "fuzz_queue" {
  name                      = "${var.deployment.id}-fuzz-queue-test" #TODO remove test when ready to deploy
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.terraform_queue_deadletter.arn
    maxReceiveCount     = 4
  })
  redrive_allow_policy = jsonencode({
    redrivePermission = "byQueue",
    sourceQueueArns   = [aws_sqs_queue.terraform_queue_deadletter.arn]
  })

  tags = {
    Environment = "production"
  }
}