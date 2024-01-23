package be_sepatu

import (
	"fmt"
	"testing"

	model "github.com/katalogsepatu/be_sepatu/model"
	module "github.com/katalogsepatu/be_sepatu/module"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "sepatu_db")
var collectionnameUser = "user"

// var collectionnameFishingspot = "fishingspot"

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
}

func TestSignUp(t *testing.T) {
	conn := db
	var user model.User
	user.Fullname = "Park Jisung"
	user.Email = "jisung@gmail.com"
	user.Password = "jisung123"
	user.ConfirmPassword = "jisung123"
	user.PhoneNumber = "622109013241"
	email, err := module.SignUp(conn, collectionnameUser, user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil SignUp : ", email)
	}
}

func TestLogInn(t *testing.T) {
	conn := db
	var user model.User
	user.Email = "admin@gmail.com"
	user.Password = "admin12345678"
	user, _ = module.LogIn(conn, collectionnameUser, user)
	tokenstring, err := module.Encode(user.ID, user.Email, "33186fcfc13ba9946bf200cf6c7808e6ebfc605140f65809e06648985b08ebda2df976efd75eacf2a37b1ce184deec8d3b72cb78f7881ed5e7a02d97351c2aef")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil LogIn : ", user.Fullname)
		fmt.Print("Berhasil LogIn : " + tokenstring)
	}
}

func TestToken(*testing.T) {
	token := "v4.public.eyJleHAiOiIyMDI0LTAxLTA0VDExOjI1OjU0WiIsImZ1bGxuYW1lIjoiYWRtaW5AZ21haWwuY29tIiwiaWF0IjoiMjAyNC0wMS0wNFQwOToyNTo1NFoiLCJpZCI6IjY1OTY1ZWNkY2MxOGQxNmNkNGNhNGY4YSIsIm5iZiI6IjIwMjQtMDEtMDRUMDk6MjU6NTRaIn22kA21UMcQv-6lNrkBu88rV3XGGgToTBqulQui3HrZcYb_Go-qyCBdzje7Qg3Omj-hI5lXRRFj1afCzeMdyG0B"
	tokenstring, err := module.Decode("2df976efd75eacf2a37b1ce184deec8d3b72cb78f7881ed5e7a02d97351c2aef", token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Id Token : " + tokenstring.Id.Hex())
		fmt.Print("Email Token : " + tokenstring.Email)
	}
}

// func TestCheckLatitudelongitude(t *testing.T) {
// 	err := module.CheckLatitudeLongitude(db, "1234", "1234")
// 	fmt.Println(err)
// }

// Katalog Sepatu
func TestTambahKatalogSepatu(t *testing.T) {
	id, err := module.InsertOneDoc(db, "katalogsepatu", model.KatalogSepatu{
		ID:       primitive.NewObjectID(),
		Brand:    "Adidas",
		Name:     "Adidas Samba",
		Category: "Sepatu Pria",
		Price:    "1500000",
		Color:    "Hitam",
		Diskon:   "25%",
		Image: "https://www.google.com",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil TambahKatalogSepatu : ", id)
	}
}

func TestUpdateKatalogSepatu(t *testing.T) {
	id := "65ab46c7c0ba7ca17d5e23cc"
	objectId, err := primitive.ObjectIDFromHex(id)

	data := module.UpdateOneDoc(objectId, db, "katalogsepatu", model.KatalogSepatu{
		ID:       objectId,
		Brand:    "Adidas",
		Name:     "Superstar Shoes",
		Category: "Sepatu Pria",
		Price:    "1000000",
		Color:    "Hitam",
		Diskon:   "25%",
		Image: "https://www.google.com",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil UpdateKatalogSepatu", data)
	}
}

func TestDeleteKatalogSepatu(t *testing.T) {
	conn := db
	id := "65ab46c7c0ba7ca17d5e23cc"
	objectId, err := primitive.ObjectIDFromHex(id)
	err = module.DeleteKatalogSepatu(objectId, "katalogsepatu", conn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Delete Katalog Sepatu")
	}
}


// Favorite Sepatu
func TestTambahFavoriteSepatu(t *testing.T) {
	id, err := module.InsertOneDoc(db, "favoritesepatu", model.FavoriteSepatu{
		ID:       primitive.NewObjectID(),
		Brand:    "Adidas",
		Name:     "Adidas Samba",
		Category: "Sepatu Pria",
		Price:    "1500000",
		Color:    "Hitam",
		Diskon:   "25%",
		Image: "https://www.google.com",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Tambah Favorite Sepatu : ", id)
	}
}

func TestUpdateFavoriteSepatu(t *testing.T) {
	id := "65ab470fbcc1dce41393dbb9"
	objectId, err := primitive.ObjectIDFromHex(id)

	data := module.UpdateOneDoc(objectId, db, "favoritesepatu", model.FavoriteSepatu{
		ID:       objectId,
		Brand:    "Adidas",
		Name:     "Superstar Shoes",
		Category: "Sepatu Pria",
		Price:    "1000000",
		Color:    "Hitam",
		Diskon:   "25%",
		Image: "https://www.google.com",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Update Favorite Sepatu", data)
	}
}

func TestDeleteFavoriteSepatu(t *testing.T) {
	conn := db
	id := "65ab470fbcc1dce41393dbb9"
	objectId, err := primitive.ObjectIDFromHex(id)
	err = module.DeleteFavoriteSepatu(objectId, "favoritesepatu", conn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Delete Favorite Sepatu")
	}
}

// Kategori Sepatu
func TestTambahKategoriSepatu(t *testing.T) {
	id, err := module.InsertOneDoc(db, "kategorisepatu", model.KategoriSepatu{
		ID:       primitive.NewObjectID(),
		Brand:    "Adidas",
		Name:     "Adidas Samba",
		Category: "Sepatu Pria",
		Price:    "1500000",
		Color:    "Hitam",
		Diskon:   "25%",
		Image: "https://www.google.com",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Tambah Kategori Sepatu : ", id)
	}
}

func TestUpdateKategoriSepatu(t *testing.T) {
	id := "65ab470fbcc1dce41393dbb9"
	objectId, err := primitive.ObjectIDFromHex(id)

	data := module.UpdateOneDoc(objectId, db, "kategorisepatu", model.KategoriSepatu{
		ID:       objectId,
		Brand:    "Adidas",
		Name:     "Superstar Shoes",
		Category: "Sepatu Pria",
		Price:    "1000000",
		Color:    "Hitam",
		Diskon:   "25%",
		Image: "https://www.google.com",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Update Kategori Sepatu", data)
	}
}

func TestDeleteKategoriSepatu(t *testing.T) {
	conn := db
	id := "65ab470fbcc1dce41393dbb9"
	objectId, err := primitive.ObjectIDFromHex(id)
	err = module.DeleteKategoriSepatu(objectId, "kategorisepatu", conn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Delete Kategori Sepatu")
	}
}


// func TestTambahFishingSpot(t *testing.T) {
//     conn := db
//     var spot model.FishingSpot
//     spot.Name = "Spot 1"
//     spot.Phonenumber = "6285718177810"
//     spot.TopFish = "Ikan 1"
//     spot.Rating = "5"
//     spot.OpeningHour = "08:00 - 17:00"
//     spot.Description = "Deskripsi Spot 1"
//     spot.Image = "https://www.google.com"
//     spot.Address = "Alamat Spot 1"
//     spot.Longitude = "0.000000"
//     spot.Latitude = "0.000000"

//     // Perbaikan #1: Memastikan tipe data return dan argumen yang benar
//     _, err := module.PostFishingSpot(conn, collectionnameFishingspot, &spot)
//     if err != nil {
//         fmt.Println(err)
//     } else {
//         fmt.Println("Berhasil TambahFishingSpot : ")
//     }
// }
