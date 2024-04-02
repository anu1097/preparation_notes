## Cardinality:
In database design, cardinality refers to the relationship between two entities (tables) in terms of the number of occurrences or records they can have in relation to each other. It describes how instances of one entity are associated with instances of another entity.

Cardinality can be classified into three types:

**One-to-One (1:1) Cardinality**: In a one-to-one relationship, each record in one entity is associated with exactly one record in the related entity, and vice versa. For example, in a database of employees and their social security numbers, each employee has a unique social security number.

**One-to-Many (1:N) Cardinality**: In a one-to-many relationship, each record in one entity can be associated with multiple records in the related entity, but each record in the related entity is associated with only one record in the first entity. For example, in a database of customers and their orders, each customer can have multiple orders, but each order is associated with only one customer.

**Many-to-Many (N:M) Cardinality**: In a many-to-many relationship, each record in one entity can be associated with multiple records in the related entity, and vice versa. This type of relationship requires an intermediate table, often called a junction or join table, to represent the associations. For example, in a database of students and courses, each student can enroll in multiple courses, and each course can have multiple students.

Understanding cardinality is crucial for database modeling, as it helps define the appropriate relationships and constraints between entities, ensuring data integrity and efficient data retrieval.