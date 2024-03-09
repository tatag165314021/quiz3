package controllers

import (
	"net/http"
	"quiz3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

var DataSegitiga = structs.Segitiga{}
var DataPersegi = structs.Persegi{}
var DataPersegiPanjang = structs.PersegiPanjang{}
var DataLingkaran = structs.Lingkaran{}

func GetSegitiga(ctx *gin.Context) {
	alas := ctx.Query("alas")
	tinggi := ctx.Query("tinggi")
	hitung := ctx.Query("hitung")

	DataSegitiga.Alas, _ = strconv.ParseFloat(alas, 64)

	DataSegitiga.Tinggi, _ = strconv.ParseFloat(tinggi, 64)

	if alas == "" || tinggi == "" || hitung == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "segitiga cant count",
		})
		return
	}

	if hitung == "keliling" {
		keliling := DataSegitiga.Alas * 3
		ctx.JSON(http.StatusOK, gin.H{
			"datasegitiga":      DataSegitiga,
			"keliling segitiga": keliling,
		})
	} else if hitung == "luas" {
		luas := DataSegitiga.Alas * DataSegitiga.Tinggi / 2
		ctx.JSON(http.StatusOK, gin.H{
			"datasegitiga":  DataSegitiga,
			"luas segitiga": luas,
		})
	}

}

func GetPersegi(ctx *gin.Context) {
	sisi := ctx.Query("sisi")
	hitung := ctx.Query("hitung")

	DataPersegi.Sisi, _ = strconv.Atoi(sisi)

	if sisi == "" || hitung == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "persegi cant count",
		})
		return
	}

	if hitung == "keliling" {
		keliling := DataPersegi.Sisi * 4
		ctx.JSON(http.StatusOK, gin.H{
			"datapersegi":      DataPersegi,
			"keliling persegi": keliling,
		})
	} else if hitung == "luas" {
		luas := DataPersegi.Sisi * DataPersegi.Sisi
		ctx.JSON(http.StatusOK, gin.H{
			"datapersegi":  DataSegitiga,
			"luas persegi": luas,
		})
	}
}

func GetPersegiPanjang(ctx *gin.Context) {
	panjang := ctx.Query("panjang")
	lebar := ctx.Query("lebar")
	hitung := ctx.Query("hitung")

	DataPersegiPanjang.Panjang, _ = strconv.Atoi(panjang)
	DataPersegiPanjang.Lebar, _ = strconv.Atoi(lebar)

	if panjang == "" || lebar == "" || hitung == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "persegi panjang cant count",
		})
		return
	}

	if hitung == "keliling" {
		keliling := 2 * (DataPersegiPanjang.Panjang + DataPersegiPanjang.Lebar)
		ctx.JSON(http.StatusOK, gin.H{
			"datapersegipanjang":       DataPersegiPanjang,
			"keliling persegi panjang": keliling,
		})
	} else if hitung == "luas" {
		luas := DataPersegiPanjang.Panjang * DataPersegiPanjang.Lebar
		ctx.JSON(http.StatusOK, gin.H{
			"datapersegipanjang":   DataPersegiPanjang,
			"luas persegi panajng": luas,
		})
	}
}

func GetLingkaran(ctx *gin.Context) {
	jariJari := ctx.Query("jariJari")
	hitung := ctx.Query("hitung")

	DataLingkaran.JariJari, _ = strconv.ParseFloat(jariJari, 64)

	if jariJari == "" || hitung == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "lingkaran cant count",
		})
		return
	}

	if hitung == "keliling" {
		keliling := 2 * 3.14 * DataLingkaran.JariJari
		ctx.JSON(http.StatusOK, gin.H{
			"datalingkaran":      DataLingkaran,
			"keliling lingkaran": keliling,
		})
	} else if hitung == "luas" {
		luas := 3.14 * DataLingkaran.JariJari * DataLingkaran.JariJari
		ctx.JSON(http.StatusOK, gin.H{
			"datalingkaran":  DataLingkaran,
			"luas lingkaran": luas,
		})
	}
}
