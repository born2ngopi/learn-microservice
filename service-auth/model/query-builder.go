package model

import (
  "fmt"
  "github.com/learn-microservice/service-auth/entities"
)

func Init(){
  return entities.Users
}

/*
* function for get all data user in table Users
* @return array struct
* @see package.entities
*/
func FetchAll() []entities.Users {

  //Connect database
  db, err := Connect();
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  //get all data from table users
  rows, err := db.Query("select id, username, email from users")
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  defer rows.Close()

  //slice form user entities
  var users []entities.Users

  for rows.Next(){
    var user = entities.Users{}
    var err = rows.Scan(user.Id, user.Username, user.Email)

    if err != nil {
      fmt.Println(err.Error())
      return
    }

    users = append(users, user)
  }

  return users
}

/*
* function for get data user by username
* @param string username
* @return struct
* @see package.entities
*/

func fetchByUsername(username string) entities.Users {
  //get connection database
  db, err := Connect();
  if err != nil {
      fmt.Println(err.Error());
      return
  }
  defer db.Close();

  // get data user by username
  row, err := db.Query("select id, username, email form users where username = ? limit=1",username)
  if err != nil {
    fmt.Println(err.Error());
    return
  }

  user := entities.Users{}
  err = row.Scan(user.id, user.Username, user.Email)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  return user
}
