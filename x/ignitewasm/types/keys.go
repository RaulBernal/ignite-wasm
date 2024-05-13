package types

const (
	// ModuleName defines the module name
	ModuleName = "ignitewasm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ignitewasm"
)

var (
	ParamsKey = []byte("p_ignitewasm")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
