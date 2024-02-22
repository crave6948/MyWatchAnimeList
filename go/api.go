package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Anime คือโครงสร้างข้อมูลสำหรับ anime
type Anime struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Watch uint `json:"watch"`
	Image string `json:"image"`
}

var animeList []Anime
var maxID int

func main() {
	// โหลดข้อมูลจากไฟล์ JSON (หากมี)
	loadData()
	updateID()

	// สร้าง router ด้วย gorilla/mux
	router := mux.NewRouter()

	// Add CORS middleware
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// สร้างเส้นทาง (route) สำหรับ API endpoint
	router.HandleFunc("/anime", getAnimeList).Methods("GET")
	router.HandleFunc("/anime/add", addAnime).Methods("POST")
	router.HandleFunc("/anime/edit", updateAnime).Methods("PUT")
	router.HandleFunc("/anime/remove", removeAnime).Methods("DELETE")

	// เริ่มต้น worker สำหรับบันทึกข้อมูลทุก 1 นาที
	go saveDataEveryMinute()

	// กำหนดพอร์ตและเริ่มเซิร์ฟเวอร์
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	// ใช้ http.Handle และ CORS middleware จาก gorilla/handlers
	http.ListenAndServe(fmt.Sprintf(":%d", port), cors(router))
}

// Handler สำหรับการดึงข้อมูลทั้งหมดใน REST API
func getAnimeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animeList)
}

// Handler สำหรับการเพิ่ม anime ใน REST API
func addAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ดึงข้อมูลจาก request body
	var anime Anime
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&anime)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// กำหนด ID ให้ anime
	maxID++
	anime.ID = fmt.Sprintf("%d", maxID)

	// เพิ่ม anime ใน slice
	animeList = append(animeList, anime)
	// ส่ง response กลับไป
	json.NewEncoder(w).Encode(anime)
}
// Handler สำหรับการอัปเดตค่า Watch และ Image ของ anime ใน REST API
func updateAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ดึงข้อมูลจาก request body
	var updateData struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Watch uint   `json:"watch"`
		Image string `json:"image"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ค้นหา anime ที่ต้องการอัปเดต
	for i, anime := range animeList {
		if anime.ID == updateData.ID {
			// อัปเดตค่า Name, Watch และ Image ของ anime และอัปเดตใน slice
			animeList[i].Watch = updateData.Watch
			animeList[i].Image = updateData.Image
			animeList[i].Name = updateData.Name
			// ส่ง response กลับไป
			json.NewEncoder(w).Encode(animeList[i])
			return
		}
	}

	// หากไม่พบ anime ที่ต้องการอัปเดต
	http.Error(w, "Anime not found", http.StatusNotFound)
}

// Handler สำหรับการลบ anime ใน REST API
func removeAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ดึงข้อมูลจาก request body
	var removeData struct {
		ID    string `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&removeData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ค้นหาและลบ anime จาก slice
	for i, a := range animeList {
		if a.ID == removeData.ID {
			// ใช้ append เพื่อลบข้อมูลที่ต้องการ
			animeList = append(animeList[:i], animeList[i+1:]...)
			// ส่ง response กลับไป
			json.NewEncoder(w).Encode(a)
			return
		}
	}

	// หากไม่พบ anime ที่ต้องการลบ
	http.Error(w, "Anime not found", http.StatusNotFound)
}

// บันทึกข้อมูลลงในไฟล์ JSON
func saveData() {
	file, err := os.Create("anime_data.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(animeList)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Data saved to file at", time.Now().Format("2006-01-02 15:04:05"))
}

// โหลดข้อมูลจากไฟล์ JSON
func loadData() {
	file, err := os.Open("anime_data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&animeList)
	if err != nil {
		fmt.Println("Error decoding data:", err)
	}
}

// บันทึกข้อมูลทุก 1 นาที
func saveDataEveryMinute() {
	for {
		saveData()
		updateID()
		time.Sleep(time.Minute)
	}
}

func updateID() {
	maxID = 0
	updateAnimeListID()
	for _, anime := range animeList {
		id, _ := strconv.Atoi(anime.ID)
		if id > maxID {
			maxID = id
		}
	}	
}
// Function to update the ID of animeList from 0 to n
func updateAnimeListID() {
	for i := 0; i < len(animeList); i++ {
		animeList[i].ID = strconv.Itoa(i)
	}
}
