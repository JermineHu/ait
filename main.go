package main

import (
	"github.com/labstack/echo"
	"github.com/jerminehu/ait/common"
	"github.com/graphql-go/handler"
	"./route"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/jerminehu/ait/graphql"
)

func init()  {


}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

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


	e.Debug=true
	e.HideBanner=true
	e.Logger.Fatal(route.Mean{e}.Engine().StartTLS(":443","certs/server.crt","certs/server.key"))

}