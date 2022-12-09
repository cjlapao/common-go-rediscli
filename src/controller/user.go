package controller

// import (
// 	"github.com/cjlapao/common-go-database/mongodb"
// 	"github.com/cjlapao/common-go-identity/models"
// 	"github.com/cjlapao/rediscli-go/databaseservice"
// )

// type UserContext struct{}

// func (u UserContext) GetUserById(id string) *models.User {
// 	var result models.User
// 	repo := GetRepository()
// 	dbUsers := repo.FindOne("id", id)
// 	dbUsers.Decode(&result)
// 	return &result
// }

// func (u UserContext) GetUserByEmail(email string) *models.User {
// 	var result models.User
// 	repo := GetRepository()
// 	dbUsers := repo.FindOne("email", email)
// 	dbUsers.Decode(&result)
// 	return &result
// }

// func (u UserContext) UpsertUser(user models.User) {

// }

// func GetRepository() mongodb.Repository {
// 	factory, database := databaseservice.GetDatabase()
// 	userRepo := mongodb.NewRepository(factory, database, identity.IdentityUsersCollection)
// 	return userRepo
// }
