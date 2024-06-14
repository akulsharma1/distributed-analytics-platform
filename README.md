# distributed-analytics-platform




## Kafka Setup
In order to use the platform, you need to set up Kafka.

First, install [Docker](https://docs.docker.com/engine/install/) if you haven't already.
Then, do all the following commands in your terminal:

1. `cd kafka`
2. `docker-compose up -d`

Once it is done running, use the command `docker ps` to check on the status.

Use the command `docker-compose down` to stop the kafka cluster.