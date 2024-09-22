# Differ

Used to generate human-readable diff of structs, maps, lists. Use cases:
- You need to create an audit log. You can just pass the before and after structs to `differ` and it will generate
  human-readable diff of what has changed for you. No need to code the diff by hand.
- You have a bug, so you need to log changes within an entity to make troubleshooting easier.

## Implementation

Current implementation uses `json.Marshal` internally. It's not the fastest, and has issues, like not supporting types 
like `any` or `int` keys on `map`, but should be good enough for most domain objects. Plan is to improve this by 
changing implementation to use reflection entirely. 
