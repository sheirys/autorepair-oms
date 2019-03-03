package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

func SegmentUInt32(mux map[string]string, seg string) uint32 {
	v, _ := strconv.Atoi(mux[seg])
	return uint32(v)
}

func SegmentBsonID(mux map[string]string, seg string) bson.ObjectId {
	v, _ := mux[seg]
	return bson.ObjectIdHex(v)
}

func SegmentString(mux map[string]string, seg string) string {
	v, _ := mux[seg]
	return string(v)
}

func BindJSON(r *http.Request, data interface{}) (bool, error) {
	err := json.NewDecoder(r.Body).Decode(&data)
	return err == nil, err
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	output, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(output)
	return
}
