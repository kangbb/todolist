package controller
import (
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/kangbb/todolist/core/models/service"
)

func dataProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postData(w,r)
	} else if r.Method == "GET" {
		getData(w)
	} else {
		http.Error(w, "An Error taken place!", 501)
	}
}

//处理Post数据
//数据格式 {method:add/update/delete, label: string, isFinished: bool }
func postData(w http.ResponseWriter, r *http.Request) {
	method := r.Form["method"][0]
	if method == "add" {
		addData(w, r.Form)
	} else if method == "update" {
    updateData(w, r.Form)
	} else {
    deleteData(w, r.Form)
	}
}

func getData(w http.ResponseWriter) {
	data := service.GetAllItems()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	b, _ := json.Marshal(data)
	w.Write(b)
}

/*----------------------------------------------*/
func addData(w http.ResponseWriter, data map[string][]string) {
	isFinished, _ := strconv.ParseBool(data["isFinished"][0])
	_, err := service.CreateItem(data["label"][0], isFinished, data["time"][0])
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("unknown wrong!"))
	}
	w.WriteHeader(200)
}
func updateData(w http.ResponseWriter, data map[string][]string) {
	isFinished, _ := strconv.ParseBool(data["isFinished"][0])
  _, err := service.UpdateItem(data["label"][0], isFinished, data["time"][0])
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("unknown wrong!"))
	}
  w.WriteHeader(200)
}
func deleteData(w http.ResponseWriter, data map[string][]string) {
	isFinished, _ := strconv.ParseBool(data["isFinished"][0])
  _, err := service.DeleteItem(data["label"][0], isFinished, data["time"][0])
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("unknown wrong!"))
	}
  w.WriteHeader(200)
}