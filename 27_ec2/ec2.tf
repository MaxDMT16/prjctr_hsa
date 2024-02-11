locals {
  ami = {
    ubuntu = "ami-0905a3c97561e0b69"
  }

  servers = {
    server-1 = {
      ami            = local.ami.ubuntu
      instance_type  = "t2.micro"
      subnet_id      = aws_subnet.prjctr_public_a.id
      user_data_path = "scripts/server1.tpl"
    },
    server-2 = {
      ami            = local.ami.ubuntu
      instance_type  = "t2.micro"
      subnet_id      = aws_subnet.prjctr_public_b.id
      user_data_path = "scripts/server2.tpl"
    }
  }
}

resource "aws_key_pair" "ec2" {
  key_name   = "prjctr-terraform-ec2"
  public_key = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMvNfw1zJaZlI9H0oWk8I8lcoe3aXD7bze5cXCXMgZcr prjctr-aws-ec2-key"
}

resource "aws_instance" "app_server" {
  for_each = local.servers

  ami           = each.value.ami
  instance_type = each.value.instance_type
  subnet_id     = each.value.subnet_id
  key_name      = aws_key_pair.ec2.key_name

  vpc_security_group_ids = [aws_security_group.ec2.id]

  user_data = file(each.value.user_data_path)

  tags = {
    project = "prjctr_ec2"
    Name    = each.key
  }
}


# TMP             TODO: delete after testing

# Default VPC
# resource "aws_default_vpc" "tmp" {}

# # Security group
# resource "aws_security_group" "tmp_sg" {
#   name        = "demo_sg"
#   description = "allow ssh on 22 & http on port 80"
#   vpc_id      = aws_default_vpc.tmp.id

#   ingress {
#     from_port   = 22
#     to_port     = 22
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   ingress {
#     from_port   = 80
#     to_port     = 80
#     protocol    = "tcp"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }

# resource "aws_instance" "tmp1" {
#   ami           = local.ami.ubuntu
#   instance_type = "t2.micro"

#   key_name = aws_key_pair.ec2.key_name

#   # subnet_id              = aws_subnet.prjctr_public_a.id
#   # vpc_security_group_ids = [aws_security_group.ec2.id]

#   user_data = file("scripts/server1.tpl")

#   tags = {
#     "Name" = "tmp"
#   }
# }
