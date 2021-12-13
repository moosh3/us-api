# us-api

us-api is a simple service that returns the US state code based on the state. It
does not support creating, updating nor deleting data.

## Local Development

```
docker build -t us-api .
docker run -d -p 8080:8080 us-api
```

## Deploying

The `infrastructure` directory contains the terraform required to stand up an
EKS cluster within AWS. It additionally creates a VPC of which to deploy the cluster
to.

Given you have the correct AWS access, you can run the following to stand up a cluster:

```
cd infrastructure/terraform
terraform init
terraform plan
terraform apply
```

Once you have access to a kubernetes cluster, you can use Helm chart within the `chart` directory to deploy the actual service.

```
helm upgrade -i us-api ./chart/us-api
```

### HA

HA is achieved by running a stateless replica on three nodes across AZ's, along with
a HorizontalPodAutoscaler object that scales the deployment out based on resource usage.
