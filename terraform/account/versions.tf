terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.100.0"
    }
    github = {
      source  = "integrations/github"
      version = "6.11.1"
    }
  }
  required_version = "1.14.6"
}
