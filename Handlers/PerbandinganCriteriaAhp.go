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

type handlerperbandingancriteriaahp struct {
	PerbandinganCriteriaAhpRepository repositories.PerbandinganCriteriaAhpRepository
}

func Handlerperbandingancriteriaahp(PerbandinganCriteriaAhpRepository repositories.PerbandinganCriteriaAhpRepository) *handlerperbandingancriteriaahp {
	return &handlerperbandingancriteriaahp{PerbandinganCriteriaAhpRepository}
}

func (h *handlerperbandingancriteriaahp) FindPerbandinganCriteriaAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kriteria, err := h.PerbandinganCriteriaAhpRepository.FindPerbandinganCriteriaAhp()

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

func (h *handlerperbandingancriteriaahp) InsertPerbandinganCriteriaAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.PerbandinganCriteriaAhpRequest)
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

	createPerbandinganCriteriaAhp := models.PerbandinganCriteriaAhp{
		Kode:      request.Kode,
		Deskripsi: request.Deskripsi,
		Nilai:     request.Nilai,
		CreatedAt: time.Now(),
	}

	data, err := h.PerbandinganCriteriaAhpRepository.CreatePerbandinganCriteriaAhp(createPerbandinganCriteriaAhp)
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

func (h *handlerperbandingancriteriaahp) GetPerbandinganCriteriaAhpByKode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kode, _ := mux.Vars(r)["kode"]

	var perbandinganCriteriaAhp models.PerbandinganCriteriaAhp

	perbandinganCriteriaAhp, err := h.PerbandinganCriteriaAhpRepository.GetPerbandinganCriteriaAhpByKode(kode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: perbandinganCriteriaAhp}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerperbandingancriteriaahp) UpdatePerbandinganCriteriaAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.PerbandinganCriteriaAhpRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kodePerbandinganCriteriaAhp := mux.Vars(r)["kode"]
	oldData, err := h.PerbandinganCriteriaAhpRepository.GetPerbandinganCriteriaAhpByKode(kodePerbandinganCriteriaAhp)
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
	if request.Nilai >= 0 {
		oldData.Nilai = request.Nilai
	}

	data, err := h.PerbandinganCriteriaAhpRepository.UpdatePerbandinganCriteriaAhp(oldData)
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

func (h *handlerperbandingancriteriaahp) DeletePerbandinganCriteriaAhp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kodePerbandinganCriteriaAhp := mux.Vars(r)["kode"]
	perbandingan, err := h.PerbandinganCriteriaAhpRepository.GetPerbandinganCriteriaAhpByKode(kodePerbandinganCriteriaAhp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deletePerbandinganCriteriaAhp, err := h.PerbandinganCriteriaAhpRepository.DeletePerbandinganCriteriaAhp(perbandingan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deletePerbandinganCriteriaAhp}
	json.NewEncoder(w).Encode(response)
}
