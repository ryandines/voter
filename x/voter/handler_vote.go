package voter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ryandines/voter/x/voter/types"
	"github.com/ryandines/voter/x/voter/keeper"
)

func handleMsgCreateVote(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateVote) (*sdk.Result, error) {
	k.CreateVote(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateVote(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateVote) (*sdk.Result, error) {
	var vote = types.Vote{
		Creator: msg.Creator,
		Id:      msg.Id,
    	PollID: msg.PollID,
    	Value: msg.Value,
	}

    if msg.Creator != k.GetVoteOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdateVote(ctx, vote)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteVote(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteVote) (*sdk.Result, error) {
    if !k.HasVote(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if msg.Creator != k.GetVoteOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeleteVote(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
