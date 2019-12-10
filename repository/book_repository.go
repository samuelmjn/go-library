package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/samuelmjn/go-library/repository/model"
)

// BookRepository :nodoc:
type BookRepository interface {
	Create(req *model.Book) error
	Unissue(bookID int64) (err error)
	Issue(req *model.Issue) (err error)
	FindByID(id int64) (book *model.Book, err error)
	FindIssueByBookID(bookID int64) (issue *model.Issue, err error)
	FindMostIssued() (book *model.Book, err error)
	Delete(id int64) error
}

type bookRepo struct {
	db *gorm.DB
}

// NewBookRepository :nodoc:
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) Create(req *model.Book) (err error) {
	tx := r.db.Begin()
	err = tx.Create(&req).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepo) Issue(req *model.Issue) (err error) {
	tx := r.db.Begin()
	err = tx.Create(&req).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	var book model.Book
	err = tx.Where("id = ?", req.IssuedBook).First(book).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	book.IsIssued = true
	book.IssueCount++
	err = tx.Save(&book).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepo) Unissue(bookID int64) (err error) {
	tx := r.db.Begin()
	var book model.Book
	err = tx.Where("id = ?", bookID).First(&book).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	book.IsIssued = false
	err = tx.Save(book).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepo) FindByID(id int64) (book *model.Book, err error) {
	err = r.db.Where("id = ?", id).Take(&book).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (r *bookRepo) FindIssueByBookID(bookID int64) (issue *model.Issue, err error) {
	err = r.db.Where("issued_book = ?", bookID).Take(&issue).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (r *bookRepo) FindMostIssued() (book *model.Book, err error) {
	var maximumIssue int64
	row := r.db.Table("books").Select("MAX(issue_count)").Row()
	row.Scan(&maximumIssue)

	err = r.db.Where("issue_count = ?", maximumIssue).Take(&book).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (r *bookRepo) Delete(id int64) (err error) {
	tx := r.db.Begin()
	err = tx.Where("id = ?", id).Delete(&model.Book{}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return
}
