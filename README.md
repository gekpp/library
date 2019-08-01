# Books Library and Board Games

Library and board games is a projectinitiated by my sons who were eager to share their books and find mates for playing their boardgames. They decided to build a website for this purpose and run ad campaign in local heighbourhood.

This repo is the backend built on go with [goa](https://github.com/goadesign/goa) framework and code generator.

## Run and test

### 1. Run http-server

```bash
make run
```

After design is changed

```bash
make gen example build run
```

### 2. For generating token. This step could be ommited because not all requests require token

```bash
bin/books-cli auther signin -username librarian -password library
# if jq installed
# TOKEN=`bin/books-cli auther signin -username librarian -password library | jq -r ".JWT"`
 ```

### 3. Call the books service

```bash
bin/books-cli books reserve -bookid abc-der -token $TOKEN
```
