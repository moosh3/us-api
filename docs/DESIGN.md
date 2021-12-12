# Design Decisions

## Database

Store the JSON within the binary due to the size and immutability of the dataset.
The web library, Gin, handles each HTTP request as a separate Goroutine, making the retrieval of the key upon each request extremely fast and allows for up to millions of rps.

## Scalability

Using the default values in the Helm chart, the API is deployed with three replicas,
along with node affinity that ensures each pod is scheduled in a different AZ.
To ensure the service can scale with traffic there exists a HorizontalPodAutoscaler
object that will create a new pod whenever it sees CPU usage over 80%.
