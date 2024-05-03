package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"kiichain/x/kiichain/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreatePostResponse{}, nil
}