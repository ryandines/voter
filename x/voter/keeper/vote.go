package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ryandines/voter/x/voter/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetVoteCount get the total number of vote
func (k Keeper) GetVoteCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteCountKey))
	byteKey := types.KeyPrefix(types.VoteCountKey)
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

// SetVoteCount set the total number of vote
func (k Keeper) SetVoteCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteCountKey))
	byteKey := types.KeyPrefix(types.VoteCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateVote(ctx sdk.Context, msg types.MsgCreateVote) {
	// Create the vote
    count := k.GetVoteCount(ctx)
    var vote = types.Vote{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        PollID: msg.PollID,
        Value: msg.Value,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
    key := types.KeyPrefix(types.VoteKey + vote.Id)
    value := k.cdc.MustMarshalBinaryBare(&vote)
    store.Set(key, value)

    // Update vote count
    k.SetVoteCount(ctx, count+1)
}

func (k Keeper) UpdateVote(ctx sdk.Context, vote types.Vote) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	b := k.cdc.MustMarshalBinaryBare(&vote)
	store.Set(types.KeyPrefix(types.VoteKey + vote.Id), b)
}

func (k Keeper) GetVote(ctx sdk.Context, key string) types.Vote {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	var vote types.Vote
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.VoteKey + key)), &vote)
	return vote
}

func (k Keeper) HasVote(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	return store.Has(types.KeyPrefix(types.VoteKey + id))
}

func (k Keeper) GetVoteOwner(ctx sdk.Context, key string) string {
    return k.GetVote(ctx, key).Creator
}

// DeleteVote deletes a vote
func (k Keeper) DeleteVote(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	store.Delete(types.KeyPrefix(types.VoteKey + key))
}

func (k Keeper) GetAllVote(ctx sdk.Context) (msgs []types.Vote) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.VoteKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Vote
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
