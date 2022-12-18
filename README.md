# Triage
Triage is an open source Kafka consumer proxy service that solves head-of-line blocking while preventing data-loss through increased parallelism, commit tracking, and a dead-letter store. 

Triage operates as an AWS Fargate service deployed to a users AWS account. It sits between your Kafka cluster and consumer applications and can be interacted with via a thin client library. 

To read more about Triage, visit: https://team-triage.github.io/

<h2>How to Use Triage</h2>


Users can deploy Triage using the `triage-cli` command line tool, which is available via `npm`. It can be installed using the following command:

	npm install -g triage-cli
For detailed instructions, see the README at: https://github.com/Team-Triage/triage-cli

To interact with Triage, consumer applications should use the `triage-client-go` library, found at https://github.com/Team-Triage/triage-client-go.
