package hash

// https://github.com/umpc/go-zrange
// How about modifying this type of hash but for different types of filters?
// For example, we do long is directory, increasing with depth, and lat is a
// filter inside the folder, can be alpha, can be size, etc. We can have
// multiple versions for each type. Then we radix it.

//rangeParams := zrange.RadialRangeParams{
//  Radius:    (maybe this is filetype actually)
//  Directory: (BitSet based on pathing from the root path),
//  Sort:      (Sort of a specific type, like alpha, or size, or filetype),
//}
