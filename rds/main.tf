resource "aws_db_instance" "gopinger_db" {
    identifier = "gopinger-database-2"
    allocated_storage = 10
    db_name = "gopinger"
    engine = "postgres"
    engine_version = "13.7"
    instance_class = "db.t3.micro"
    username = "postgres"
    password = "password1234"
    publicly_accessible = false
    vpc_security_group_ids = [data.aws_security_group.selected.id]
    # db_subnet_group_name = "default-vpc-011d1fc2f111c2eb8"
    # final_snapshot_identifier = "snap1"
    skip_final_snapshot = true
}