# Terraform Web Server Deployment

This project demonstrates the deployment of a web server using Terraform. The web server serves a simple "Hello World" page over HTTPS. The infrastructure is provisioned on AWS.

## Prerequisites

Before you begin, ensure that you have the following:

- Terraform installed on your local machine
- AWS account credentials configured
- An SSL certificate for the web server (self-signed or issued by a trusted CA)

## Project Structure

The project structure is organized as follows:

- `main.tf`: Defines the infrastructure resources (VPC, subnet, security group, EC2 instance).
- `variables.tf`: Contains the variables used in the Terraform configuration.
- `outputs.tf`: Specifies the output values exposed by the Terraform deployment.
- `tests/web_server_test.tf`: Defines automated tests to validate the correctness of the server configuration.
- `certs/private_key.pem`: Contains the private key for the SSL certificate.
- `certs/certificate.pem`: Contains the SSL certificate.
- `README.md`: Provides instructions and explanations for the project.

## Getting Started

1. Clone the repository:

   ```shell
   git clone <repository-url>

2. Navigate to the project directory:

cd terraform-web-server

3. Update the variables.tf file with your desired configuration.

4. Place your SSL certificate files (private_key.pem and certificate.pem) in the certs directory.

5. Initialize the Terraform project:

terraform init

6. Deploy the infrastructure:

terraform apply

7. After the deployment is complete, the web server URL will be displayed in the output.

##Testing
To run the automated tests, use the following command:

terraform test

The tests will validate the deployment of the web server and verify the web page content.
