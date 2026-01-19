1. in server.go start server by fiber {app:=fiber.New()}
2. in handlers make user handlers
give user handler access to the user service through a struct
User handler {
    svc service.UserService
   
   }

3.make struct of request handler struct {app *fiber.App} and pass it to the routes of user handlers so that there they can app.POST/GET/PUT/UPDATE/DELETE