# Due to bugs with CS2 we cannot get this information right now
# Will be added in the future

#resource "aws_dynamodb_table" "map_last_played" {
#  name           = "map_last_played"
#  billing_mode   = "PAY_PER_REQUEST"
#
#  attribute {
#    name = "id" // SteamID
#    type = "S"
#  }
#
#  attribute {
#    name = "match-code"
#    type = "S"
#  }
#
#  attribute {
#    name = "auth-token"
#    type = "S"
#  }
#
#  dynamic "attribute" {
#    for_each = var.cs2_maps
#    content {
#      name = attribute.value
#      type = "S"
#    }
#  }
#
#  hash_key = "id"
#
#}

resource "aws_dynamodb_table" "report_list" {
  name = "report_list"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "id" //Random ID
    type = "S"
  }

  attribute {
    name = "steamId" //SteamID of the reported player
    type = "S"
  }

  attribute {
    name = "reportedBy" //SteamID of the reporter
    type = "S"
  }

  global_secondary_index {
    name = "SteamIdIndex"
    hash_key = "steamId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name = "SteamIdIndex"
    hash_key = "reportedBy"
    projection_type = "ALL"
  }

  hash_key = "id"
}
