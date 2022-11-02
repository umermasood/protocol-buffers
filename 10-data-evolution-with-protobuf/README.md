# 10. Data Evolution with Protobuf

Data Evolution is one of the major design decisions that were made in Protocol Buffers

We generally need:
- Backward compatability
- Forward compatability

Rules about changing Protobuf: 
- Don't change tags
- Add new fields
- Use reserved tags
- Before changing types:
  - Refer to official documentation
  - If the type compatability is trivial - change the type
  - OR
  - Add a new field