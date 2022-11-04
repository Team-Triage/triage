# Triage

Todo's Tomorrow:

Mob: Dispatch:
  - Implement logic to invoke SendMessage goRoutines for each downstream consumer instance(network address)
    - Listen on channel from Consumer Manager (to be implemented, as well)
    - When new network address is read off channel, create new gRPC connection/client
    - call SendMessage goRoutine with respective gRPC client

Split:
  - Thin client
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