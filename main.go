package main

import (
	"github.com/labstack/echo"
	"github.com/JermineHu/ait/common"
	"github.com/graphql-go/handler"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/JermineHu/ait/graphql"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/JermineHu/ait/route"
)
// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// a jwt test for login handler
func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {

		// Set custom claims
		claims := &jwtCustomClaims{
			"Jon Snow",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}


var (
	upgrader = websocket.Upgrader{}
)

// a websocket test for handler
func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	e := echo.New()

	h := handler.New(&handler.Config{
		Schema: graphql.GetGraphqlSchema().Scheam,
		Pretty: true,
		GraphiQL: true,
	})
	common.SetGraphqlHandler(e,"/graphql",h)


	// Save JSON of full schema introspection for Babel Relay Plugin to use
	//result := graphql.Do(graphql.Params{
	//	Schema:   testutil.StarWarsSchema,
	//	RequestString: testutil.IntrospectionQuery,
	//})
	//if result.HasErrors() {
	//	log.Fatalf("ERROR introspecting schema: %v", result.Errors)
	//	return
	//} else {
	//	b, err := json.MarshalIndent(result, "", "  ")
	//	if err != nil {
	//		log.Fatalf("ERROR: %v", err)
	//	}
	//	err = ioutil.WriteFile("./docs/schema.json", b, os.ModePerm)
	//	if err != nil {
	//		log.Fatalf("ERROR: %v", err)
	//	}
	//
	//}

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)


	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "assets")
	e.GET("/ws", hello)
	//e.Logger.Fatal(e.Start(":1323"))

	e.Debug=true
	e.HideBanner=true
	e.Logger.Fatal(route.Mean{e}.Engine().Echo.StartTLS(":443","certs/server.crt","certs/server.key"))

}