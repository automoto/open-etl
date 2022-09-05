# open-etl
Open source implementations and architecture for different ETL workloads.

## sqs

to setup for first time create a queue:
`aws --endpoint http://localhost:4556 sqs create-queue --queue-name results-queue --attributes MessageRetentionPeriod=1209600,ContentBasedDeduplication=false`