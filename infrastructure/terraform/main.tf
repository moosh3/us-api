provider "aws" {
  region = var.aws_region
}

locals {
  cluster_name = "example"
}

data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_id
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  token                  = data.aws_eks_cluster_auth.cluster.token
}

data "aws_availability_zones" "available" {}

data "aws_caller_identity" "current" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "2.64.0"

  name                 = "test-vpc"
  cidr                 = "10.0.0.0/16"
  azs                  = data.aws_availability_zones.available.names
  public_subnets       = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  enable_dns_hostnames = true

  public_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/elb"                      = "1"
  }
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "17.1.0"

  cluster_name    = local.cluster_name
  cluster_version = "1.21"
  subnets         = module.vpc.public_subnets
  vpc_id          = module.vpc.vpc_id
  enable_irsa     = true

  node_groups = {
    ondemand_instances = {
      desired_capacity = 3
      max_capacity     = 5
      min_capacity     = 1

      instance_types = ["m5.xlarge"]
      k8s_labels = {
        Environment = "production"
        ManagedBy   = "terraform"
      }
    }
  }
}

module "aws_load_balancer" {
  source = "git::https://XXXXXXXXXXXX/sre/tf-modules/aws-load-balancer.git?ref=v1.0.0"

  cluster_oidc_issuer_url = data.aws_eks_cluster.cluster.identity.0.oidc.0.issuer
  cluster_name            = data.aws_eks_cluster.cluster.id
}

module "external_dns" {
  source = "git::https://XXXXXXXXXXXX/sre/tf-modules/aws-external-dns.git?ref=v1.0.0"

  cluster_oidc_issuer_url = data.aws_eks_cluster.cluster.identity.0.oidc.0.issuer
  cluster_name            = data.aws_eks_cluster.cluster.id
}
