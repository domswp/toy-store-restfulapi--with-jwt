## Toy Store Advance with Json Web Token (Work In Progress)
A RESTful API built in Golang with Json Web Token, Study Case Toy store (Work In Progress) 

💡 Roadmap RESTful API + JWT

    ✅ main.go

    ✅ .env

    ✅ config/config.go

    ✅ database/db.go

    🔜 models/user.go → make User & Credentials struct for login

    🔜 utils/jwt.go → function generate & validating token

    🔜 controllers/auth_controller.go → endpoint /login, /register

    🔜 middlewares/jwt_middleware.go → protect route using JWT

    🔜 routes/router.go → setup all endpoint

    🔜 models/toy.go → struct dan integrated produk

    🔜 controllers/toy_controller.go → endpoint CRUD toy

    🔜 Testing using Postman
