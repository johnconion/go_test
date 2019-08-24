package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func clockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <body>
            %d時%d分%d秒ですよ！今！
        </body>
        </html>
    `, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}

func main() {
	//ディレクトリを指定する
	// fs := http.FileServer(http.Dir("static"))
	//ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
	http.HandleFunc("/", clockHandler)

	log.Println("Listening...")
	// 3000ポートでサーバーを立ち上げる
	http.ListenAndServe(":3000", nil)
}
