terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.88.0"
    }
    github = {
      source  = "integrations/github"
      version = "6.5.0"
    }
  }
  required_version = ">= 1.1.0"
}
