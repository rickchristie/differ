package differ

// For list, with long-list use cases:
//  - Zero item list, new item.
//  - Only 1 or 2 fields changed within the list.
//  - 1 item is inserted at the top.
//     - ...
//	- 1 item is inserted in the middle.
//     - All the same fields as the previous item.
//     - All the same fields as the next item.
//     - Only 1-2 fields are different from the previous item.
//     - Only 1-2 fields are different from the next item.
//     - All fields are different from the previous and next items.
//     - Fields are all the same with another item in the list.
//         - Equal with an item at the top.
//         - Equal with an item at the bottom.
//         - Equal with an item at the middle (2 before, 2 after).
//	- 1 item is inserted at the bottom.
//     - ...
//  - Top item(s) is deleted.
//  - Middle item(s) is deleted.
//  - Bottom item(s) is deleted.
//  - Item is deleted until it becomes a zero list item.

// For maps, use cases:
//   - Key is updated.
//   - Key is deleted.
//   - Value A and Value B's keys are swapped.

// Multi-structure tests:
//   - Struct
//     - Map of Struct
//        - Map of fields.
//        - List of structs
//        - List of fields.
//        - Map of Structs
//     - List of Structs
//        - List of values.
//        - List of fields.
//        - Map of fields.
//        - Map of structs.
