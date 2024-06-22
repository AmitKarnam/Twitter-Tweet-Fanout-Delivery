package users

type User string

var Users []User

var UserFollowers map[User][]User
