# Protocol Buffers

## Why there is a need for Protocol Buffers?

Why do we need another serialization format? Surely there must be some advantages of using one over the other

### Advantages:
- Typed (The schema defines the type of each field)
    - This might not seem like a major improvement over other serialization formats
    - Because in JSON we also have types
    - Protocol Buffers strongly emphasize the field types
- Generated Code
    - We can generate code for any language
    - This allows us to focus on data schema rather than worrying about the serialization / deserialization
- Schema Evolution
  - Supports the idea that the schema evolves over time
  - If we need evolution, we need backwards/forwards compatability
- Comments
  - Internal documentation unlike JSON
- Binary
  - Data is serialized in binary in an optimal way, not as a string
  - Serialized data that is much smaller, easier to pass and cheaper to send over the network

### Disadvantages:
- Binary serialization
  - Editing / Reading the serialized data is relatively harder than other data exchange formats like JSON / XML

---

## How are protocol buffers used?

In a nutshell, it is used to share data across different programming languages

Steps:
- first we write the schema (protobuf message)
- then the protobuf compiler generates the source code in desired programming language
- then we set the value of the fields in the schema
- then the generated code will serialize that data in binary

### What it is used for?
- It is used for communication (gRPC)
- The gRPC framework uses protocol buffers to serialize/deserialize the data between the client and the server
