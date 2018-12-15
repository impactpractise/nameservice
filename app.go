//in app.go we define what the application does once it receives a transaction

//To receive transactions in correct order we use the Tendermint consensus engine

package app

import (
  "github.com/tendermint/tendermint/libs/log"
    "github.com/cosmos/cosmos-sdk/x/auth")

    bam "github.com/cosmos/cosmos-sdk/baseapp"
    dbm "hithub.com/tendermint/tendermint/libs/db"

/*links to godocs for each module and package imported:

- log => Tendermint's logger
- auth => The auth module for the Cosmos SDK
- dbm => Code for working with the Tendermind Database
- baseapp => See below

Included Tendermint packages are used to pass transaction from the network to the
application through an interface called ABCI.

The achitecture of the blockchain node we are building looks like the following:

+-----------------------+
|                       |

|      application      |

|                       |
+-------+------+--------+

        ^     |
        |
              | ABCI
        |     v

+-----------------------+
|                       |

|      Tendermint       |

|                       |
+-------+------+--------+

The Cosmos SDK provides a boilerplate implementation of the ABCI interface the form of baseapp

Here is what baseapp does:

- Decode transactions received from the Tendermint consensus engine.

- Extract messages from transactions and do basic sanity checks.

- Route the message to the appropriate module so that in can be processed. Note that baseapp
  has no knowledge of the specific modules we want to use. It is our job to declare such modules
  in app.go, as we will see soon.
  baseapp implements the core routing logic that can be applied to any module.

- Commit if the ABCI message is DeliverTx (CheckTx changes are not persistent).

- Help setup BEGINBLOCK and ENDBLOCK, two messages that enables us to define logic executed
  at the beginning and end of each block.
  In Practice, each module implements its own Beginblock and Endblock sub-logic,
  and the role of the app is to aggregate everything together

- Help Initialise our state

- Help set up queries.


Now we are going to create the skeleton of our a custom type nameserviceApp application.

This type will embed baseapp (embedding in Go is similar to inheritance in other languages),
meaning it will have access to all of baseapps methods.

*/


const (
    appName = "nameservice"
)

type nameserviceApp struct {
  *bam.BaseApp
}

// lets add a simple constructor for our application

func NewnameserviceApp(logger log.logger, db dbm.DB) *nameserviceApp{

  //First we define the top level codec that will be shared by the different modules
  cdc := MakeCodec()

  //BaseApp handles interations with Tendermint through the ABCI protocol
  bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

  var app = &nameserviceApp{
    BaseApp: bApp,
    cdc:      cdc,
  }

  return app
}

/* Cool, we have a basic skeleton, yet it still lacks functionality.

- baseapp has no knowledge of the routes or user interactions we want to use in our application.
  Main roles of app:
  define routes
  define initial state

  Both require that we add modules to our application.

  For our nameservice app, we need three modules:
  -Auth
  -Bank
  -nameservice

  The first are provided by the SDK. Lets build the nameservice module
  and define the bulk of our state machine.
