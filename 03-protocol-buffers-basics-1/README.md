# 03 - Protocol Buffers Basics 1

## Defaults

When field values are not set:
- Unset fields will not be serialized
  - No payload overhead for empty data
- The side that deserializes the data will use the default values for unset fields

## Scalars
- int (default `0`)
    - int32, int64, sint32, sint64
    - uint32, uint64
    - fixed32, fixed64, sfixed32, sfixed64
    - float32, double
- bool
  - true, false (default `false`)
- string
  - keyword: `string`
  - length: arbitrary length text
  - default: empty string
  - Accepts only UTF-8 encoded strings OR 7-bit ASCII encoded strings
- bytes
  - keyword: `bytes`
  - length: arbitrary length byte sequence
  - default: empty array of bytes
  - It's upto the programmer to interpret the contents of bytes
    - a jpeg image might be serialized in a protobuf
    - once we receive those bytes in our code, we need to treat them as an image
    - otherwise the data will be meaningless

## Tags
- Tags make serialization and deserialization possible
- Field names are not important for serialization
- Rules:
  - Tags range is (1 to 536,870,911)
  - Google reserved tags (19000 to 19999)
- Payload is affected by the tag
- `1 to 15` -> 1 byte
- `16 to 2047` -> 2 byte
- ...
- This makes it a good practice to use smaller tags for the most used / populated fields in the schema
- Saves bandwidth / resources

## Repeated Fields
We can also have lists as our types
- Keyword: `repeated <type> <name> = <tag>;`
- Value: we can have any number of values for a repeated field
- Default value: empty list

## Enums
- keyword: `enum`
- Default Value: First value that is defined with enum
- First tag of an `enum` should be `0`