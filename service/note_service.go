package service

import (
	"github.com/lipzpoom/data/request"
	"github.com/lipzpoom/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteRespone
	FindAll() []response.NoteRespone
}
