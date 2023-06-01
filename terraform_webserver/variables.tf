variable "aws_region" {
  description = "AWS region where the web server will be deployed"
  default     = "us-west-2"
}

variable "ami_id" {
  description = "AMI ID for the web server instance"
  default     = "ami-0c94855ba95c71c99"  # Replace with your desired AMI ID
}

variable "instance_type" {
  description = "EC2 instance type for the web server"
  default     = "t2.micro"
}

variable "vpc_id" {
  description = "VPC ID where the web server will be deployed"
  default     = "vpc-0123456789abcdef0"  # Replace with your desired VPC ID
}

variable "subnet_ids" {
  description = "List of subnet IDs where the web server will be deployed"
  type        = list(string)
  default     = ["subnet-0123456789abcdef0"]  # Replace with your desired subnet IDs
}

variable "certificate_arn" {
  description = "ARN of the SSL certificate for HTTPS"
  default     = "arn:aws:acm:us-west-2:123456789012:certificate/abcdef12-3456-7890-abcd-ef1234567890"  # Replace with your desired certificate ARN
}
