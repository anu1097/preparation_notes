## ACID

ACID is an acronym that stands for Atomicity, Consistency, Isolation, and Durability. These properties are fundamental principles in database systems that ensure reliable and robust transaction processing. Let's dive into each property in detail:

**Atomicity:**: 
Atomicity guarantees that a transaction is treated as a single, indivisible unit of work. It means that either all the changes made within a transaction are committed, or none of them are. If any part of a transaction fails, the entire transaction is rolled back, and the database remains unchanged. Atomicity ensures data integrity and prevents incomplete or inconsistent updates.

- Example: Performing a funds transfer from one bank account to another. Either the entire transaction of debiting one account and crediting the other account is completed, or none of it happens.

- [x] Debit Account A
- [x] Credit Account B

**Consistency:**: 
Consistency ensures that a transaction brings the database from one valid state to another. It defines a set of rules or constraints that the data must follow. If a transaction violates these rules, it is rolled back, and the database remains in its original state. Consistency ensures that data remains valid and reflects the integrity constraints defined by the database schema.

- Example: Enforcing a unique constraint on a username field in a user registration process. If a username is already taken, the transaction is rolled back, and the user is prompted to choose a different username.

- [x] Check if username is available
- [x] Register new user

**Isolation:**:
Isolation ensures that concurrent transactions do not interfere with each other. Each transaction should execute as if it is the only transaction running on the system. Isolation prevents issues like dirty reads, non-repeatable reads, and phantom reads. It ensures that the final outcome of a set of concurrent transactions is the same as if they were executed serially.

- Example: Simultaneous booking of airline tickets by multiple users. Each user's transaction should be isolated from others, ensuring that no conflicts or inconsistencies occur.

- [x] User 1: Reserve Seat A
- [x] User 2: Reserve Seat B

**Durability:**: 
Durability guarantees that once a transaction is committed, its changes are permanent and will survive any subsequent failures, such as power outages or system crashes. Durability is typically achieved by writing the transaction's changes to disk or other non-volatile storage. This property ensures that committed data is not lost and can be recovered in case of failures.

- Example: After successfully submitting an online order, ensuring that the order details are permanently stored and can be retrieved even if there is a system failure.

- [x] Store order details in database
- [x] Send confirmation email to customer

 
ACID properties provide a strong foundation for reliable and consistent database operations. However, it's important to note that enforcing strict ACID properties can impact performance. In some cases, relaxing certain properties (such as isolation) may be acceptable based on specific application requirements. This trade-off is often managed through different isolation levels and transaction management techniques.
It's worth mentioning that while ACID properties are commonly associated with traditional relational databases, other types of databases, such as NoSQL databases, may have different consistency and isolation models that relax some ACID properties to achieve scalability or performance benefits.