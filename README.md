# Triage

Weekend Todos:

Go Get/Module Management - Mike, Aashish
  - How do we publish a module so a developer can retrieve using go get

Understand gRPC liveness - Jordan, Aashish
  - What happens if one end of the gRPC connection dies?

Securing gRPC connection - Jordan, Aashish

Securing connection to Kafka Cluster - Mike
  - How do we authenticate with the Cluster?
    - What are the most common auth methods?
    - How do we integrate said methods with Triage?

Wait Groups - Aryan, Mike
  - How to establish/manage wait groups correctly

Lightest/simplest HTTP Server for Go - Aashish, Aryan


Split:
  - Reaper
  - Consumer Manager
  - Commit Calculator

Low Priority:
  - Channel direction management

Redo: Diagram

Error Handling:

Dispatch Function:

  1. Consumer Connection - what if network address in Dispatch function is invalid?? 
    - return an error saying "This is invalid" 
  2. Add a kill channel for sendMessage goroutine