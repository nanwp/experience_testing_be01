package main

//exercise testing BE01

//buat func PembayaranBarang
//parameter : hargaTotal (float64), metode pembayaran (string), dicicil (bool)

//func PembayaranBarang(ht float64, m string, cicil bool) error

//- cek hargaTotal > 0 apabila <= 0 maka return error "harga tidak bisa nol"

//- cek metode pembayaran: "cod", "transfer", "debit", "credit", "gerai" kalo tidak yg diatas maka return error "metode tidak dikenali"

//- cek dicicil atau tidak, kalo true, metode harus credit, dan hargaTotal harus >= 500.000 return error "cicilan tidak memenuhi syarat"

//kalo false, metode bukan credit. return error "credit harus dicicil"

import (
	"errors"
)

func PembayaranBarang(ht float64, metode string, cicil bool) error {
	if ht <= 0 {
		return errors.New("harga tidak bisa nol")
	}
	if metode != "cod" && metode != "transfer" && metode != "debit" && metode != "credit" && metode != "gerai" {
		return errors.New("metode tidak dikenali")
	}
	if cicil {
		if metode != "credit" {
			return errors.New("credit harus dicicil")
		}
		if ht < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		if metode == "credit" {
			return errors.New("credit harus dicicil")
		}
	}
	return nil

}
