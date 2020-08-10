# UPP - Content Collection Unfolder

Content Collection Unfolder finds added/deleted collection members and lead articles through the `relations-api`.

## Code

content-collection-unfolder

## Primary URL

https://upp-prod-delivery-glb.upp.ft.com/__content-collection-unfolder/

## Service Tier

Platinum

## Lifecycle Stage

Production

## Delivered By

content

## Supported By

content

## Known About By

- dimitar.terziev
- elitsa.pavlova
- kalin.arsov
- hristo.georgiev
- georgi.ivanov
- elina.kaneva
- robert.marinov
- tsvetan.dimitrov

## Host Platform

AWS

## Architecture

This service forwards mapped content collections to `content-collection-rw-neo4j`. If the response is 200, it retrieves the elements in the collection from `document-store-api` and places them in Kafka on the `PostPublicationEvents` topic so that notifications will be created for each one of them.

## Contains Personal Data

No

## Contains Sensitive Data

No

## Dependencies

* upp-relations-api
* upp-content-collection-rw-neo4j
* document-store-api
* upp-kafka

## Failover Architecture Type

ActiveActive

## Failover Process Type

FullyAutomated

## Failback Process Type

FullyAutomated

## Failover Details

The service is deployed in both Delivery clusters. The failover guide for the cluster is located here: https://github.com/Financial-Times/upp-docs/tree/master/failover-guides/delivery-cluster.

## Data Recovery Process Type

NotApplicable

## Data Recovery Details

The service does not store data, so it does not require any data recovery steps.

## Release Process Type

PartiallyAutomated

## Rollback Process Type

Manual

## Release Details

The release is triggered by making a Github release which is then picked up by a Jenkins multibranch pipeline. The Jenkins pipeline should be manually started in order for it to deploy the helm package to the Kubernetes clusters.

## Key Management Process Type

NotApplicable

## Key Management Details

There is no key rotation procedure for this system.

## Monitoring

Service in UPP K8S delivery clusters:

* Delivery-Prod-EU health: https://upp-prod-delivery-eu.upp.ft.com/__health/__pods-health?service-name=content-collection-unfolder
* Delivery-Prod-US health: https://upp-prod-delivery-us.upp.ft.com/__health/__pods-health?service-name=content-collection-unfolder

## First Line Troubleshooting

[First Line Troubleshooting guide](https://github.com/Financial-Times/upp-docs/tree/master/guides/ops/first-line-troubleshooting)

## Second Line Troubleshooting

Please refer to the GitHub repository README for troubleshooting information.