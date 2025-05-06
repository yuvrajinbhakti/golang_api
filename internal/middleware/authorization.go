package middleware 
import {
	"errors"
	"net/http" 

	"github.com/yuvrajinbhakti/golang_api/internal/api"
	"github.com/yuvrajinbhakti/golang_api/internal/tools"
	log "github.com/sirupsen/logrus"
}

var unAuthorized = errors.New("Invalid username or token");