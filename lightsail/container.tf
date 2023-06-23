resource "aws_lightsail_container_service" "pinger-service" {
  name = "pinger-service"
  power = "micro"
  scale = 1
  is_disabled = false
}