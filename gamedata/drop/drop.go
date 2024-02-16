package drop

// implement item drops, provided 2 types of drop lists:
// - equal weight random
// - weighted random
// each type provides the following methods:
// - AddItem(content client.Content) / AddItem(content client.Content, weight int32)
//   - add an item to the random list, the state is guaranteed to be correct and useable after each call
//   - the weight params is only present for weighted random.
// - GetRandomDrop() client.Content {}
//   - returns a single random item
//   - for equal weight random, this is O(1)
//   - for weighted random, this is O(log(n)) with n being the amount of different drops.
// - GetRandomDrops(count int32) []client.Content {}:
//   - returns a slice of random items
//   - this is like calling the above "count" times, so it's either O(nlog(n)) or O(n)
