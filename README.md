# Bookstore

Create  
Read  
Update  
Delete  

### Hexagonal Architecture
![Image](https://miro.medium.com/v2/resize:fit:500/format:webp/1*LF3qzk0dgk9kfnplYYKv4Q.png)

```mermaid
sequenceDiagram
  Input Port->>Logic: CreateRequest
  Logic->>Output Port: Create(Title, Author, Section)
  Output Port->>Logic: Id
  Logic->>Input Port: CreateResponse(Id)
  ```
---
```mermaid
sequenceDiagram
  Input Port->>Logic: CreateRequest
  Logic->>Output Port: Create(Title, Author, Section)
  Output Port->>Logic: Error
  Logic->>Input Port: ErrorResponse(error)
  ```

```
type Section = Fiction | NonFiction

type Book =
  { id: int
    title: string
    author: string
    section: Section }

type CreateRequest =
  { title: string
    author: string
    section: Section }

type CreateResponse =
  { id: int }

type ErrorResponse = 
  { error: string }
```

### Things to try

- Complete the missing requests
- Add validation, maybe a title has to be less than 50 chars
- Create a CLI or HTTP adaptor for the input port, add `main.go` and run the app!
- Create a real database implementation, this could just be writing to the file system