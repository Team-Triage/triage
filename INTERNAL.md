Development environment setup:

IMPORTANT: Kill main process with Ctrl + \ , NOT Ctrl + C

1. Kafka
  - From /triage
  - Move to the /dev/kafka directory
    - cd /dev/kafka
  - Start kafka cluster using docker compose
    - docker-compose up -d

2. Producing test messages to Kafka
  - From /triage
  - Move to the /dev/tmp
    - cd /dev/tmp
  - Run the producer script
    - ./producer devConfig.properties

3. Consuming test messages from Kafka
  - From /triage
    - go run main.go