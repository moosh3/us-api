variable "aws_region" {
  default     = "us-east-1"
  description = "Default AWS Region"
}

variable "vpc_cidr" {
  default     = "172.20.0.0/18"
  description = "VPC CIDR range"
}
