# NewAPITestground
Just an area to muck around with API tech


The structure of this API project is as follows.

With a main.go as the starting point, handlers will be where all the functions for clients/services are called. Clients will be specialized handlers for SQL or for Database queries.

Services will be handling any kind of logic that needs to be done, while handlers is for the directing of these requests.

A 'router.go' will be generated in handlers as well to call the end points.


Currently, just building a little something to mess with any new Go based API tech I find interesting.