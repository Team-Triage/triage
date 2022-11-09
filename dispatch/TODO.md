- When the connection closes (due to violation of keepalive params)
  - The downstream consumer instance is 'dead'
  
  - Currently, keepalive does not seem to be working as intended; however,  downstream consumer death throws an error on the next gRPC method invocation.
    - This is suitable for our needs, since we still have the message in scope when we receive this error, and can append to the messages channel and break out of the senderRoutine
  - If the consumer instance dies before we send them a message, we'll get the same error.
  - If the consumer dies after we send them a message, but before we get a response, we'll get the same error.

  
  - Place the message in the messages channel
  - Break the goroutine


- When we get an error because we've violated the TIMEOUT on the individual message
  - We are in our 'just in case' scenario where we treat the message as a poison pill (potentially make this timeout configurable as a nice to have)
  - Send the message to filter as a nack.