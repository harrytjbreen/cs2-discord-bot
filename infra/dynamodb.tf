resource "aws_dynamodb_table" "map_last_played" {
  name           = "map_last_played"
  billing_mode   = "PAY_PER_REQUEST"

  attribute {
    name = "id" // SteamID
    type = "S"
  }

  dynamic "attribute" {
    for_each = var.cs2_maps
    content {
      name = attribute.value
      type = "S"
    }
  }

  hash_key = "id"

}

resource "aws_dynamodb_table" "report_list" {
  name = "report_list"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "id" //Random ID
    type = "S"
  }

  attribute {
    name = "banned"
    type = "BOOL"
  }

  attribute {
    name = "timestamp"
    type = "N"
  }

  attribute {
    name = "steamId" //SteamID of the reported player
    type = "S"
  }

  attribute {
    name = "reportedBy" //SteamID of the reporter
    type = "S"
  }

  hash_key = "id"
}
