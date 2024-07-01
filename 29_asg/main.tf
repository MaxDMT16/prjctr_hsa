locals {
  ami = {
    ubuntu = "ami-0905a3c97561e0b69"
  }

  instance_type = "t2.micro"
}

resource "aws_launch_template" "main" {
  name_prefix = "${local.instance_type}-ubuntu"
  image_id = local.ami.ubuntu
  instance_type = local.instance_type
}

resource "aws_autoscaling_group" "main" {
  depends_on = [ aws_launch_template.main ]

  max_size = 3
  min_size = 1
  
  launch_template {
    id = aws_launch_template.main.id
    version = "$Latest"
  }

  availability_zones = [ "eu-west-1a" ]

  tag {
    key = "Name"
    value = "prjctr-terraform-asg"
    propagate_at_launch = true
  }
}

# resource "aws_launch_template" "template" {
#   name_prefix     = "test"
#   image_id        = "ami-1a2b3c"
#   instance_type   = "t2.micro"
#   vpc_security_group_ids = ["sg-12345678"]
# }

# resource "aws_autoscaling_group" "autoscale" {
#   name                  = "test-autoscaling-group"  
#   availability_zones    = ["us-west-2"]
#   desired_capacity      = 1
#   max_size              = 6
#   min_size              = 1
#   health_check_type     = "EC2"
#   termination_policies  = ["OldestInstance"]
#   vpc_zone_identifier   = ["subnet-12345678"]

#   launch_template {
#     id      = aws_launch_template.template.id
#     version = "$Latest"
#   }
# }

# resource "aws_autoscaling_policy" "scale_down" {
#   name                   = "test_scale_down"
#   autoscaling_group_name = aws_autoscaling_group.autoscale.name
#   adjustment_type        = "ChangeInCapacity"
#   scaling_adjustment     = -1
#   cooldown               = 120
# }

# resource "aws_cloudwatch_metric_alarm" "scale_down" {
#   alarm_description   = "Monitors CPU utilization"
#   alarm_actions       = [aws_autoscaling_policy.scale_down.arn]
#   alarm_name          = "test_scale_down"
#   comparison_operator = "LessThanOrEqualToThreshold"
#   namespace           = "AWS/EC2"
#   metric_name         = "CPUUtilization"
#   threshold           = "25"
#   evaluation_periods  = "5"
#   period              = "30"
#   statistic           = "Average"

#   dimensions = {
#     AutoScalingGroupName = aws_autoscaling_group.autoscale.name
#   }
# }
