package routes

import (
	"go_mysql/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           //책 등록
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")               // 책 리스트
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   // ID로 책정보 가져옴
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    // ID값으로 책정보 업데이트
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") // ID값으로 책정보 삭제
}
