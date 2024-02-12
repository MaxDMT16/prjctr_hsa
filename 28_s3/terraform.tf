terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.23.1"
    }
  }

  required_version = "~> 1.2"
}

provider "aws" {
  region = var.aws_region
}
