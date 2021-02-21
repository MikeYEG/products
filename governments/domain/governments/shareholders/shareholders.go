package shareholders

import "github.com/deepvalue-network/software/libs/hash"

type shareHolders struct {
	hash hash.Hash
	list []ShareHolder
}

func createShareHolders(
	hash hash.Hash,
	list []ShareHolder,
) ShareHolders {
	out := shareHolders{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *shareHolders) Hash() hash.Hash {
	return obj.hash
}

// All returns the shareholders
func (obj *shareHolders) All() []ShareHolder {
	return obj.list
}

// Has returns true if the given shareHolder is in the list, false otherwise
func (obj *shareHolders) Has(shareHolder ShareHolder) bool {
	hsh := shareHolder.Hash()
	for _, oneShareHolder := range obj.list {
		if oneShareHolder.Hash().Compare(hsh) {
			return true
		}
	}

	return false
}

// Same returns true if the given pubkey hashes are validated by 1 shareholder
func (obj *shareHolders) Same(pubKeyHashes []hash.Hash) bool {
	for _, oneHolder := range obj.list {
		if oneHolder.Same(pubKeyHashes) {
			return true
		}
	}

	return false
}
