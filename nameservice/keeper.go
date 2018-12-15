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

      allows code in this module to call functions from the bank module.

      The SDK uses an object capabilities approach to access sections of the application state.

      This allows us developers to employ a least authority approach, limiting the capabilities of a

      faulty or malicious module from affecting parts of state it doesnt need access to.

  - *codec.Codec => A pointer to the codec is used by Amino to encode and decode binary structs

  - sdk.StoreKey => gates access to a sdk.KVStore that persits the state of our application

Our module has 3 store keys:

  - namesStoreKey => The main store that stores value strings
                     that the name points to
                     (i.e. map[name]value).

  - ownersStoreKey => This store contains the current owner of any given name
                    (i.e. map[sdk_address]name).

  - priceStoreKey => This store contains the price that the current owner paid for a given name.
                    Anyone buying this name must spend more than the current owner
                    (i.e. map[name]price).



***Getters and Setters***

*/


//SetName - sets the value string that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
    store := ctx.KVStore(k.namesStoreKey)
    store.Set([]byte(name), []byte(value))
}


/* We first get the store object for the map[name]value using namesStoreKey from Keeper

NOTE: This function uses the sdk.Context. This object holds functions to access a number
      of important pieces of the state like:


      func (c Context) BlockHeight() int64

      func (c Context) BlockHeader() abci.Header

      func(c Context) chainID() string

      func (c Context) ConsensusParams() *abci.ConsensusParams

      func (c Context) KVStore(key StoreKey) KVStore


Next, insert the <name, value> pair into the store using:

    .Set([]byte, []byte method.

    As the store only takes []byte, we first cast the strings to []byte

    and then use them as parameters into the Set method.)


*/

func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
  store := ctx.KVStore(k.namesStoreKey)
  bz := store.Get([]byte(name))
  return string(bz)
}



/* again, like in SetName, we first access the store using StoreKey.

  Next, instead of using Set method on the store key, use the

  .Get([]byte) []byte method.

  As the parameter into the function, we pass the key, which is the name string casted to
  []byte.

  We get back the result in the form of []byte.

  Cast this to a string and return the result.


  ***getting and setting name owners***

  */


// HasOwner - returns wether or not the name already has an owner

  func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
    store := ctx.KVStore(k.ownersStoreKey)
    bz := store.Get([]byte(name))
    return bz != nil
  }


// GetOwner - get the current owner of a name

  func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
    store := ctx.KVStore(k.ownersStoreKey)
    bz := store.Get([]byte(name))
    return bz
  }

// SetOwner - sets the current owner of a name

  func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress {
    store := ctx.KVStore(k.ownersStoreKey)
    store.Set([]byte(name), owner)
  }


/*

Notes on code above:

    -Instead of accessing the data from namesStoreKey store, get it from ownersStoreKey store.

    -Because sdk.AccAddress is a type alias for []byte, it can natively be casted to it.

    -There is an extra function,
     HasOwner that is a convenience to be used in conditional statements



***Finally, add a getter and a setter for the price of a name***

*/


func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
  if !k.HasOwner(ctx, name) {
    return sdk.Coins{sdk.NewInt64Coin("educoin", 1)}
  }
  store := ctx.KVStore(k.pricesStoreKey)
  bz := store.Get([]byte(name))
  var price sdk.Coins
  k.cdc.MustUnmarshalBinaryBare(bz, &price)
  return price
}

//SetPrice - sets current price of name

func (k keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
  store := ctx.KVStore(k.pricesStoreKey)
  store.Set([]byte(name), k.cdc.MustUnmarshalBinaryBare(price))
}


/*

Notes on code above:

    - sdk.Coins has no own bytes encoding => price needs to be marshalled and
      unmarshalled using Amino to be inserted or removed from store.

    -When getting price for a name that has no owner (=> no price) return 1steak as price.


    ***constructor function for keeper***
*/


//NewKeeper creates new instances of the nameservice Keeper

    function NewKeeper(coinKeeper bank.Keeper, namesStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, priceStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
      return Keeper {
        coinKeeper:     coinKeeper,
        namesStoreKey:  namesStoreKey,
        ownersStoreKey: ownersStoreKey,
        pricesStoreKey: priceStoreKey,
        cdc:            cdc,
      }

    }


// Next => describe how users interact with store using Msgs and Handlers




*/




