/* Register types with Amino to be encoded / decoded.

   Any interface we create and any struct that implements and interface needs
   to be declared in the RegisterCodec function.

   In this module the Msg implementations (SetName and BuyName) need to be registered,
   but our Whois query return type does not
*/
package nameservice

import (

    "github.com/cosmos/cosmos-sdk/codec"

    )

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
    cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
    cdc.RegisterConcrete(MsgBuyName{}, "nameservice/BuyName", nil)
}

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
  var cdc = codec.New()
  auth.RegisterCodec(cdc)
  bank.RegisterCodec(cdc)
  nameservice.RegisterCodec(cdc)
  stake.RegisterCodec(cdc)
  sdk.RegisterCodec(cdc)
  codec.RegisterCrypto(cdc)
  return cdc
}
