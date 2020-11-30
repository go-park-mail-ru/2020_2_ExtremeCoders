package test

import (
	"testing"
)

func TestServer(t *testing.T) {
	//var db = Postgres.DataBase{}
	//DataBase, err := db.Init(config.DbUser, config.DbPassword, config.DbDB)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//var uDB = UserPostgres.New(DataBase)
	//var uUC = UserUseCase.New(uDB)
	//var uDE = UserDelivery.New(uUC)
	//
	//var lDB = LetterPostgres.New(DataBase)
	//var lUC = LetterUseCase.New(lDB)
	//var lDE = LetterDelivery.New(lUC)
	//
	//mux := http.NewServeMux()
	//
	//mux.HandleFunc("/session", uDE.Session)
	//mux.HandleFunc("/user", uDE.Profile)
	//mux.HandleFunc("/user/avatar", uDE.GetAvatar)
	//mux.HandleFunc("/letter", lDE.SendLetter)
	//mux.HandleFunc("/user/letter/sent", lDE.GetSendLetters)
	//mux.HandleFunc("/user/letter/received", lDE.GetRecvLetters)
	//
	//siteHandler := middleware.AccessLogMiddleware(mux)
	//siteHandler = middleware.PanicMiddleware(siteHandler)
	//a := middleware.AuthMiddleware{Sessions: uDB}
	//siteHandler = a.Auth(siteHandler)
	//handler := cors.New(cors.Options{
	//	AllowedOrigins:   config.AllowedOriginsCORS,
	//	AllowedHeaders:   config.AllowedHeadersCORS,
	//	AllowedMethods:   config.AllowedMethodsCORS,
	//	AllowCredentials: true,
	//}).Handler(siteHandler)
	//File := http.Server{
	//	Addr:         config.Port,
	//	Handler:      handler,
	//	ReadTimeout:  config.ReadTimeout,
	//	WriteTimeout: config.WriteTimeout,
	//}
	//fmt.Println("starting File at ", config.Port)
	//File.ListenAndServe()
	//
	//
	//convey.Convey("Create should return ID of newly created user", func() {
	//	user := &User{Name: "Test user"}
	//	data, err := json.Marshal(user)
	//	convey.So(err, convey.ShouldBeNil)
	//	buf := bytes.NewBuffer(data)
	//
	//	req, err := http.NewRequest("POST", "http://localhost/api/v1/users", buf)
	//	convey.So(err, convey.ShouldBeNil)
	//	w := httptest.NewRecorder()
	//	File.
	//	r.ServeHTTP(w, req)
	//
	//	convey.So(w.Code, convey.ShouldAlmostEqual, http.StatusOK)
	//	body := strings.TrimSpace(w.Body.String())
	//	convey.So(body, convey.ShouldAlmostEqual, "1")
	//})

}
