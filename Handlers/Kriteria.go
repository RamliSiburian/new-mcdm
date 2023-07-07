package handlers

import (
	"encoding/json"
	dto "mcdm/Dto"
	models "mcdm/Models"
	repositories "mcdm/Repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerkriteria struct {
	KriteriaRepository repositories.KriteriaRepository
}

func HandlerKriteria(KriteriaRepository repositories.KriteriaRepository) *handlerkriteria {
	return &handlerkriteria{KriteriaRepository}
}

func (h *handlerkriteria) FindKriteria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kriteria, err := h.KriteriaRepository.FindKriteria()

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

func (h *handlerkriteria) InsertKriteria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.KriteriaRequest)
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

	lastcode, err := h.KriteriaRepository.GetLastKode()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	newKode := "C" + strconv.Itoa(lastcode[0])

	createKriteria := models.Kriteria{
		Kode:         newKode,
		NamaKriteria: request.NamaKriteria,
		Bobot:        request.Bobot,
		Kategory:     request.Kategori,
		CreatedAt:    time.Now(),
	}

	data, err := h.KriteriaRepository.CreateKriteria(createKriteria)
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

func (h *handlerkriteria) GetKriteriaByKode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kode, _ := mux.Vars(r)["kode"]

	var kriteria models.Kriteria

	kriteria, err := h.KriteriaRepository.GetKriteriaByKode(kode)
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

func (h *handlerkriteria) UpdateKriteria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.KriteriaUpdateRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kodekriteria := mux.Vars(r)["kode"]
	oldData, err := h.KriteriaRepository.GetKriteriaByKode(kodekriteria)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.NamaKriteria != "" {
		oldData.NamaKriteria = request.NamaKriteria
	}
	if request.Bobot > 0 {
		oldData.Bobot = request.Bobot
	}
	if request.Kategori != "" {
		oldData.Kategory = request.Kategori
	}

	data, err := h.KriteriaRepository.UpdateKriteria(oldData)
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

func (h *handlerkriteria) DeleteKriteria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kodeKriteria := mux.Vars(r)["kode"]
	kriteria, err := h.KriteriaRepository.GetKriteriaByKode(kodeKriteria)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteKriteria, err := h.KriteriaRepository.DeleteKriteria(kriteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deleteKriteria}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerkriteria) LastCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lastcode, err := h.KriteriaRepository.GetLastKode()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	newKode := "C" + strconv.Itoa(lastcode[0])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: newKode}
	json.NewEncoder(w).Encode(response)
}
