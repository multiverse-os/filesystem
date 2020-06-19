package hash

import "hash"

type crc32Hasher struct {
	crc32 hash.Hash32
}

func (self *crc32Hasher) Write(p []byte) (n int, err error) { return self.crc32.Write(p) }
func (self *crc32Hasher) Sum(b []byte) []byte               { return self.crc32.Sum(b) }
func (self *crc32Hasher) Reset()                            { self.crc32.Reset() }
func (self *crc32Hasher) Size() int                         { return self.crc32.Size() }
func (self *crc32Hasher) BlockSize() int                    { return self.crc32.BlockSize() }
func (self *crc32Hasher) Sum32() uint32                     { return self.crc32.Sum32() }
func (self *crc32Hasher) Sum64() uint64                     { return uint64(self.crc32.Sum32()) }

var _ hash.Hash32 = (*crc32Hasher)(nil)
var _ hash.Hash64 = (*crc32Hasher)(nil)
