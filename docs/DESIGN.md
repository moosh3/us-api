# Design Decisions

## Database

Two implementations exist within this repository:

1. Store the JSON within the binary due to the size and immutability of the dataset
2. Use go-cache to store the data in a memory efficient, scalable embedded cache 

## Scalability

Using the default values in the Helm chart, the API is deployed with three replicas,
along with node affinity that ensures each pod is scheduled in a different AZ.
To ensure the service can scale with traffic there exists a HorizontalPodAutoscaler
object that will create a new pod whenever it sees CPU usage over 80%.
