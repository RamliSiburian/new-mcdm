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

type handlerperbandinganahp struct {
	PerbandinganAhpRepository repositories.PerbandinganAhpRepository
}

func Handlerperbandinganahp(PerbandinganAhpRepository repositories.PerbandinganAhpRepository) *handlerperbandinganahp {
	return &handlerperbandinganahp{PerbandinganAhpRepository}
}

func (h *handlerperbandinganahp) FindPerbandinganAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kriteria, err := h.PerbandinganAhpRepository.FindPerbandinganAhp()

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

func (h *handlerperbandinganahp) InsertPerbandinganAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.PerbandinganAhpRequest)
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

	createPerbandinganAhp := models.PerbandinganAhp{
		Kode:      request.Kode,
		Deskripsi: request.Deskripsi,
		Nilai:     request.Nilai,
		CreatedAt: time.Now(),
	}

	data, err := h.PerbandinganAhpRepository.CreatePerbandinganAhp(createPerbandinganAhp)
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

func (h *handlerperbandinganahp) GetPerbandinganAhpByKode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kode, _ := mux.Vars(r)["kode"]

	var perbandinganAhp models.PerbandinganAhp

	perbandinganAhp, err := h.PerbandinganAhpRepository.GetPerbandinganAhpByKode(kode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: perbandinganAhp}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerperbandinganahp) UpdatePerbandinganAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.PerbandinganAhpRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kodePerbandinganAhp := mux.Vars(r)["kode"]
	oldData, err := h.PerbandinganAhpRepository.GetPerbandinganAhpByKode(kodePerbandinganAhp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Kode > 0 {
		oldData.Kode = request.Kode
	}

	if request.Deskripsi != "" {
		oldData.Deskripsi = request.Deskripsi
	}
	if request.Nilai > 0 {
		oldData.Nilai = request.Nilai
	}

	data, err := h.PerbandinganAhpRepository.UpdatePerbandinganAhp(oldData)
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

func (h *handlerperbandinganahp) DeletePerbandinganAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kodePerbandinganAhp := mux.Vars(r)["kode"]
	perbandingan, err := h.PerbandinganAhpRepository.GetPerbandinganAhpByKode(kodePerbandinganAhp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deletePerbandinganAhp, err := h.PerbandinganAhpRepository.DeletePerbandinganAhp(perbandingan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deletePerbandinganAhp}
	json.NewEncoder(w).Encode(response)
}
