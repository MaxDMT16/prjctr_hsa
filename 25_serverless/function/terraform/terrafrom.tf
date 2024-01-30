terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.23.1"
    }

    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.4.0"
    }

    null = {
      source  = "hashicorp/null"
      version = "~> 3.0.0"
    }
  }

  required_version = "~> 1.2"
}

provider "aws" {
  region = var.aws_region
}
