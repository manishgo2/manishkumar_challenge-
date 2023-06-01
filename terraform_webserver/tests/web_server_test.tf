terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.0.0"
    }
  }
}

terraform_test "web_server_test" {
  test_folder = "./tests"
}
