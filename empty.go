// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <https://unlicense.org>

package verkle

import "errors"

type Empty struct{}

func (Empty) Insert([]byte, []byte, NodeResolverFn) error {
	return errors.New("an empty node should not be inserted directly into")
}

func (e Empty) InsertOrdered(key []byte, value []byte, _ NodeFlushFn) error {
	return e.Insert(key, value, nil)
}

func (Empty) Delete([]byte) error {
	return errors.New("cant delete an empty node")
}

func (Empty) Get([]byte, NodeResolverFn) ([]byte, error) {
	return nil, nil
}

func (Empty) ComputeCommitment() *Fr {
	return &FrZero
}

func (Empty) GetProofItems(keylist) (*ProofElements, byte, []byte) {
	panic("trying to produce a commitment for an empty subtree")
}

func (Empty) Serialize() ([]byte, error) {
	return nil, errors.New("can't encode empty node to RLP")
}

func (Empty) Copy() VerkleNode {
	return Empty(struct{}{})
}

func (Empty) toDot(string, string) string {
	return ""
}
