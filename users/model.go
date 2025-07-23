package user 

import "time" 

type User struct {
   ID string 
   FirstName string 
   LastName string 
   DateOfBirth time.Time 
   Country string 
   Language string 
   Email string 
   CreatedAt time.Time
}
