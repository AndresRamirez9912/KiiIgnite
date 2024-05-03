package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"kiichain/x/kiichain/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdatePostResponse{}, nil
}
