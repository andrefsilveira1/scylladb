# ScyllaDB

This is part of my journey into the land of advanced Golang building HTTP Servers with ScyllaDB.

# Reporting System
This simple system store some reports by ReportID, UserID, Status and Title. 

# Running

## First, create a scyllaDB container:
`docker run --name scylla -d scylladb/scylla`

In this case, a container called "scylla" was created.

## Second, runs cqlsh in scylla container:
`docker exec -it scylla cqlsh`

Check out the host and the port connected by cqlsh


## Second, run the migration:
`SCYLLA_HOST=YOUR_HOST_HERE SCYLLA_KEYSPACE=reporting SCYLLA_MIGRATIONS=migrations go run exec/migration/main.go`

Notice that you will have to change **YOUR_HOST_HERE** by the host and the server returned by cqlsh command.

## Have fun!

Now, just have fun with **INSERT** and **SELECT** querys. Remember to run **USE reporting;**, this will able you to use the keyspace defined.



# References:

Using tutorial of Amin-mir: https://github.com/amin-mir/reporting
