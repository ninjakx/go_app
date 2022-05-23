package handler

import (
	"encoding/json"
	"net/http"
	"GO_APP/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"strconv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// getUserOr404 gets a User instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	// var user User
	if err := db.First(&user, model.User{Model: gorm.Model{ID: id}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &user
}

// getUserAddressOr404 gets a User Address instance if exists, or respond the 404 error otherwise
func getUserAddressOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) [] model.UserAddress {
	userAddress := []model.UserAddress{}
	if err := db.Find(&userAddress, model.UserAddress{UserId: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return userAddress
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	err := db.Create(&user).Error
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

func CreateUserAddress(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["id"])
    id, err := strconv.ParseUint(vars["id"], 10, 32)
    if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
    }
    wd := uint(id)

	user := getUserOr404(db, wd, w, r)
	if user == nil {
		return
	}
    // wd := uint(id)
	userAddress := model.UserAddress{UserId: wd}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userAddress); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	err = db.Create(&userAddress).Error
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, userAddress)
}

func GetAllUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}


func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["id"])

    id, err := strconv.ParseUint(vars["id"], 10, 32)
    if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
    }
    wd := uint(id)

	user := getUserOr404(db, wd, w, r)
	if user == nil {
		return
	}
	addresses := []model.UserAddress{}
	err = db.Find(&addresses, "user_id = ?", id).Error
	user.Addrs = addresses

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}
		
	respondJSON(w, http.StatusOK, user)
}

func GetUserAddress(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["id"])

    id, err := strconv.ParseUint(vars["id"], 10, 32)
    if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
    }
    wd := uint(id)

	userAddress := getUserAddressOr404(db, wd, w, r)
	if userAddress == nil {
		return
	}
		
	respondJSON(w, http.StatusOK, userAddress)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 32)
    if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
    }
    wd := uint(id)

	user := getUserOr404(db, wd, w, r)
	if user == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}


func UpdateUserAddress(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 32)
	addr_id, err := strconv.ParseUint(vars["addr_id"], 10, 32)
    wd := uint(id)
    wd_addr := uint(addr_id)
    if err != nil {
        // fmt.Println(err)
		respondError(w, http.StatusInternalServerError, err.Error())
    }

	userAddress := model.UserAddress{}
	if err := db.Find(&userAddress, model.UserAddress{UserId: wd, Model: gorm.Model{ID: wd_addr}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return 
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userAddress); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&userAddress).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, userAddress)
}

func DisableUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id, _ := strconv.ParseUint(vars["id"], 10, 32)
    wd := uint(id)
	user := getUserOr404(db, wd, w, r)
	if user == nil {
		return
	}
	user.Disable()
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func EnableUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

    id, _ := strconv.ParseUint(vars["id"], 10, 32)
    wd := uint(id)
	
	user := getUserOr404(db, wd, w, r)
	if user == nil {
		return
	}
	user.Enable()
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

    id, _ := strconv.ParseUint(vars["id"], 10, 32)
    wd := uint(id)

	user := getUserOr404(db, wd, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DeleteUserAddress(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

    id, _ := strconv.ParseUint(vars["id"], 10, 32)
	addr_id, _ := strconv.ParseUint(vars["addr_id"], 10, 32)
    wd := uint(id)
    wd_addr := uint(addr_id)

	userAddress := model.UserAddress{}
	// var userAddress UserAddress
	if err := db.Find(&userAddress, model.UserAddress{UserId: wd, Model: gorm.Model{ID: wd_addr}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return 
	}

	if err := db.Delete(&userAddress).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// // // {
// // // 	"Id":  1,
// // // 	"UserId": 240,
// // // 	"AddressLine1": "abcdf",
// // // 	"AddressLine2": "lucknow",
// // // 	"City": "kanpur",
// // // 	"PostalCode": "226021",
// // // 	"Country": "India",
// // // 	"Phone": "9818476950",
// // // 	"Telephone": "783232"
// // // }

// // {
// //     "Id":2140,
// //     "Username":"Kriti",
// //     "Password":"",
// //     "ModifiedAt":"",
// //     "Addrs":null
// // }



