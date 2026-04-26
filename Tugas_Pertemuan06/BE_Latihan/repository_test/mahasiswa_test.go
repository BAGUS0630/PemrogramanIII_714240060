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

	err := config.GetDB().AutoMigrate(&model.Mahasiswa{})
	if err != nil {
		t.Fatalf("Gagal migrasi database: %v", err)
	}
}

func TestInsertMahasiswa(t *testing.T) {
	setupTest(t)

	mhs := model.Mahasiswa{
		NPM:    int64(1775465520665812100),
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

	npm := int64(1775465520665812100)

	mhs, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		t.Errorf("GetMahasiswaByNPM Gagal : %v", err)
	}

	if mhs.NPM != npm {
		t.Errorf("Expected NPM %d, got %d", npm, mhs.NPM)
	}

	fmt.Printf("Data Mahasiswa Ditemukan: %v\n", mhs)
}

func TestUpdateMahasiswa(t *testing.T) {
	setupTest(t)

	npm := int64(1775467813399742100)

	_, err := repository.UpdateMahasiswa(npm, &model.Mahasiswa{
		NPM:    int64(1775465520665812100),
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

	npm := int64(1775467813399742100)
	err := repository.DeleteMahasiswa(npm)
	if err != nil {
		t.Errorf("Delete Failed : %v", err)
	}

}
