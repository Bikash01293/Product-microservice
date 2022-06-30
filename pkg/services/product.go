package services

import (
	"context"
	"net/http"
	"product-micro/pkg/db"
	"product-micro/pkg/models"
	"product-micro/pkg/pb"
)

type Server struct {
	H db.Handler
}
//Create the product when the request is accepted.
func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	product.Title = req.Title
	product.Desc = req.Desc
	product.Img = req.Img
	product.Categories = req.Categories
	product.Size = req.Size
	product.Color = req.Color
	product.Price = req.Price
	product.Stock = req.Stock
	// fmt.Println("this is a product:", product)
	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}
//Get the product by the product id.
func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product
	// fmt.Println("product id request to find product : ", req.Id)
	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Title:  product.Title,
		Desc: product.Desc,
		Img: product.Img,
		Categories: product.Categories,
		Size: product.Size,
		Color: product.Color,
		Price: product.Price,
		Stock: product.Stock,

	}


	// fmt.Println("This is the data to send for findOne response: ", data)
	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}
//Decrease the stock by the product id when the order is placed.
func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1

	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}

//Get all the products.
func (s *Server) FindAllProduct(ctx context.Context, req *pb.FindAllProductRequest) (*pb.FindAllProductResponse, error) {
	var product []models.Product
	if result := s.H.DB.Find(&product); result.Error != nil {
		return &pb.FindAllProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	data := &pb.FindOneData{}
	data2 := &pb.FindAllProductResponse{Data: []*pb.FindOneData{}}
	data3 := data2.Data
	for i := 0; i < len(product); i++ {
		data = &pb.FindOneData{
			Id:    product[i].Id,
			Title:  product[i].Title,
			Desc: product[i].Desc,
			Img: product[i].Img,
			Categories: product[i].Categories,
			Size: product[i].Size,
			Color: product[i].Color,
			Price: product[i].Price,
			Stock: product[i].Stock,
		}
		data3 = append(data3, data)
	}
	

	return &pb.FindAllProductResponse{
		Status: http.StatusOK,
		Data:   data3,
	}, nil
}


//Update the product by the product id.
func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	var prod models.Product
	var productreq models.Product
	// fmt.Println("product id request to find product : ", req.Id)
	// fmt.Println("the request is:", req.Color)
	productreq.Title = req.Title
	productreq.Desc = req.Desc
	productreq.Img = req.Img
	productreq.Categories = req.Categories
	productreq.Size = req.Size
	productreq.Color = req.Color
	productreq.Price = req.Price
	productreq.Stock = req.Stock
	if result := s.H.DB.Model(prod).Where(req.Id).Updates(productreq); result.Error != nil {
		return &pb.UpdateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	if findResult := s.H.DB.First(&prod, req.Id); findResult.Error != nil {
		return &pb.UpdateProductResponse{
			Status: http.StatusConflict,
			Error:  findResult.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    prod.Id,
		Title:  prod.Title,
		Desc: prod.Desc,
		Img: prod.Img,
		Categories: prod.Categories,
		Size: prod.Size,
		Color: prod.Color,
		Price: prod.Price,
		Stock: prod.Stock,
	}

	// fmt.Println("This is the data to send for update product response: ", prod)
	return &pb.UpdateProductResponse{
		Status: http.StatusOK,
		Data: data,
	}, nil
}

//Delete the product by the product id.
func (s *Server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	var product models.Product
	// fmt.Println("product id request to find product : ", req.Id)
	if findResult := s.H.DB.First(&product, req.Id); findResult.Error != nil {
		return &pb.DeleteProductResponse{
			Status: http.StatusConflict,
			Error:  findResult.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    product.Id,
		Title:  product.Title,
		Desc: product.Desc,
		Img: product.Img,
		Categories: product.Categories,
		Size: product.Size,
		Color: product.Color,
		Price: product.Price,
		Stock: product.Stock,
	}

	if result := s.H.DB.Model(product).Where(req.Id).Delete(product); result.Error != nil {
		return &pb.DeleteProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	// fmt.Println("This is the data to send for delete product response: ", prod)
	return &pb.DeleteProductResponse{
		Status: http.StatusOK,
		Data: data,
	}, nil
}