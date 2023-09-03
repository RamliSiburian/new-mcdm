package handlers

import (
	"encoding/json"
	dto "mcdm/Dto"
	models "mcdm/Models"
	repositories "mcdm/Repositories"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerperbandinganMopa struct {
	PerbandinganMopaRepository repositories.PerbandinganMopaRepository
}

func HandlerperbandinganMopa(PerbandinganMopaRepository repositories.PerbandinganMopaRepository) *handlerperbandinganMopa {
	return &handlerperbandinganMopa{PerbandinganMopaRepository}
}

func (h *handlerperbandinganMopa) FindPerbandinganMopa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kriteria, err := h.PerbandinganMopaRepository.FindPerbandinganMopa()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: kriteria}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerperbandinganMopa) InsertPerbandinganMopa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.PerbandinganMopaRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// validation
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	createPerbandinganMopa := models.PerbandinganMopa{
		Kode:      request.Kode,
		Deskripsi: request.Deskripsi,
		Nilai:     request.Nilai,
		CreatedAt: time.Now(),
	}

	data, err := h.PerbandinganMopaRepository.CreatePerbandinganMopa(createPerbandinganMopa)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerperbandinganMopa) GetPerbandinganMopaByKode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kode, _ := mux.Vars(r)["kode"]

	var perbandinganMopa models.PerbandinganMopa

	perbandinganMopa, err := h.PerbandinganMopaRepository.GetPerbandinganMopaByKode(kode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: perbandinganMopa}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerperbandinganMopa) UpdatePerbandinganMopa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.PerbandinganMopaRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kodePerbandinganMopa := mux.Vars(r)["kode"]
	oldData, err := h.PerbandinganMopaRepository.GetPerbandinganMopaByKode(kodePerbandinganMopa)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Kode >= 0 {
		oldData.Kode = request.Kode
	}

	if request.Deskripsi != "" {
		oldData.Deskripsi = request.Deskripsi
	}
	if request.Nilai >= 0 {
		oldData.Nilai = request.Nilai
	}

	data, err := h.PerbandinganMopaRepository.UpdatePerbandinganMopa(oldData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerperbandinganMopa) DeletePerbandinganMopa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kodePerbandinganMopa := mux.Vars(r)["kode"]
	perbandingan, err := h.PerbandinganMopaRepository.GetPerbandinganMopaByKode(kodePerbandinganMopa)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deletePerbandinganMopa, err := h.PerbandinganMopaRepository.DeletePerbandinganMopa(perbandingan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deletePerbandinganMopa}
	json.NewEncoder(w).Encode(response)
}
