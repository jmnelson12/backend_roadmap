# Relational Databases

Going to focus on **PostgreSQL** because it is what I use at work, and I want a deeper understanding of it.

## What is PostgreSQL

PostgreSQL is an open source object-relational database system that uses and extends the SQL language combined with many features that store and scale data workloads.

- Runs on all major operating systems
- [ACID](https://en.wikipedia.org/wiki/ACID)-compliant

## Why use PostgreSQL

- you can define your own data types
- build out custom functions
- write code from different programming languages without recompiling your database
- highly extensible
- It is an excellent DBMS, which is scalable, performs well, allows replication, supports ACID and a big subset of the SQL standard (in several cases, it is a superset).
- Is much better at data types than the other DBMSs, with a more rich semantics, with geo-spatial types, complex numbers, etc.
- It supports several methods of indexation, including B-trees, Genetic Algorithms based indexes, and GIN indexes that accelerates full-text searches.

## Tutorial

### How To Install PostgreSQL on Ubuntu 20.04 [Quickstart]

https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart

### Command-Line Commands

- `createdb <dbname>`
  - `dbname` defaults to user account name
- `dropdb <dbname>`
- `psql <dbname>`
  - `dbname` defaults to user account name
  - interactive terminal which allows you to interactively enter, edit, and execute SQL commands.

### `psql` Commands

- `\q` - exit the interactive terminal
- `\h` - get help on the syntax of various PostgreSQL SQL commands
- `\?` - get list of available commands
- `\! clear` - clear screen
- `COPY weather FROM '/home/user/weather.txt';`
  - Load large amounts of data from flat-text files.
    This is usually faster because the COPY command is optimized for this application while allowing less flexibility than `INSERT`.

### General Notes

- `mydb=#` means superuser vs `mydb=>`
- Can use multi-lines on the command line. The command isn't run until `;` is used
> While `SELECT *` is useful for off-the-cuff queries, it is widely considered bad style in production code, since adding a column to the table would change the results.

> If no matching row is found we want some “empty values” to be substituted for the cities table's columns. This kind of query is called an `outer join`.

> `left outer join` because the table mentioned on the left of the join operator will have each of its rows in the output at least once, whereas the table on the right will only have those rows output that match some row of the left table
```sql
-- the first way is more commonly used
-- this
SELECT w.city
  , w.temp_lo
  , w.temp_hi
  , w.prcp
  , w.date
  , c.location
FROM weather w, cities c
WHERE c.name = w.city;

-- is the same as
SELECT *
FROM weather
INNER JOIN cities
  ON (weather.city = cities.name);

```

## References

- https://www.postgresql.org/about/
- https://www.postgresql.org/docs/current/tutorial.html

# Left off at https://www.postgresql.org/docs/current/tutorial-agg.html