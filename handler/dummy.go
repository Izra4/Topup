package handler

import (
	"TopUpWEb/database"
	"TopUpWEb/entity"
	"fmt"
	"gorm.io/gorm"
)

func GamesData() {
	g1 := entity.Games{
		Model:      gorm.Model{},
		Nama:       "Mobile Legends",
		Developer:  "Moonton",
		Gambar:     "https://veolwtbyepcnwgvrwnwu.supabase.co/storage/v1/object/public/Gambar_Game/1674803407ml%20(1).jpg?t=2023-06-28T05%3A42%3A27.856Z",
		Bookings:   nil,
		ListTopUps: nil,
	}
	g2 := entity.Games{
		Model:      gorm.Model{},
		Nama:       "PUBG Mobile",
		Developer:  "Tencent Games",
		Gambar:     "https://veolwtbyepcnwgvrwnwu.supabase.co/storage/v1/object/public/Gambar_Game/1674803473pubgm%20(1).jpg",
		Bookings:   nil,
		ListTopUps: nil,
	}

	g3 := entity.Games{
		Model:      gorm.Model{},
		Nama:       "Valorant",
		Developer:  "Riot",
		Gambar:     "https://veolwtbyepcnwgvrwnwu.supabase.co/storage/v1/object/public/Gambar_Game/1674803553valorant.jpg",
		Bookings:   nil,
		ListTopUps: nil,
	}
	if err := database.DB.Create(&g1).Error; err != nil {
		fmt.Println(err.Error())
	}

	if err := database.DB.Create(&g2).Error; err != nil {
		fmt.Println(err.Error())
	}

	if err := database.DB.Create(&g3).Error; err != nil {
		fmt.Println(err.Error())
	}
}

func MLTopUp() {
	t1 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "86",
		Harga:      "Rp. 20.000",
	}
	t2 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "172",
		Harga:      "Rp. 40.000",
	}
	t3 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "257",
		Harga:      "Rp. 60.000",
	}
	t4 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "706",
		Harga:      "Rp. 158.000",
	}
	t5 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "2.195",
		Harga:      "Rp. 465.000",
	}
	t6 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "3.688",
		Harga:      "Rp. 775.000",
	}
	t7 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "5.532",
		Harga:      "Rp. 1.162.000",
	}
	t8 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    1,
		JenisPaket: "9.288",
		Harga:      "Rp. 1.928.000",
	}

	if err := database.DB.Create(&t1).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t2).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t3).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t4).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t5).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t6).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t7).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t8).Error; err != nil {
		fmt.Println(err.Error())
	}
}

func PUBGTopUp() {
	t1 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "131 UC",
		Harga:      "Rp. 23.000",
	}
	t2 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "263 UC",
		Harga:      "Rp. 45.000",
	}
	t3 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "530 UC",
		Harga:      "Rp. 88.400",
	}
	t4 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "1100 UC",
		Harga:      "Rp. 175.000",
	}
	t5 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "1363 UC",
		Harga:      "Rp. 216.500",
	}
	t6 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "1630 UC",
		Harga:      "Rp. 262.000",
	}
	t7 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "2200 UC",
		Harga:      "Rp. 347.200",
	}
	t8 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "3300 UC",
		Harga:      "Rp. 518.600",
	}
	t9 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "4400 UC",
		Harga:      "Rp. 690.000",
	}
	t10 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "5500 UC",
		Harga:      "Rp. 862.000",
	}
	t11 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "6800 UC",
		Harga:      "Rp. 1.034.000",
	}
	t12 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "7700 UC",
		Harga:      "Rp. 1.208.000",
	}
	t13 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "8800 UC",
		Harga:      "Rp. 1.378.500",
	}
	t14 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "9900 UC",
		Harga:      "Rp. 1.549.800",
	}
	t15 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    2,
		JenisPaket: "11000 UC",
		Harga:      "Rp. 1.725.000",
	}

	if err := database.DB.Create(&t1).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t2).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t3).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t4).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t5).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t6).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t7).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t8).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t9).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t10).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t11).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t12).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t13).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t14).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t15).Error; err != nil {
		fmt.Println(err.Error())
	}
}

func ValorantTopUp() {
	t1 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "300 Points",
		Harga:      "Rp. 33.000",
	}
	t2 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "625 Points",
		Harga:      "Rp. 65.000",
	}
	t3 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "1125 Points",
		Harga:      "Rp. 112.500",
	}
	t4 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "1250 Points",
		Harga:      "Rp. 130.000",
	}
	t5 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "1650 Points",
		Harga:      "Rp. 162.000",
	}
	t6 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "2275 Points",
		Harga:      "Rp. 225.000",
	}
	t7 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "2775 Points",
		Harga:      "Rp. 271.000",
	}
	t8 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "3400 Points",
		Harga:      "Rp. 320.000",
	}
	t9 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "4525 Points",
		Harga:      "Rp. 419.000",
	}
	t10 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "5050 Points",
		Harga:      "Rp. 468.000",
	}
	t11 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "7000 Points",
		Harga:      "Rp. 623.500",
	}
	t12 := entity.ListTopUp{
		Model:      gorm.Model{},
		GamesID:    3,
		JenisPaket: "8650 Points",
		Harga:      "Rp. 791.000",
	}
	if err := database.DB.Create(&t1).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t2).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t3).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t4).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t5).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t6).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t7).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t8).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t9).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t10).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t11).Error; err != nil {
		fmt.Println(err.Error())
	}
	if err := database.DB.Create(&t12).Error; err != nil {
		fmt.Println(err.Error())
	}
}
