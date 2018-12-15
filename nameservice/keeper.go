/* Main core of the Cosmos SDK module is a piece called the KEEPER.
   The keeper handles interactions with the store, has references to other keepers
   for cross-module interations, and contains most of the core functionality of a module.

   In Cosmos SDK applications, the convention is that modules livce in the ./x/ folder.

   lets define our nameservice.keeper

*/


package nameservice

   import (
   "github.com/cosmos/cosmos-sdk/codec"
   "github.com/cosmos/cosmos-sdk/x/bank"

   sdk "github.com/cosmosc/cosmos-sdk/types"

)

    /*Keeper maintains the link to data storage and exposes getter/setter methods
    for various parts of the machine*/
    type Keeper struct {
      coinKeeper bank.Keeper

      //Unexposed keys to access name, owners and prices store from sdk.Context
      namesStoreKey  sdk.StoreKey
      ownersStoreKey sdk.StoreKey
      pricesStoreKey sdk.StoreKey

      cdc *codec.Codec //The wire codec for binary encoding/decoding.

    }


/* Notes for above code:

  - We used 3 different cosmos-sdk packages:
      -the codec, provides tools to work with the Cosmos encoding format, Amino
      -the bank module, controls accounts and coin transfers
      -types, containing commonly used types throughout the SDK

  - The Keeper struct. In this keeper, we have the following key pieces:
      -bank.Keeper - a reference from the bank module.

      Including it allows code in this module to call functions from the bank module.

      The SDK uses an object capabilities approach to access sections of the application state.

      This allows us developers to employ a least authority approach, limiting the capabilities of a

      faulty or malicious module from affecting parts of state it doesnt need access to.




















