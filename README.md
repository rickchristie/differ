# Differ

Used to generate human-readable diff of structs, maps, lists. Use cases:
- You need to create an audit log. You can just pass the before and after structs to `differ` and it will generate
  human-readable diff of what has changed for you. No need to code the diff by hand.
- You have a bug, so you need to log changes within an entity to make troubleshooting easier.

Alternatives:
- Just use `json.Marshal` and compare the before and after using text diff algorithm. This is good for most use cases.
  When you have `map[comparable]struct` or `struct[]` with many values, and only 1 field change in one of the items,
  it's difficult to know which item actually changed.

## Implementation

Implementation is very simple, but it's not the fastest.
- Use reflection to convert everything to 