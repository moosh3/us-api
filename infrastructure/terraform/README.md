# Looker infrastructure

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.12.26 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 2.49 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 3.40.0 |
| <a name="provider_kubernetes"></a> [kubernetes](#provider\_kubernetes) | n/a |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_aws_load_balancer"></a> [aws\_load\_balancer](#module\_aws\_load\_balancer) | git::https://gitlab.amwell.systems/sre/tf-modules/aws-load-balancer.git | v1.0.0 |
| <a name="module_eks"></a> [eks](#module\_eks) | terraform-aws-modules/eks/aws | 17.1.0 |
| <a name="module_external_dns"></a> [external\_dns](#module\_external\_dns) | git::https://gitlab.amwell.systems/sre/tf-modules/aws-external-dns.git | v1.0.0 |
| <a name="module_vpc"></a> [vpc](#module\_vpc) | terraform-aws-modules/vpc/aws | 2.64.0 |

## Resources

| Name | Type |
|------|------|
| [aws_availability_zones.available](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones) | data source |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_eks_cluster.cluster](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/eks_cluster) | data source |
| [aws_eks_cluster_auth.cluster](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/eks_cluster_auth) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | Default AWS Region | `string` | `"us-east-1"` | no |
| <a name="input_database_name"></a> [database\_name](#input\_database\_name) | looker | `string` | `"Initial RDS MySQL database name"` | no |
| <a name="input_database_password"></a> [database\_password](#input\_database\_password) | RDS MySQL database user password | `string` | `"lookerpassword"` | no |
| <a name="input_database_username"></a> [database\_username](#input\_database\_username) | looker | `string` | `"RDS MySQL database username"` | no |
| <a name="input_vpc_cidr"></a> [vpc\_cidr](#input\_vpc\_cidr) | VPC CIDR range | `string` | `"172.20.0.0/18"` | no |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
