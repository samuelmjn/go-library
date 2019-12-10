package repository

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/samuelmjn/go-library/repository/model"
	"github.com/samuelmjn/go-library/utils"
)

// BookRepository :nodoc:
type BookRepository interface {
	Create(req *model.Book) error
	Unissue(issueID int64) (err error)
	Issue(req *model.Issue) (err error)
	FindByID(id int64) (book *model.Book, err error)
	FindCurrentIssueByBookID(bookID int64) (issue *model.Issue, err error)
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
	// Define defaults
	req.ID = utils.GenerateID()
	req.IsIssued = false
	req.IssueCount = 0

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
	// Generate Defaults
	req.ID = utils.GenerateID()

	err = tx.Create(&req).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	var book model.Book
	err = tx.Where("id = ? AND is_issued = ?", req.IssuedBook, false).First(&book).Error
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

func (r *bookRepo) Unissue(issueID int64) (err error) {
	tx := r.db.Begin()

	var issue model.Issue
	err = tx.Where("id = ? AND return_time IS NULL", issueID).First(&issue).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	currentTime := time.Now()
	issue.ReturnTime = &currentTime
	err = tx.Save(issue).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	var book model.Book
	err = tx.Where("id = ?", issue.IssuedBook).First(&book).Error
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
	var res model.Book
	err = r.db.Where("id = ?", id).First(&res).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	book = &res
	return
}

func (r *bookRepo) FindCurrentIssueByBookID(bookID int64) (issue *model.Issue, err error) {
	var res model.Issue
	err = r.db.Where("issued_book = ? AND return_time IS NULL", bookID).First(&res).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	issue = &res
	return
}

func (r *bookRepo) FindMostIssued() (book *model.Book, err error) {
	var maximumIssue int64
	row := r.db.Table("books").Select("MAX(issue_count)").Row()
	row.Scan(&maximumIssue)

	var res model.Book
	err = r.db.Where("issue_count = ?", maximumIssue).Take(&res).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	book = &res
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
