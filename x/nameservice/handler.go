package nameservice

import (
"fmt"

sdk "github.com/cosmos/cosmos-sdk/types"
)



/*

NOTE: The naming convention for handler names in the SDK is handleMsg{ .Action }

NewHandler is essentially a sub-router, directing messages coming into this module
to the proper handler

*/

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.handler {
  return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
    switch msg := msg.(type) {

    case MsgSetName:
      return handleMsgSetName(ctx, keeper, msg)

    case MsgBuyName:
      return handleMsgBuyName(ctx, keeper, msg)

    default:
      errorMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
      return sdk. ErrUnknownRequest(errMsg).Result()
    }
  }
}

// Handle MsgSetName
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
  if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.NameID)) { // checks if msg sender is same as current owner
          return sdk.ErrUnauthorized("Incorrect Owner").Result() //If not, throw an error
        }
     keeper.SetName(ctx, msg.NameID, msg.value) //If so, set the name to the value specified in msg.
     return sdk.Result{}
   }

// Handle MsgBuyName
   func handleMsgBuyName(ctx sdk.Context, keeper Keeper, MsgBuyName) sdk.Result {
  if keeper.GetPrice(ctx, msg.NameID).IsAllgt(msg.Bid) // checks if bid price is greater than price paid by current owner
          return sdk.ErrInsufficientCoins("Bid not high enough").Result() //if not, throw an error
        }
        if keeper.HasOwner(ctx, msg.NameID) {
          _, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.NameID), msg.Bid)
          if err != nil {
            return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
          }
        } else {
              _,_, err := keeper.coinKeeper.SubstractCoins(ctx, msg.Buyer, msg.Bid) //if so, deduct bid amount from the sender
              if err != nil {
                return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
              }
            }
            keeper.SetOwner(ctx, msg.NameID, msg.Buyer)
            keeper.SetPrice(ctx, msg.NameID, msg.Bid)
            return sdk.Result()
          }

