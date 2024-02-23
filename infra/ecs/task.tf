resource "aws_ecs_task_definition" "discord_bot_task" {
  family                = "discord-bot-task"
  container_definitions = jsonencode([
    {
      name      = "discord-bot"
      image     = "nginx:latest"
      cpu       = 0.25
      memory    = 0.5
      essential = true
    }
  ])
}
