# Provider Configuration
provider "aws" {
  region = var.aws_region
}

# Resource Configuration
resource "aws_instance" "webserver" {
  ami           = var.ami_id
  instance_type = var.instance_type

  tags = {
    Name = "Web Server"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y apache2",
      "sudo systemctl enable apache2",
      "sudo systemctl start apache2",
      "sudo a2enmod rewrite",
      "echo '<html><head><title>Hello World</title></head><body><h1>Hello World!</h1></body></html>' | sudo tee /var/www/html/index.html",
      "sudo systemctl restart apache2"
    ]
  }
}

# HTTPS Listener Configuration
resource "aws_lb_listener" "https_listener" {
  load_balancer_arn = aws_lb.webserver_lb.arn
  port              = 443
  protocol          = "HTTPS"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.webserver_target_group.arn
  }

  ssl_policy      = "ELBSecurityPolicy-2016-08"
  certificate_arn = var.certificate_arn
}

# Target Group Configuration
resource "aws_lb_target_group" "webserver_target_group" {
  name        = "webserver-target-group"
  port        = 80
  protocol    = "HTTP"
  vpc_id      = var.vpc_id
  target_type = "instance"

  health_check {
    healthy_threshold   = 2
    unhealthy_threshold = 2
    interval            = 30
    path                = "/"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
  }
}

# Load Balancer Configuration
resource "aws_lb" "webserver_lb" {
  name               = "webserver-lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.webserver_sg.id]
  subnets            = var.subnet_ids
}

# Security Group Configuration
resource "aws_security_group" "webserver_sg" {
  name        = "webserver-sg"
  description = "Web server security group"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Outputs
output "webserver_url" {
  value = "http://${aws_instance.webserver.public_ip}"
}
