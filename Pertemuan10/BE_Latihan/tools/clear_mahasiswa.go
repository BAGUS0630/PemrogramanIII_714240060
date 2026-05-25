package main

import (
    "BE_LATIHAN/config"
    "BE_LATIHAN/model"
    "log"
)

func main() {
    config.InitDB()
    db := config.GetDB()

    result := db.Where("1 = 1").Delete(&model.Mahasiswa{})
    if result.Error != nil {
        log.Fatalf("Gagal menghapus data mahasiswa: %v", result.Error)
    }

    log.Printf("Berhasil menghapus %d baris di tabel mahasiswa", result.RowsAffected)
}
