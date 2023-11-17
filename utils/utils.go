package utils

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
)

// エラー情報をJSONで返す
func RespondWithError(w http.ResponseWriter, code int, msg string) {
    RespondWithJSON(w, code, map[string]string{"error": msg})
}

// JSONを返す
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

// DBとのコネクションを張る
func GetConnection() *gorm.DB {
    db, err := gorm.Open("mysql", "ChickenClisp:Urgrn24yuyk#@/score_db?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatalf("DB connection failed %v", err)
    }
    // 詳細なログを表示
    db.LogMode(true)

    return db
}

// リクエストからIDを取得する
func GetID(r *http.Request) (id int, err error) {
    vars := mux.Vars(r)
    return strconv.Atoi(vars["id"])
}

// JSONリクエストを構造体にパースする
func GetStruct(r *http.Request, i interface{}) string {
    // リクエストボディ取得
    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        return "Invalid request"
    }
    // 読み込んだJSONを構造体に変換
    if err := json.Unmarshal(body, i); err != nil {
        return "JSON Unmarshaling failed ."
    }
    return ""
}
