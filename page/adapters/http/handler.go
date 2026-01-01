package http

import (
	"encoding/json"
	"errors"
	"io"
	nethttp "net/http"

	page "Trilhos/page/core"
)

type Handler struct {
	service *page.Service
}

func NewHandler(service *page.Service) *Handler {
	return &Handler{service: service}
}

func mapError(err error) (status int, code, message string) {
	if err == nil {
		return nethttp.StatusInternalServerError, "internal_error", "erro interno nao mapeado"
	}
	var syntaxErr *json.SyntaxError
	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &syntaxErr) || errors.As(err, &typeErr) || errors.Is(err, io.EOF) {
		return nethttp.StatusBadRequest, "invalid_json", "payload json invalido"
	}
	if errors.Is(err, page.ErrNotFound) {
		return nethttp.StatusNotFound, "not_found", "pagina nao encontrada com o id informado"
	}
	if errors.Is(err, page.ErrInvalidTitle) {
		return nethttp.StatusBadRequest, "invalid_title", "titulo precisa ser informado"
	}
	if errors.Is(err, page.ErrInvalidValue) {
		return nethttp.StatusBadRequest, "invalid_value", "titulo precisa ter entre 1 e 100 caracteres"
	}
	if errors.Is(err, page.ErrInvalidTime) {
		return nethttp.StatusInternalServerError, "invalid_time", "erro na validacao da data de atualizacao"
	}
	if errors.Is(err, page.ErrInvalidID) {
		return nethttp.StatusBadRequest, "invalid_id", "id nao pode ser vazio"
	}
	return nethttp.StatusInternalServerError, "internal_error", "erro interno nao mapeado"
}

func writeError(w nethttp.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Error:   code,
		Message: message,
	})
}

func (h *Handler) Create(w nethttp.ResponseWriter, r *nethttp.Request) {
	var req CreatePageRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		status, code, message := mapError(err)
		writeError(w, status, code, message)
		return
	}
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			writeError(w, nethttp.StatusBadRequest, "invalid_json", "payload deve conter um unico objeto json")
			return
		}
		status, code, message := mapError(err)
		writeError(w, status, code, message)
		return
	}

	p, err := h.service.Create(r.Context(), req.Title)
	if err != nil {
		status, code, message := mapError(err)
		writeError(w, status, code, message)
		return
	}

	resp := PageResponse{
		ID:        string(p.ID),
		Title:     p.Title,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(nethttp.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}
