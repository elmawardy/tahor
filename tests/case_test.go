package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/elmawardy/tahor/global"
	"github.com/elmawardy/tahor/handlers"
	"github.com/elmawardy/tahor/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TestInsertCase(t *testing.T) {
	Init()
	var tests = []struct {
		Data     models.Case
		Metadata struct {
			Name string
		}
	}{
		{
			Data: models.Case{
				Comment:        "",
				Collected:      20000,
				Target:         100000,
				Currency:       "EGP",
				OrganizationID: 1,
			},
			Metadata: struct {
				Name string
			}{
				Name: "Basic insert",
			},
		},
	}

	w := httptest.NewRecorder()

	r := getRouter(true)

	r.POST("/api/case/insert", handlers.CaseInsertHandler())

	for _, tt := range tests {

		fmt.Printf("========= RUN Testing : %s\n", tt.Metadata.Name)
		payload, _ := json.Marshal(tt.Data)
		req, _ := http.NewRequest("POST", "/api/case/insert", strings.NewReader(string(payload)))
		req.Header.Add("Content-Type", "application/json")
		// req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fail()
		}

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Fail()
		}

		jsonData := struct {
			State string
			Msg   string
		}{}

		if err := json.Unmarshal(p, &jsonData); err != nil {
			t.Errorf("error unmarshalling response")
		}

		if jsonData.State != "Success" {
			t.Fail()
		}

		// restoreLists()

	}

}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	return r
}

func Init() {

	if global.DB == nil {
		// Check Database and migrations

		var err error
		global.DB, err = gorm.Open(global.Config["driver"], global.Config["constring"])
		defer global.DB.Close()
		if err != nil {
			panic("failed to connect database")
		}
		// Migrate the schema
		global.DB.AutoMigrate(&models.Case{}, &models.Organization{}, &models.Tag{}, &models.User{})

		//  Web Server
	}

}
