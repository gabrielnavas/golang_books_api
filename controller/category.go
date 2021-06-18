package controller

import (
	"books_api/entity"
	servicecategory "books_api/service"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request)
}

var (
	ErrNoDataBody = errors.New("EOF")
)

type categoryControllerImpl struct {
	categoryService servicecategory.InsertOneCategory
}

func NewCategoryRepository(categoryService servicecategory.InsertOneCategory) *categoryControllerImpl {
	return &categoryControllerImpl{categoryService}
}

func (c *categoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	category := entity.Category{}

	_, err := getIP(r)
	if err != nil {
		fmt.Print(err)
	}

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil && err == io.EOF {
		err = errors.New("no data on body found")
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	category_created, err := c.categoryService.InsertOne(category)
	if err != nil {
		if err == servicecategory.ErrRepositoryError {
			responseJsonErr(&w, err, http.StatusInternalServerError)
			return
		}
		responseJsonErr(&w, err, http.StatusBadRequest)
		return
	}

	responseCreated(&w, category_created)
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		fmt.Printf("X-REAL-IP: %v\n", ip)
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			fmt.Printf("X-REAL-IP: %v\n", ip)
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		fmt.Printf("X-REAL-IP: %v\n", ip)
		return ip, nil
	}
	return "", fmt.Errorf("no valid ip found")
}
