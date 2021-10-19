package service

import (
	"tubespbbo/domains/product/dto"
	"tubespbbo/domains/product/model"
	"tubespbbo/domains/product/repository"
)

func FindProduct() (*[]model.Product, error) {
	return repository.FindProduct()
}

func FindOneProduct(id int64) (*model.Product, error) {
	return repository.FindOneProduct(id)
}

func CreateProduct(products *model.Product) error {
	return repository.CreateProduct(products)
}

func UpdateProduct(updateDto *dto.UpdateProductDTO, id int64) (*model.Product, error) {
	products, err := repository.FindOneProduct(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Nama != "" {
		products.Nama = updateDto.Nama
	}
	if updateDto.Kode != "" {
		products.Kode = updateDto.Kode
	}
	if updateDto.Description != "" {
		products.Description = updateDto.Description
	}
	if updateDto.Price != "" {
		products.Price = updateDto.Price
	}
	if updateDto.Quantity != "" {
		products.Quantity = updateDto.Quantity
	}
	if updateDto.ImageUrl != "" {
		products.ImageUrl = updateDto.ImageUrl
	}
	products, err = repository.UpdateProduct(products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func DeleteProduct(id int64) (*model.Product, error) {
	products, err := repository.FindOneProduct(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteProduct(products)
}
