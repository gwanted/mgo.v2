The MongoDB driver for Go
-------------------------

Add a mock model for mongodb test

The origin code is from gopkg.in/mgo.v2

## Usage
get this package and replace the gopkg.in/mgo.v2

## example:
   	mgo.Mock = true // open mock model

	mgo.SetMckC("Query-All",0) // in the next function ,the first query.All() will throw an error out

	// TODO:add the function need to be tested here

	mgo.ClearMock() // close mock query which is seted in mgo.SetMckC()
