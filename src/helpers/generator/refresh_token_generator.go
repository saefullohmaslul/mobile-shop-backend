package generator

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
	"github.com/saefullohmaslul/mobile-shop-backend/src/db/entity"
	"github.com/saefullohmaslul/mobile-shop-backend/src/repositories"
	uuid "github.com/satori/go.uuid"
)

// StoreRefreshToken is method to create refresh token in database
func StoreRefreshToken(
	userID uuid.UUID,
	authInformationRepository repositories.AuthInformationRepository,
	refreshToken string,
	errors []response.Error,
) {
	authInformation, err := authInformationRepository.GetUserID(entity.AuthInformation{
		UserID: userID,
	})

	if gorm.IsRecordNotFoundError(err) {
		createRefreshToken(userID, authInformationRepository, refreshToken, errors)
		return
	}

	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "AUTHINFO_GET_DB_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}

	if (authInformation == entity.AuthInformation{}) {
		createRefreshToken(userID, authInformationRepository, refreshToken, errors)
		return
	}

	if err := authInformationRepository.UpdateRefreshToken(
		authInformation,
		entity.AuthInformation{
			RefreshToken: refreshToken,
		}); err != nil {
		errors = append(errors, response.Error{
			Flag:    "REFRESHTOKEN_UPDATE_DB_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}
}

func createRefreshToken(
	userID uuid.UUID,
	authInformationRepository repositories.AuthInformationRepository,
	refreshToken string,
	errors []response.Error,
) {
	err := authInformationRepository.CreateRefreshToken(entity.AuthInformation{
		RefreshToken: refreshToken,
		UserID:       userID,
	})
	if err != nil {
		errors = append(errors, response.Error{
			Flag:    "REFRESHTOKEN_CREATE_DB_ERROR",
			Message: err.Error(),
		})
		response.InternalServerError("Internal server error", errors)
	}
}
