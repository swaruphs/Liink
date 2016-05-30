package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

func WriteBadRequestError(w http.ResponseWriter, err error) {

	w.WriteHeader(http.StatusBadRequest)
	resp := map[string]string{}
	resp["error"] = err.Error()

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func SendInternalServerError(w http.ResponseWriter, err error) {

	w.WriteHeader(http.StatusInternalServerError)
	resp := map[string]string{}
	resp["error"] = err.Error()

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func SendResponse(w http.ResponseWriter, v interface{}) {

	w.WriteHeader(http.StatusOK)

	jsonResp, err := json.Marshal(v)

	if err != nil {
		SendInternalServerError(w, err)
		return
	}

	w.Write(jsonResp)
}

func SendEmptyResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func GetUserIDFromRequest(r *http.Request) (int, bool) {
	var iUserId int
	userID, ok := context.GetOk(r, "userid")
	fmt.Println(userID)

	if ok == false {
		return iUserId, ok
	}

	iUserId, ok = userID.(int)
	return iUserId, ok
}
