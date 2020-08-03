package database

import (
   "../../domain"
   "fmt"
)

type UserRepository struct {
  SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
  result, err := repo.Execute(
    "INSERT INTO users (first_name, last_name) VALUES (?, ?)", u.FirstName, u.LastName,
  )
  if err != nil {
    return
  }

  id64, err := result.LastInsertId()
  if err != nil {
    return
  }

  id = int(id64)
  return
}


func (repo *UserRepository) FindById(identifier int) (user domain.User, err error){
  row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
  defer row.Close()
  var id int
  var firstName string
  var lastName string
  row.Next()
  if err = row.Scan(&id, &firstName, &lastName); err != nil {
    return
  }

  user.Id = id
  user.FirstName = firstName
  user.LastName = lastName
  //user.Build()
  return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
    rows, err := repo.Query("SELECT * FROM users")
    defer rows.Close()
    if err != nil {
        fmt.Print("hello")
         return
    }

   var id int
   var firstName string
   var lastName string


   for rows.Next() {
        if err := rows.Scan(&id, &firstName, &lastName); err != nil {
            continue
        }
   }
   user := domain.User{
        Id:        id,
        FirstName: firstName,
        LastName:  lastName,
   }
   users = append(users, user)

   return
}

