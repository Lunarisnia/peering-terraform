data "aws_security_group" "selected" {
  filter {
    name   = "tag:Name"
    values = ["default"]
  }
}