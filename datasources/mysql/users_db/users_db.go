package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const(
  userName = "root"
  password = "murali"
  host = "localhost:3306"
  schema = "users_db"
)
// userName := os.Getenv(MYSQL_USERS_USER)
// password := os.Getenv(MYSQL_USERS_PASSWORD)
// host := os.Getenv(MYSQL_USERS_HOST)
// schema := os.Getenv(MYSQL_USERS_SCHEMA)


// type Config struct {
//   USERNAME string `json:"username"`
//   PASSWORD string `json:"password"`
//   HOST string `json:"host"`
//   SCHEMA string `json:"schema"`
// }

var (
  Client *sql.DB
) 

func init(){
  // fptr := flag.String("fpath", "config.json", "config file path  to read from")
  // flag.Parse()
  // data, err := ioutil.ReadFile(*fptr)
  // if err != nil {
  //     fmt.Print(err)
  //     return
  // }
  // // json data
  // var obj Config
  // // unmarshall it
  // err = json.Unmarshal(data, &obj)
  // if err != nil {
  //    fmt.Println("error:", err)
  //    return 
  // }
  // // can access using struct now
  // fmt.Println(obj.USERNAME);
  // fmt.Println(obj.PASSWORD);
  // fmt.Println(obj.HOST);
  // fmt.Println(obj.SCHEMA);
  // userName := obj.USERNAME
  // password := obj.PASSWORD
  // host := obj.HOST
  // schema := obj.SCHEMA
  dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
   userName,
   password,
   host,
   schema,
  ) 
  fmt.Println(dataSourceName)
  var err error
  Client, err = sql.Open("mysql", dataSourceName)  
  if err != nil{
    panic(err)
    return
  }
  if err= Client.Ping(); err !=nil {
    panic(err)
    return 
  }
  fmt.Println("Successfully connected to database")
  log.Println("Database connected successfully ")
  
}