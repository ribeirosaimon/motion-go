package Company

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/ribeirosaimon/motion-go/internal/db"
	"github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"
	"github.com/ribeirosaimon/motion-go/internal/repository"
	"github.com/ribeirosaimon/motion-go/test"
	"github.com/stretchr/testify/assert"
)

var (
	e         = test.CreateEngine(NewCompanyRouter)
	companyId uint64
)

func TestSaveCompanyController(t *testing.T) {

	company := createCompany()
	jsonData, _ := json.Marshal(company)
	w := test.PerformRequest(e, http.MethodPost, "/company", "ADMIN", bytes.NewReader(jsonData))

	var res sqlDomain.Company
	json.Unmarshal([]byte(w.Body.String()), &res)

	conn := db.Conn.GetPgsqTemplate()
	defer db.Conn.ClosePostgreSQL()
	dbCompany, _ := repository.NewCompanyRepository(conn).FindById(res.Id)

	assert.NotEqual(t, 0, res.Id)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, company.Name, dbCompany.Name)
	assert.Equal(t, company.Image, dbCompany.Image)
	companyId = dbCompany.Id
}

func TestGetCompanyController(t *testing.T) {
	conn := db.Conn.GetPgsqTemplate()
	defer db.Conn.ClosePostgreSQL()
	companyRepository := repository.NewCompanyRepository(conn)

	newCompany := createCompany()
	savedCompany, _ := companyRepository.Save(newCompany)
	w := test.PerformRequest(e, http.MethodGet, fmt.Sprintf("/company/%d", savedCompany.Id), "ADMIN", nil)

	var responseCompany sqlDomain.Company
	json.Unmarshal([]byte(w.Body.String()), &responseCompany)

	foundCompany, _ := companyRepository.FindById(responseCompany.Id)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, foundCompany.Id, responseCompany.Id)
	assert.Equal(t, foundCompany.Name, responseCompany.Name)
	assert.Equal(t, foundCompany.Image, responseCompany.Image)
}

func TestPutCompanyController(t *testing.T) {
	conn := db.Conn.GetPgsqTemplate()
	defer db.Conn.ClosePostgreSQL()
	companyRepository := repository.NewCompanyRepository(conn)
	companyDb, _ := companyRepository.Save(createCompany())

	updatedCompany := createCompany()
	companyToUpdate, _ := json.Marshal(updatedCompany)
	w := test.PerformRequest(e, http.MethodPut, fmt.Sprintf("/company/%d", companyDb.Id), "ADMIN", bytes.NewReader(companyToUpdate))

	var response sqlDomain.Company
	json.Unmarshal([]byte(w.Body.String()), &response)

	dbCompany, _ := companyRepository.FindById(response.Id)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, companyDb.Id, dbCompany.Id)
	assert.Equal(t, updatedCompany.Name, dbCompany.Name)
	assert.Equal(t, updatedCompany.Image, dbCompany.Image)
}

func createCompany() sqlDomain.Company {
	rand.Seed(time.Now().UnixNano())
	nameRandom := strconv.Itoa(rand.Intn(1000000))
	imageRandom := strconv.Itoa(rand.Intn(1000000))

	var company sqlDomain.Company
	company.Name = nameRandom
	company.Image = fmt.Sprintf("https://%s", imageRandom)
	return company
}