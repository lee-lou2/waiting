package test

import (
	"waiting/cmd/server"
	"waiting/config/db"
)

// BaseData 기본 데이터
func BaseData() {
	tx, _ := db.GetDatabase()
	brand := server.Brand{Name: "test"}
	tx.Create(&brand)
	store := server.Store{Brand: brand, Name: "seoul"}
	tx.Create(&store)
	form1 := server.StoreForm{Store: store, Key: "key1", Type: "text_box"}
	tx.Create(&form1)
	form2 := server.StoreForm{Store: store, Key: "key2", Type: "select_box"}
	tx.Create(&form2)
	form3 := server.StoreForm{Store: store, Key: "key3", Type: "combo_box"}
	tx.Create(&form3)
	location := server.StoreLocation{Store: store, Latitude: "11.111", Longitude: "22.222"}
	tx.Create(&location)
}
