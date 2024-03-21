package repository

import "github.com/lipzpoom/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Delete(noteId int)
	FindById(noteId int) (model.Note, error)
	FindAll() []model.Note
}
