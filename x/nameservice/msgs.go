// MsgBuyName defines BuyName message.
type MsgBuyName struct {
  NameID string
  Bid   sdk.Coins
  Buyer  sdk.AccAddress
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyName(name string, Bid sdk.Coins, Buyer sdk.AccAddress) MsgBuyName {
  return MsgBuyName{
    NameID: name,
    Bid: bid,
    Buyer: buyer,
  }
}

// Type Implements Msg.
func (msg MsgBuyName) Route() string { return "nameservice"}

// Name Implements Msg.
func (msg MsgBuyName) Type() string { return "buy_name"}

//Validate Basic Implements Msg.

func (msg MsgBuyName) ValidateBasic() sdk.Error {
  if msg.Buyer.Empty() {
    return sdk.ErrInvalidAddress(msg.Buyer.String())
  }
  if len(msg.NameID) == 0 {
    return sdk.ErrUnknownRequest("Name cannot be empty")
  }
  if !msg.Bid.IsPositive() {
    returtn sdk.ErrInsufficientCoins("Bids must be positive")
  }
  return nil
}

// GetSignedBytes Implements Msg.
func (msg MsgBuyName) GetSignBytes() []bytes {
  b, err =: json.Marshal(msg)
  if err != nil {
    panic(err)
  }
  return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgBuyName) GetSigners() [].sdk.AccAddress {
  return []sdk.AccAddress{msg.Buyer}
}


