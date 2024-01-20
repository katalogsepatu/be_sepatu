package module

import (
	"encoding/json"
	"net/http"
	"os"

	model "github.com/katalogsepatu/be_sepatu/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	credential model.Credential
	response   model.Response
	user       model.User
	password   model.UpdatePassword
)

func SignUpHandler(MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	email, err := SignUp(conn, collectionname, user)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil SignUp"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data": bson.M{
			"email": email,
		},
	}
	return GCFReturnStruct(responData)
}

func LogInHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	user, err := LogIn(conn, collectionname, user)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	tokenstring, err := Encode(user.ID, user.Email, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		response.Message = "Gagal Encode Token : " + err.Error()
		return GCFReturnStruct(response)
	}
	//
	credential.Message = "Selamat Datang " + user.Fullname
	credential.Token = tokenstring
	credential.Status = 200
	responData := bson.M{
		"status":  credential.Status,
		"message": credential.Message,
		"data": bson.M{
			"token": credential.Token,
			"email": user.Email,
		},
	}
	return GCFReturnStruct(responData)
}

// func GetProfileHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
// 	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
// 	response.Status = 400
// 	//
// 	payload, err := GetUserLogin(PASETOPUBLICKEYENV, r)
// 	if err != nil {
// 		response.Message = err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	user, err := GetUserFromID(payload.Id, conn)
// 	if err != nil {
// 		response.Message = err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	//
// 	response.Status = 200
// 	response.Message = "Get Success"
// 	responData := bson.M{
// 		"status":  response.Status,
// 		"message": response.Message,
// 		"data": bson.M{
// 			"_id":          user.ID,
// 			"nama_lengkap": user.Fullname,
// 			"email":        user.Email,
// 			"phonenumber":  user.PhoneNumber,
// 		},
// 	}
// 	return GCFReturnStruct(responData)
// }

// func EditProfileHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
// 	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
// 	response.Status = 400
// 	//
// 	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
// 	if err != nil {
// 		response.Message = "Gagal Decode Token : " + err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	err = json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		response.Message = "error parsing application/json: " + err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	data, err := EditProfile(user.Id, conn, r)
// 	if err != nil {
// 		response.Message = err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	//
// 	response.Status = 200
// 	response.Message = "Berhasil mengubah profile" + user.Email
// 	responData := bson.M{
// 		"status":  response.Status,
// 		"message": response.Message,
// 		"data":    data,
// 	}
// 	return GCFReturnStruct(responData)
// }

func EditPasswordHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = "Gagal Decode Token : " + err.Error()
		return GCFReturnStruct(response)
	}
	err = json.NewDecoder(r.Body).Decode(&password)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	data, err := EditPassword(user.Id, conn, password)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil mengubah password" + user.Email
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

func EditEmailHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user_login, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = "Gagal Decode Token : " + err.Error()
		return GCFReturnStruct(response)
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	data, err := EditEmail(user_login.Id, conn, user)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil mengubah email" + user_login.Email
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

// Katalog Sepatu
func TambahKatalogSepatuHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses, email anda : " + user.Email
		return GCFReturnStruct(response)
	}
	data, err := PostKatalogSepatu(conn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 201
	response.Message = "Berhasil menambah Katalog Sepatu"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

func GetKatalogSepatuHandler(MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	id := GetID(r)
	if id == "" {
		data, err := GetAllKatalogSepatu(conn, collectionname)
		if err != nil {
			response.Message = err.Error()
			return GCFReturnStruct(response)
		}
		responData := bson.M{
			"status":  200,
			"message": "Get Success",
			"data":    data,
		}
		//
		return GCFReturnStruct(responData)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	katalogsepatu, err := GetKatalogSepatuById(conn, collectionname, idparam)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Get Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    katalogsepatu,
	}
	return GCFReturnStruct(responData)
}

func EditKatalogSepatuHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	id := GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return GCFReturnStruct(response)
	}
	data, err := PutKatalogSepatu(idparam, conn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil mengubah Katalog Sepatu"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

func DeleteKatalogSepatuHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	id := GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return GCFReturnStruct(response)
	}
	err = DeleteKatalogSepatu(idparam, collectionname, conn)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 204
	response.Message = "Berhasil menghapus Katalog Spot"
	return GCFReturnStruct(response)
}

// Favorite Sepatu
func TambahFavoriteSepatuHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses, email anda : " + user.Email
		return GCFReturnStruct(response)
	}
	data, err := PostFavoriteSepatu(conn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 201
	response.Message = "Berhasil menambah Katalog Sepatu"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

func GetFavoriteSepatuHandler(MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	id := GetID(r)
	if id == "" {
		data, err := GetAllFavoriteSepatu(conn, collectionname)
		if err != nil {
			response.Message = err.Error()
			return GCFReturnStruct(response)
		}
		responData := bson.M{
			"status":  200,
			"message": "Get Success",
			"data":    data,
		}
		//
		return GCFReturnStruct(responData)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	favoritesepatu, err := GetFavoriteSepatuById(conn, collectionname, idparam)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Get Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    favoritesepatu,
	}
	return GCFReturnStruct(responData)
}

func EditFavoriteSepatuHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	id := GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return GCFReturnStruct(response)
	}
	data, err := PutFavoriteSepatu(idparam, conn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil mengubah Katalog Sepatu"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

func DeleteFavoriteSepatuHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "admin@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	id := GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return GCFReturnStruct(response)
	}
	err = DeleteKatalogSepatu(idparam, collectionname, conn)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 204
	response.Message = "Berhasil menghapus Katalog Spot"
	return GCFReturnStruct(response)
}
