package controllers

import (
	//"html/template"
	//"log"
	//"database/sql"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qadrina/go-jwt-university-app/initializers"
	"github.com/qadrina/go-jwt-university-app/models"
	"golang.org/x/crypto/bcrypt"
)

// func Welcome(w http.ResponseWriter, r *http.Request) {
// 	temp, err := template.ParseFiles("static/index.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	temp.Execute(w, nil)
// }

func Welcome(c *gin.Context) {
	c.Request.URL.Path = "/static/index.html"
	//r.HandleContext(c)
}

func SignUp(c *gin.Context) {
	var body struct {
		ID        string
		Email     string
		Password  string
		FirstName string
		LastName  string
		FacultyID string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	student := models.Student{ID: (body.FacultyID + randSeq(5)), Email: body.Email, Password: string(hash), FirstName: body.FirstName, LastName: body.LastName, FacultyID: body.FacultyID}

	result := initializers.DB.Create(&student)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Student Account",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create user succeeded",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var studs models.Student
	initializers.DB.First(&studs, "email = ?", body.Email)

	if studs.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(studs.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": studs.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Token successfully created",
		"token":   tokenString,
	})
}

// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

func Validate(c *gin.Context) {
	student, _ := c.Get("student")
	c.JSON(http.StatusOK, gin.H{
		"message": student,
	})
}

func GetAllStudents(c *gin.Context) {
	/*
		err2:= db.QueryRow( "EXEC PRCENVIACOMANDO @IDVEICULO=?, @CMD=?",urlqry.Get("idveiculo") ,2).Scan(&ret)
	*/
	// err := initializers.DB.Exec(`EXEC spGetAllStudents;`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//var result sql.NullString

	//var students []Student

	var students []models.Student
	initializers.DB.Raw("EXEC spGetAllStudents;").Scan(&students)
	c.JSON(200, gin.H{
		"result": students,
	})
}

/*
package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mssql"
    "github.com/jinzhu/gorm"
)

func main() {
    db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := gin.Default()
    router.GET("/call-stored-procedure", func(c *gin.Context) {
        var result sql.NullString
        db.Raw("EXEC stored_procedure_name @param1 = ?, @param2 = ?", "value1", "value2").Scan(&result)
        c.JSON(200, gin.H{
            "result": result.String,
        })
    })

    router.Run(":8080")
}


*/

/*
package main

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mssql"
    "github.com/jinzhu/gorm"
)

type Student struct {
    ID        int    `gorm:"column:id"`
    FirstName string `gorm:"column:first_name"`
    LastName  string `gorm:"column:last_name"`
}

func main() {
    db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := gin.Default()
    router.GET("/get-all-students", func(c *gin.Context) {
        var students []Student
        db.Table("Student").Find(&students)
        c.JSON(200, gin.H{
            "students": students,
        })
    })

    router.Run(":8080")
}

*/
