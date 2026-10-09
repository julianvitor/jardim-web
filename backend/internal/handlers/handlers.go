package handlers

import (
	"net/http"
)
// Admin
func AdminListDevices(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Admin!! "))
}
func AdminAddDevices(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello Admin Add Devices!! "))

}