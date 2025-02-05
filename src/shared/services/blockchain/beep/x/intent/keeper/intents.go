package keeper

import (
	"context"
	"encoding/binary"

	"beep/x/intent/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetIntentsCount get the total number of intents
func (k Keeper) GetIntentsCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.IntentsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetIntentsCount set the total number of intents
func (k Keeper) SetIntentsCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.IntentsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendIntents appends a intents in the store with a new id and update the count
func (k Keeper) AppendIntents(
	ctx context.Context,
	intents types.Intents,
) uint64 {
	// Create the intents
	count := k.GetIntentsCount(ctx)

	// Set the ID of the appended value
	intents.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IntentsKey))
	appendedValue := k.cdc.MustMarshal(&intents)
	store.Set(GetIntentsIDBytes(intents.Id), appendedValue)

	// Update intents count
	k.SetIntentsCount(ctx, count+1)

	return count
}

// SetIntents set a specific intents in the store
func (k Keeper) SetIntents(ctx context.Context, intents types.Intents) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IntentsKey))
	b := k.cdc.MustMarshal(&intents)
	store.Set(GetIntentsIDBytes(intents.Id), b)
}

// GetIntents returns a intents from its id
func (k Keeper) GetIntents(ctx context.Context, id uint64) (val types.Intents, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IntentsKey))
	b := store.Get(GetIntentsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIntents removes a intents from the store
func (k Keeper) RemoveIntents(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IntentsKey))
	store.Delete(GetIntentsIDBytes(id))
}

// GetAllIntents returns all intents
func (k Keeper) GetAllIntents(ctx context.Context) (list []types.Intents) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IntentsKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Intents
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetIntentsIDBytes returns the byte representation of the ID
func GetIntentsIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.IntentsKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
