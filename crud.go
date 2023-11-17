package main

import (
	"fmt"
	"net/http"

	"github.com/ChickenClisp/golang_RestAPI_DB/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

// すべてのユーザを表示
func findAllUsers(w http.ResponseWriter, r *http.Request) {
	// DB接続
    db := utils.GetConnection()
    defer db.Close()

	// SELECT文が発行されて結果がuserListに入る
    var userList []User
	db.Table("users").Find(&userList)

	// JSONを返す
    utils.RespondWithJSON(w, http.StatusOK, userList)
}

// ユーザをIDで検索
func findById(w http.ResponseWriter, r *http.Request) {
	// IDを取得
	id, err := utils.GetID(r)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter")
        return
    }
	
    // DB接続
    db := utils.GetConnection()
    defer db.Close()

	// IDで検索
    var user User
    db.Where("id = ?", id).Find(&user)

    // JSONを返す
    utils.RespondWithJSON(w, http.StatusOK, user)
}

// INSERT　ユーザー追加処理
func createUser(w http.ResponseWriter, r *http.Request) {
    // リクエストボディ取得後、jSONを構造体に変換したものを取得
	var user User
    msg := utils.GetStruct(r, &user)
    if msg != "" {
        utils.RespondWithError(w, http.StatusBadRequest, msg)
        return
    }

    // DB接続
    db := utils.GetConnection()
    defer db.Close()

    // DBに新しいデータを挿入する
    db.Create(&user)

	// JSONを返す
    utils.RespondWithJSON(w, http.StatusOK, user)
}


// ユーザー更新
func updateUser(w http.ResponseWriter, r *http.Request) {
    // リクエストボディ取得後、jSONを構造体に変換したものを取得
	var user User
    msg := utils.GetStruct(r, &user)
    if msg != "" {
        utils.RespondWithError(w, http.StatusBadRequest, msg)
        return
    }

    // DB接続
    db := utils.GetConnection()
    defer db.Close()

    // DBを更新
    db.Save(&user)
    
	// JSONを返す
    utils.RespondWithJSON(w, http.StatusOK, user)
}

// ユーザー削除
func deleteUser(w http.ResponseWriter, r *http.Request) {
    // リクエストボディ取得後、jSONを構造体に変換したものを取得
	var user User
    msg := utils.GetStruct(r, &user)
    if msg != "" {
        utils.RespondWithError(w, http.StatusBadRequest, msg)
        return
    }

    // IDがセットされていない場合はエラーを返す
    if user.ID == 0 {
        utils.RespondWithError(w, http.StatusBadRequest, "ID is not set .")
        return
    }

    // DB接続
    db := utils.GetConnection()
    defer db.Close()

    // DELETE実行
    db.Delete(&user)

	// JSONを返す
    utils.RespondWithJSON(w, http.StatusOK, user)
}