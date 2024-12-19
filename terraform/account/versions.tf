terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.82.0"
    }
    github = {
      source  = "integrations/github"
      version = "6.4.0"
    }
  }
  required_version = ">= 1.1.0"
}
