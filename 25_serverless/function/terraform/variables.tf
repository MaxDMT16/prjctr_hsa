variable "aws_region" {
  description = "AWS region for all resources."
  type        = string
  default     = "eu-west-1"
}

variable "bucket_name" {
  description = "Name of the bucket"
  type        = string
}

variable "force_destroy_bucket" {
  description = "A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error."
  type        = bool
  default     = false  
}
