variable "dyno_table"{
  type = string
  description = "Table for dyno app"
}

resource "aws_dynamodb" "dyno_table"{
  name = var.dyno_table
  attribute {
    name = "clientID"
    type = "S"
  }
  attribute {
    name = "Title"
    type = "S"
  }
  attribute {
    name = "Details"
    type = "S"
  }
  attribute {
    name = "Visualizer"
    type = "S"
  }
  attribute {
    name = "Body"
    type = ""
  }
  attribute {
    name = "Assignee"
    type = "S"
  }
  atrribute {
    name = "Labels"
    type = "SS"
  }

  atrribute {
    name = "State"
    type = "S"
  }
  atrribute {
    name = "Milestone"
    type = "N"
  }
}

