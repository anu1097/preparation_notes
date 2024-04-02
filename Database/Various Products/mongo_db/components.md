1. **MongoDB Cluster:** A MongoDB cluster consists of multiple servers or nodes, each running an instance of the MongoDB database. The cluster can be set up as a standalone deployment, replica set, or sharded cluster.

2. **Sharding Architecture:** In a sharded cluster, data is partitioned and distributed across multiple shards. Each shard contains a subset of the data. The shard key, which determines the data distribution, is specified during sharding. MongoDB uses a hashed shard key or a range-based shard key for efficient data distribution.

3. **Shard Servers:** Shard servers are responsible for storing and managing the data for a specific shard. Each shard server is a separate instance of MongoDB. It manages a subset of the data and processes queries and operations related to that subset.

4. **Config Servers:** Config servers store the metadata and configuration information for the sharded cluster. They keep track of the mapping between the data chunks and the shards where they reside. Config servers enable the cluster to route queries to the appropriate shard servers.

5. **Query Routers:** Query routers, also known as mongos instances, act as the interface between client applications and the sharded cluster. They receive incoming queries from clients and route them to the appropriate shard servers based on the sharding configuration and metadata stored in the config servers. Query routers provide a unified view of the entire cluster to the client applications.

6. **Data Distribution and Balancing:** MongoDB's balancer component runs on the query routers and automatically redistributes the data across the shards. It ensures even data distribution and balances the data load across the cluster. The balancer also moves data between shards when the shard key range changes or when new shards are added.

7. **Replica Set Architecture:** In a replica set deployment, MongoDB uses a primary-secondary replication model. The primary node handles all write operations and replicates data to the secondary nodes. The secondary nodes act as hot backups and can serve read operations in case of a primary node failure.

8. **Oplog:** The oplog (operation log) is a special collection in MongoDB that records all write operations performed on the primary node of a replica set. Secondary nodes read the oplog and apply the recorded operations to stay in sync with the primary. The oplog ensures data consistency across the replica set.

9. **Data Storage:** MongoDB stores data in a binary format called BSON (Binary JSON). BSON is a lightweight, schema-less, and efficient format that supports various data types. Data is organized into collections, which are analogous to tables in a relational database. Each collection can have multiple documents, which are JSON-like data structures.

10. **Indexes:** MongoDB supports different types of indexes, such as B-tree, hashed, text, and geospatial indexes. Indexes improve query performance by allowing for faster data retrieval based on specified fields. MongoDB automatically uses indexes to optimize query execution.
