# Description

Really simple hash table that I have scanned from the Internet (algorithm).

* Implements interface{} keys and values so we can hash by integer and string.
* Value can take custom type, e.g. type Org struct {}, so we can return custom values
* Hashing is O(n) complexity, where `n` is a length of the key.
