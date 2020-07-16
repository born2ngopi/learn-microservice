package entities;

type Users struct{
    Id        int     `form :"id" json:"id"`
    Username  string  `form:"username" json:"username"`
    Email     string  `form:"email" json:"email"`
}

type UserDetais struct {
  UserId    int     `json:"user_id"`
  FirstName string  `form:"firstname" json:"firstname"`
  LastName  string  `form:"lastname" json:"lastname"`
  Gender    string  `form:"gender" json:"gender"`
}
