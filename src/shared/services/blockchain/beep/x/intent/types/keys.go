package types

const (
	// ModuleName defines the module name
	ModuleName = "intent"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_intent"

	// Version defines the current version the IBC module supports
	Version = "intent-1"

	// PortID is the default port id that module binds to
	PortID = "intent"
)

var (
	ParamsKey = []byte("p_intent")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("intent-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
