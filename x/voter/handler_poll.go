package voter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ryandines/voter/x/voter/types"
	"github.com/ryandines/voter/x/voter/keeper"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePoll) (*sdk.Result, error) {
	k.CreatePoll(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdatePoll(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdatePoll) (*sdk.Result, error) {
	var poll = types.Poll{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Title: msg.Title,
    	Options: msg.Options,
	}

    if msg.Creator != k.GetPollOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdatePoll(ctx, poll)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeletePoll(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeletePoll) (*sdk.Result, error) {
    if !k.HasPoll(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if msg.Creator != k.GetPollOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeletePoll(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
