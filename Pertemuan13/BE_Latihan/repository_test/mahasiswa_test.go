package repositorytest

import (
	"BE_LATIHAN/config"
	"BE_LATIHAN/model"
	"BE_LATIHAN/repository"
	"fmt"
	"testing"
)

func setupTest(t *testing.T) {
	config.InitDB()

	db := config.GetDB()
	err := db.AutoMigrate(&model.Mahasiswa{})
	if err != nil {
		t.Fatalf("Gagal migrasi database: %v", err)
	}

	// Bersihkan tabel sebelum tiap test agar idempotent
	if err := db.Where("1 = 1").Delete(&model.Mahasiswa{}).Error; err != nil {
		t.Fatalf("Gagal membersihkan tabel mahasiswa: %v", err)
	}
}

func TestInsertMahasiswa(t *testing.T) {
	setupTest(t)

	mhs := model.Mahasiswa{
		NPM:    "1775465",
		Nama:   "Bagus",
		Prodi:  "Teknik Informatika",
		Alamat: "Bandung",
		Email:  "mbagus0111@gmail.com",
		NoHP:   "085157392215",
		Hobi:   []string{"Bermain", "olahraga"},
	}

	_, err := repository.InsertMahasiswa(&mhs)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
}

func TestGetAllMahasiswa(t *testing.T) {
	setupTest(t)

	// masukkan satu record untuk memastikan ada data
	mhs := model.Mahasiswa{
		NPM:    "1775465",
		Nama:   "Bagus",
		Prodi:  "Teknik Informatika",
		Alamat: "Bandung",
		Email:  "mbagus0111@gmail.com",
		NoHP:   "085157392215",
		Hobi:   []string{"Bermain", "olahraga"},
	}
	_, err := repository.InsertMahasiswa(&mhs)
	if err != nil {
		t.Fatalf("Insert sebelum GetAll gagal: %v", err)
	}

	data, err := repository.GetAllMahasiswa()
	if err != nil {
		t.Errorf("GetAll Gagal : %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Data tidak ditemukan")
	}

	fmt.Printf("Berhasil. Data di table: %v\n", data)
}

func TestGetMahasiswaByNPM(t *testing.T) {
	setupTest(t)

	// insert terlebih dahulu
	mhsInsert := model.Mahasiswa{
		NPM:    "1775465",
		Nama:   "Bagus",
		Prodi:  "Teknik Informatika",
		Alamat: "Bandung",
		Email:  "mbagus0111@gmail.com",
		NoHP:   "085157392215",
		Hobi:   []string{"Bermain", "olahraga"},
	}
	_, err := repository.InsertMahasiswa(&mhsInsert)
	if err != nil {
		t.Fatalf("Insert sebelum GetMahasiswaByNPM gagal: %v", err)
	}

	npm := "1775465"

	mhs, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		t.Errorf("GetMahasiswaByNPM Gagal : %v", err)
	}

	if mhs.NPM != npm {
		t.Errorf("Expected NPM %s, got %s", npm, mhs.NPM)
	}

	fmt.Printf("Data Mahasiswa Ditemukan: %v\n", mhs)
}

func TestUpdateMahasiswa(t *testing.T) {
	setupTest(t)

	// insert dulu
	mhsInsert := model.Mahasiswa{
		NPM:    "1775465",
		Nama:   "Bagus",
		Prodi:  "Teknik Informatika",
		Alamat: "Bandung",
		Email:  "mbagus0111@gmail.com",
		NoHP:   "085157392215",
		Hobi:   []string{"Bermain", "olahraga"},
	}
	_, err := repository.InsertMahasiswa(&mhsInsert)
	if err != nil {
		t.Fatalf("Insert sebelum Update gagal: %v", err)
	}

	npm := "1775465"

	_, err = repository.UpdateMahasiswa(npm, &model.Mahasiswa{
		NPM:    "1775465",
		Nama:   "Muhammad Bagus Tri Atmaja",
		Prodi:  "D4 Teknik Informatika",
		Alamat: "Jl.cibiru, Kota Bandung",
		Email:  "mbagus0111@gmail.com",
		NoHP:   "085157392215",
		Hobi:   []string{"Bermain", "olahraga"},
	})

	if err != nil {
		t.Errorf("Update Failed : %v", err)
	}

}

func TestDeleteMahasiswa(t *testing.T) {
	setupTest(t)

	// insert dulu
	mhsInsert := model.Mahasiswa{
		NPM:    "1775465",
		Nama:   "Bagus",
		Prodi:  "Teknik Informatika",
		Alamat: "Bandung",
		Email:  "mbagus0111@gmail.com",
		NoHP:   "085157392215",
		Hobi:   []string{"Bermain", "olahraga"},
	}
	_, err := repository.InsertMahasiswa(&mhsInsert)
	if err != nil {
		t.Fatalf("Insert sebelum Delete gagal: %v", err)
	}

	npm := "1775465"
	err = repository.DeleteMahasiswa(npm)
	if err != nil {
		t.Errorf("Delete Failed : %v", err)
	}

}
