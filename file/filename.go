package file

type Filename string

// TODO: This needs to be definable by the developer, probably default to xhash,
// but this needs to be one of those definable actions. The idea is that was may
// even store everything as this hash an just run the hash the other way to get
// the Filename value. This allows for all comparisons to be bytes instead of
// strings.
func (self Filename) Hash() []byte {

}

// TODO: Given a merkle root, we compare determine if the filename belongs in
// the merkle tree.
func (self Filename) ValidMerkleRoot(root string) bool {
	return false
}
