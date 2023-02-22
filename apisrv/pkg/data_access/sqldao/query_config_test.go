package sqldao_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/sqldao"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Employee struct {
	ID         int    `gorm:"column:employee_id;primarykey" json:"id,omitempty"`
	Department string `gorm:"column:department" json:"department,omitempty"`
	Salary     int    `gorm:"column:salary" json:"salary,omitempty"`
}

func TestApplyQueryConfig(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	// init db client
	dataDir := os.Getenv("DATA_DIR")
	client, err := gorm.Open(
		sqlite.Open(dataDir+"test-employees.sqlite.db"),
		&gorm.Config{},
	)
	assert.NoError(t, err)

	// apply query config and run
	conf := sqldao.QueryConfig{
		Fields: []string{"ID", "Department"},
		Limit:  5,
		Offset: 1,
		// PreloadAssociations: []string{"test_association"},
	}

	var eSlice []*Employee
	err = conf.Apply(client).
		Model(Employee{}).
		Find(&eSlice).
		Error
	assert.NoError(t, err)

	// // print debug
	// mar, err := json.MarshalIndent(eSlice, "", "\t")
	// assert.NoError(t, err)
	// fmt.Println(string(mar))
}
