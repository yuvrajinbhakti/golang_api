package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/yuvrajinbhakti/golang_api/api"
	"github.com/yuvrajinbhakti/golang_api/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request){
	var params = api.CoinBalanceParams{
		Username: r.URL.Query().Get("username"),
	}
	
	if params.Username == "" {
		log.Error("Missing username parameter")
		api.InternalErrorHandler(w)
		return 
	}

	var database *tools.DatabaseInterface 
	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return 
	}
	
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil{
		log.Error("No coin details found for user")
		api.InternalErrorHandler(w)
		return 
	}

	var response = api.CoinBalanceResponse { 
		Balance:  (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response) 
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return 
	}
}