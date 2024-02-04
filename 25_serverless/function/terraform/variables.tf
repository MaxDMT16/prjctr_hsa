variable "aws_region" {
  description = "AWS region for all resources."
  type        = string
  default     = "eu-west-1"
}

variable "bucket_name" {
  description = "Name of the bucket"
  type        = string
}
