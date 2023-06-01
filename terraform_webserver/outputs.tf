output "webserver_url" {
  description = "URL of the web server"
  value       = aws_instance.webserver.public_ip
}
