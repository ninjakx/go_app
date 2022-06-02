# GO_APP (Without GORM)golang

DB Schema: https://dbdiagram.io/d/628b853bf040f104c17c271a

![](https://github.com/KSahu1705/GO_APP/blob/master/db.png)

- User Payment have not implemented it but logic will be same as that of address one.
- To keep the work flow simple didn't implement hashing on password and credential.
- Have provided the ID manually for now.
- Done with API testing using Postman.

### TO DO:
- UNIT Testing Using MockGen.

### API
```Golang
// Routing for handling the projects
router.GET("/users", a.GetAllUser)
router.GET("/users/:id", a.GetUser)
router.GET("/users/:id/address", a.GetUserAddress)
router.POST("/users", a.CreateUser)
router.POST("/users/:id/add_address", a.CreateUserAddress)
router.PUT("/users/:id/update_user", a.UpdateUser)
router.PUT("/users/:id/update_address/:addr_id", a.UpdateUserAddress)
router.PUT("/users/:id/disable", a.DisableUser)
router.PUT("/users/:id/enable", a.EnableUser)
router.DELETE("/users/:id", a.DeleteUser)
router.DELETE("/users/:id/del/:addr_id", a.DeleteUserAddress)
```

### RUN:

`go run main.go`