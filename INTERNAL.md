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


Changes:

Commmit table is now a hashmap under the data/commitTable/commitTable.go file. Since we don't actually need to know
if a message was 'acked' vs 'nacked', (we only need to know if we've received a response for a message), we've implemented
the hashmap as a map[int]bool, where the keys of the map are offsets (integers) and the status of the message (whether its been acked OR nacked) is a boolean.

