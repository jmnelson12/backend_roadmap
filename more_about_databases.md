# More About Databases

## ORMs

### What is an ORM

An ORM is a layer between the server and the database.
The server talks with the ORM and the ORM talks to the database.
The ORM creates objects, that map to the relational data.
It handles your queries, so you don't have to write native SQL, you can query the database with your application language.

### Popular Nodejs ORMs

- sequelize
- typeORM
- mongoose

## ACID

**Atomicity**

Guarantee that either all of the transactions succeeds, or none of it does.

**Consistency**

All data will be valid according to all defined rules, including any constraints, cascades, and triggers that have been applied on the database.

**Isolation**

No transaction will be affected by any other transaction.

**Durability**

Once a transaction is committed, it will remain in the system – even if there’s a system crash immediately following the transaction.

## Transactions

A transaction bundles multiple steps into a single, all-or-nothing operation. The intermediate states between the steps are not visible to other concurrent transactions, and if some failure occurs that prevents the transaction from completing, then none of the steps affect the database at all.

They are set up by surrounding the SQL commands of the transaction with `BEGIN` and `COMMIT` commands.

The command `ROLLBACK` cancels all updates up to that point.

Have savepoints which allow you to selectively discard parts of the transaction, while committing the rest.

## N+1 Problem

Looked at these two references

- https://www.youtube.com/watch?v=Xr3hZdIwuSw
- https://thecodingmachine.io/solving-n-plus-1-problem-in-orms
- https://medium.com/slite/avoiding-n-1-requests-in-graphql-including-within-subscriptions-f9d7867a257d

## Normalization

## Index and how they work

## Data Replication

## Sharding Strategies

## CAP Theorem

## References

- https://database.guide/what-is-acid-in-databases/
- https://www.postgresql.org/docs/current/tutorial-transactions.html
- https://www.youtube.com/watch?v=Xr3hZdIwuSw
- https://thecodingmachine.io/solving-n-plus-1-problem-in-orms
