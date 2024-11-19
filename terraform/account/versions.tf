terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.76.0"
    }
    github = {
      source  = "integrations/github"
      version = "4.31.0"
    }
  }
  required_version = ">= 1.1.0"
}
