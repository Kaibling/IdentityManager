package sqlite

import (
	"reflect"
	"testing"

	"github.com/Kaibling/IdentityManager/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&person{})
	if err != nil {
		panic(err)
	}
	return db
}

func TestSQLiteRepo_ReadAll(t *testing.T) {
	db := initDB()
	test1Data := []person{
		{Domain: "Dom1", Email: "email1"},
		{Domain: "Dom2", Email: "email2"},
		{Domain: "Dom3", Email: "email3"}}
	db.Create(&test1Data[0])
	db.Create(&test1Data[1])
	db.Create(&test1Data[2])
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Person
		wantErr bool
	}{
		{
			name:    "Test 1: read",
			fields:  fields{db: db},
			want:    personArrayUnmarshal(test1Data),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SQLiteRepo{
				db: tt.fields.db,
			}
			got, err := s.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLiteRepo.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLiteRepo.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteRepo_Create(t *testing.T) {
	db := initDB()

	test1Person := models.Person{Domain: "DomTest1", Email: "EmailTest1"}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		p models.Person
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1: create",
			fields:  fields{db: db},
			args:    args{p: test1Person},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SQLiteRepo{
				db: tt.fields.db,
			}
			if err := s.Create(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("SQLiteRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLiteRepo_Delete(t *testing.T) {
	db := initDB()
	test1Data := []person{
		{Domain: "Dom1", Email: "email1"},
		{Domain: "Dom2", Email: "email2"},
		{Domain: "Dom3", Email: "email3"}}
	db.Create(&test1Data[0])
	db.Create(&test1Data[1])
	db.Create(&test1Data[2])
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1: delete",
			fields:  fields{db: db},
			args:    args{domain: test1Data[1].Domain},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SQLiteRepo{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.domain); (err != nil) != tt.wantErr {
				t.Errorf("SQLiteRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLiteRepo_Update(t *testing.T) {
	db := initDB()
	test1Data := []person{
		{Domain: "Dom1", Email: "email1"}}
	db.Create(&test1Data[0])
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		p models.Person
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test 1: delete",
			fields:  fields{db: db},
			args:    args{p: models.Person{Domain: "Dom1", Email: "emailnew"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SQLiteRepo{
				db: tt.fields.db,
			}
			if err := s.Update(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("SQLiteRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
