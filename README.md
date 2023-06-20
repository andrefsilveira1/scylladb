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


## Second, define your environment variables in .env file:
- SCYLLA_HOST=localhost (notice that you will have to change for your host)
- SCYLLA_KEYSPACE=reporting
- SCYLLA_MIGRATIONS=migrations


### Run:
In the folder src, run:
`go run exec/migration/main.go`

## Have fun!

Now, just have fun with **INSERT** and **SELECT** querys. Remember to run **USE reporting;**, this will able you to use the keyspace defined.


# Next steps:

- Refactor the code
- Create a .env file and get env by functions
- Implement efficient ways for querys
- Learn and use Network Topology Strategy
- Apply differents replication factor
- Make this repository a ease setup for ScyllaDB connection with docker image of scylla



# References:

Using tutorial of Amin-mir: https://github.com/amin-mir/reporting
