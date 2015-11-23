package helper

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func TrunctTime(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

func Decode(bytesData []byte, result interface{}) error {
	buf := bytes.NewBuffer(bytesData)
	dec := gob.NewDecoder(buf)
	e := dec.Decode(result)
	return e
}

func Encode(obj interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	gw := gob.NewEncoder(buf)
	err := gw.Encode(obj)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func PrintJSON(w http.ResponseWriter, success bool, data interface{}, message string) {
	w.Header().Set("Content-type", "application/json")

	result, err := json.Marshal(map[string]interface{}{
		"success": success,
		"data":    data,
		"message": message,
	})

	if err == nil {
		fmt.Fprintf(w, "%s\n", result)
	} else {
		result, _ := json.Marshal(map[string]interface{}{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})

		fmt.Fprintf(w, "%s\n", result)
	}
}

func BuildResponse(success bool, data interface{}, message string) interface{} {
	return map[string]interface{}{
		"success": success,
		"data":    data,
		"message": message,
	}
}

func FormatDuration(duration time.Duration) string {
	hours := int(math.Floor(duration.Hours()))
	minutes := int(math.Floor(math.Mod(duration.Minutes(), 60)))
	seconds := int(math.Floor(math.Mod(math.Mod(duration.Seconds(), 3600), 60)))
	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}

func Errorable(err error, callbacks ...func()) {
	if err != nil {
		fmt.Printf("Error %s\n", err.Error())
	}

	if len(callbacks) > 0 {
		callbacks[0]()
	}
}

func AsString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

func FloatToString(val interface{}) string {
	return fmt.Sprintf("%.2f", val)
}

func Float64ToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
