# MongoManager
Database connector with MongoDB without dependencies in your code

### Install

```shell
go get github.com/lbernardo/mongomanager
```

### Use

```go
package main

import (
	"github.com/lbernardo/mongomanager"
	"os"
)

func main() {
    conn, err := mongomanager.NewDatabaseManager(os.Getenv("DATABASE_URI"))
    if err != nil {
        panic(err)
    }
    defer conn.Close() // Disconnect
    conn.SetDatabase("mydatabase") // Select database
}
```

### Functions

`func (d *DatabaseManager) Close()`
```go
// Close connection
package main

import (
	"github.com/lbernardo/mongomanager"
	"os"
)

func main() {
    conn, err := mongomanager.NewDatabaseManager(os.Getenv("DATABASE_URI"))
    if err != nil {
        panic(err)
    }
    defer conn.Close() // Disconnect
}
```

`func (d *DatabaseManager) GetItemById(table, id string, result interface{})`
```go
/// Get only Item By ID
package main

import (
	"github.com/lbernardo/mongomanager"
	"os"
)

type User struct {
    Name string 
}

func main() {
    var user User
    conn, err := mongomanager.NewDatabaseManager(os.Getenv("DATABASE_URI"))
    if err != nil {
        panic(err)
    }
    defer conn.Close() // Disconnect
    conn.SetDatabase("mydatabase") // Select database
    conn.GetItemById("users", "fe4c5295f954ecc939a2a6900bbaaab1", &user)
}
```
`func (d *DatabaseManager) GetManyItems(table string, filter interface{}, result interface{})`
```go
/// Get many items
package main

import (
	"github.com/lbernardo/mongomanager"
	"os"
)

type User struct {
    Name string 
}

func main() {
    var users []User
    conn, err := mongomanager.NewDatabaseManager(os.Getenv("DATABASE_URI"))
    if err != nil {
        panic(err)
    }
    defer conn.Close() // Disconnect
    conn.SetDatabase("mydatabase") // Select database
    conn.GetManyItems("users", User{Name: "Lucas"}, &users)
}
```
`func (d *DatabaseManager) GetOnlyItem(table string, filter interface{}, result interface{})`
```go
/// Get only Item
package main

import (
	"github.com/lbernardo/mongomanager"
	"os"
)

type User struct {
    Name string 
}

func main() {
    var user User
    conn, err := mongomanager.NewDatabaseManager(os.Getenv("DATABASE_URI"))
    if err != nil {
        panic(err)
    }
    defer conn.Close() // Disconnect
    conn.SetDatabase("mydatabase") // Select database
    conn.GetOnlyItem("users", User{Name: "Lucas"}, &user)
}
```
`func (d *DatabaseManager) InsertItem(table string, item interface{}) (string, error)`

```go
/// Insert Item
package main

import (
	"github.com/lbernardo/mongomanager"
	"log"
	"os"
)

type User struct {
	Name string
}

func main() {
	conn, err := mongomanager.NewDatabaseManager(os.Getenv("DATABASE_URI"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()             // Disconnect
	conn.SetDatabase("mydatabase") // Select database
	id := conn.InsertItem("users", User{Name: "Lucas"})
	log.Println("Id", id)
}
```
`func (d *DatabaseManager) SetDatabase(name string)`
```go
// Select database
package main

import (
	"github.com/lbernardo/mongomanager"
	"os"
)

func main() {
    conn, err := mongomanager.NewDatabaseManager(os.Getenv("DATABASE_URI"))
    if err != nil {
        panic(err)
    }
    defer conn.Close() // Disconnect
    conn.SetDatabase("mydatabase") // Select database
}
```

`func NewDatabaseManager(databaseURI string) (MongoManager, error)`
```go
// Create new connection
package main

import (
	"github.com/lbernardo/mongomanager"
	"os"
)

func main() {
    conn, err := mongomanager.NewDatabaseManager(os.Getenv("MONGO_BASEURI"))
    if err != nil {
        panic(err)
    }
    defer conn.Close() // Disconnect
    conn.SetDatabase("mydatabase") // Select database
}
```