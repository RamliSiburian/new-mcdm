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

type handleralternatif struct {
	AlternatifRepository repositories.AlternatifRepository
}

func HandlerAlternatif(AlternatifRepository repositories.AlternatifRepository) *handleralternatif {
	return &handleralternatif{AlternatifRepository}
}

func (h *handleralternatif) FindAlternatif(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	alternatif, err := h.AlternatifRepository.FindAlternatif()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: alternatif}
	json.NewEncoder(w).Encode(response)
}

func (h *handleralternatif) InsertAlternatif(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.AlternatifRequest)
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

	// lastcode, err := h.AlternatifRepository.GetLastKodeAlternatif()
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// newKode := "A" + strconv.Itoa(lastcode[0])

	createAlternatif := models.Alternatif{
		Kode:           request.Kode,
		KodeKriteria:   request.KodeKriteria,
		NamaAlternatif: request.NamaAlternatif,
		Nilai:          request.Nilai,
		CreatedAt:      time.Now(),
	}

	data, err := h.AlternatifRepository.CreateAlternatif(createAlternatif)
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

func (h *handleralternatif) GetAlternatifByKode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kode, _ := mux.Vars(r)["kode"]

	var alternatif models.Alternatif

	alternatif, err := h.AlternatifRepository.GetAlternatifByKode(kode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: alternatif}
	json.NewEncoder(w).Encode(response)
}

func (h *handleralternatif) GetAlternatifByKodeAndKodeKriteria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kode, _ := mux.Vars(r)["kode"]
	kodeKriteria, _ := mux.Vars(r)["kodeKriteria"]

	var alternatif models.Alternatif

	alternatif, err := h.AlternatifRepository.GetAlternatifByKodeAndKodeKriteria(kode, kodeKriteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: alternatif}
	json.NewEncoder(w).Encode(response)
}

func (h *handleralternatif) UpdateAlternatif(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dto.AlternatifRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kodealternatif := mux.Vars(r)["kode"]
	kodeKriteria := mux.Vars(r)["kodeKriteria"]

	oldData, err := h.AlternatifRepository.GetAlternatifByKodeAndKodeKriteria(kodealternatif, kodeKriteria)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.NamaAlternatif != "" {
		oldData.NamaAlternatif = request.NamaAlternatif
	}

	if request.Nilai != "" {
		oldData.Nilai = request.Nilai
	}

	data, err := h.AlternatifRepository.UpdateAlternatif(oldData)
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

func (h *handleralternatif) DeleteAlternatif(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kodeAlternatif := mux.Vars(r)["kode"]
	// alternatif, err := h.AlternatifRepository.GetAlternatifByKode(kodeAlternatif)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	deleteAlternatif, err := h.AlternatifRepository.DeleteAlternatif(kodeAlternatif)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deleteAlternatif}
	json.NewEncoder(w).Encode(response)
}

func (h *handleralternatif) LastCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lastcode, err := h.AlternatifRepository.GetLastKodeAlternatif()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	newKode := "A" + strconv.Itoa(lastcode[0])

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
