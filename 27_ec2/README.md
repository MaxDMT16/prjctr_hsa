# EC2 & Load Balancer

## How to setup
**Prerequisites**

- AWS account
- configured AWS CLI
- Terraform

Run the following commands:
```bash
terraform init
terraform apply
```

## How it works

There is a private VPC. It contains 2 subnets, one for each availability zone. The subnets are private, so the EC2 instances are not accessible from the internet. The EC2 instances are placed behind an Application Load Balancer (ALB) which is placed in a public subnet. The ALB is accessible from the internet and forwards the requests to the EC2 instances.

EC2's security groups have rules to allow bidirectional traffic from the ALB (port `:8080`), outbound traffic to the internet to install updates and usefull packages (e.g. Nginx) and inbound trafic from the Ubuntu repository. 


## Additional Resources:
- [setup EC2 instances with Application Load Balancer (ALB)](https://antonputra.com/amazon/create-alb-terraform/#create-aws-alb-with-ec2-backend)
- 
