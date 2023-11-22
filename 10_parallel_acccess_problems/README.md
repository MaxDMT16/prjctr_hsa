# Parallel access problems
- lost update 
- dirty read
- non-repeatable read
- phantom read

## Percona

| Isolation level  | Lost update | Dirty read | Non-repeatable read | Phantom read |
| ---------------- | ----------- | ---------- | ------------------- | ------------ |
| Read uncommitted | No          | Yes        | Yes                 | Yes          |
| Read committed   | No          | No         | Yes                 | Yes          |
| Repeatable read  | No          | No         | No                  | No           |
| Serializable     | No          | No         | No                  | No           |

## PostgreSQL

| Isolation level  | Lost update | Dirty read | Non-repeatable read  | Phantom read  |
| ---------------- | ----------- | ---------- | -------------------  | ------------  |
| Read committed   | No          | No         | No                   | Yes           |
| Read uncommitted | No          | No         | No                   | Yes           |
| Repeatable read  | No          | No         | No                   | No            |
| Serializable     | No          | No         | No                   | No            |
