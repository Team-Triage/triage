# Triage

Todo's Tomorrow:

Mob: Dispatch:
  - Implement logic to invoke SendMessage goRoutines for each downstream consumer instance(network address)
    - Listen on channel from Consumer Manager (to be implemented, as well)
    - When new network address is read off channel, create new gRPC connection/client
    - call SendMessage goRoutine with respective gRPC client

Split: Consumer Manager & Reaper