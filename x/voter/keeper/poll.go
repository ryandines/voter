package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ryandines/voter/x/voter/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetPollCount get the total number of poll
func (k Keeper) GetPollCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollCountKey))
	byteKey := types.KeyPrefix(types.PollCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetPollCount set the total number of poll
func (k Keeper) SetPollCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollCountKey))
	byteKey := types.KeyPrefix(types.PollCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreatePoll(ctx sdk.Context, msg types.MsgCreatePoll) {
	// Create the poll
    count := k.GetPollCount(ctx)
    var poll = types.Poll{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Title: msg.Title,
        Options: msg.Options,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
    key := types.KeyPrefix(types.PollKey + poll.Id)
    value := k.cdc.MustMarshalBinaryBare(&poll)
    store.Set(key, value)

    // Update poll count
    k.SetPollCount(ctx, count+1)
}

func (k Keeper) UpdatePoll(ctx sdk.Context, poll types.Poll) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	b := k.cdc.MustMarshalBinaryBare(&poll)
	store.Set(types.KeyPrefix(types.PollKey + poll.Id), b)
}

func (k Keeper) GetPoll(ctx sdk.Context, key string) types.Poll {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	var poll types.Poll
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PollKey + key)), &poll)
	return poll
}

func (k Keeper) HasPoll(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	return store.Has(types.KeyPrefix(types.PollKey + id))
}

func (k Keeper) GetPollOwner(ctx sdk.Context, key string) string {
    return k.GetPoll(ctx, key).Creator
}

// DeletePoll deletes a poll
func (k Keeper) DeletePoll(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	store.Delete(types.KeyPrefix(types.PollKey + key))
}

func (k Keeper) GetAllPoll(ctx sdk.Context) (msgs []types.Poll) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PollKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Poll
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
