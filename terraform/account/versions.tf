terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.76.0"
    }
    github = {
      source  = "integrations/github"
      version = "6.3.1"
    }
  }
  required_version = ">= 1.1.0"
}
