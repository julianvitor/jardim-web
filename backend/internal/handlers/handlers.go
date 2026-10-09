package handlers

import (
	"bytes"
	"net/http"
)

// Admin devices

func AdminGetDevices(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Admin!! "))
}
func AdminPostDevices(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Admin post Devices!! "))

}
func AdminPatchDevices(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello admin patch devices!!"))
}

func AdminDeleteDevices(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello admin delete devices"))
}
// Admin users

func AdminGetUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello admin get users"))
}

func AdminDeleteUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello admin delete users"))
}

// Admin telemetry
func AdminGetTelemetry(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello admin get telemetry"))
}

func AdminDeleteTelemetry(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello admin delete telemetry"))
}

// Hardware

func HardwarePostTelemetry(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello hardware post telemetry"))
}

func HardwareGetSettings(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello hardware get settings"))
}

// users
func UsersGetDevices (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user get devices"))
}
func UsersPostDevices (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user post devices"))
}
func UsersPatchDevices (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user patch devices"))
}
func UsersDeleteDevices (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user delete devices"))
}

func UsersGetTelemetry (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user get telemetry"))
}
func UsersGetSettings (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user get settings"))
}

func UsersPutSettings (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello user put settings"))
}

// Auth
func Login (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello login"))
}

func Register (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello register"))
}

func Logout (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello logout"))
}