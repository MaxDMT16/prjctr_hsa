resource "aws_lb_target_group" "tg1" {
  name       = "tg1"
  port       = 8080
  protocol   = "HTTP"
  vpc_id     = aws_vpc.prjctr.id
  slow_start = 0

  load_balancing_algorithm_type = "round_robin"

  stickiness {
    enabled = false
    type    = "lb_cookie"
  }


  # TODO: make a health check endpoint

  #   health_check {
  #     enabled             = true
  #     port                = 8081
  #     interval            = 30
  #     protocol            = "HTTP"
  #     path                = "/health"
  #     matcher             = "200"
  #     healthy_threshold   = 3
  #     unhealthy_threshold = 3
  #   }
}

resource "aws_lb_target_group_attachment" "my_app_tg1" {
  for_each = aws_instance.app_server

  target_group_arn = aws_lb_target_group.tg1.arn
  target_id        = each.value.id
  port             = 8080
}

resource "aws_lb" "lb1" {
  name               = "lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb.id]

  # access_logs {
  #   bucket  = aws_s3_bucket.lb-logs.bucket
  #   prefix  = "my-app-lb"
  #   enabled = true
  # }

  subnets = [
    aws_subnet.prjctr_public_a.id,
    aws_subnet.prjctr_public_b.id
  ]
}

resource "aws_lb_listener" "http_eg1" {
  load_balancer_arn = aws_lb.lb1.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.tg1.arn
  }
}
