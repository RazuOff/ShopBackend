package postgre

import "ginexample.com/pkg/models"

var testProducts = []models.Product{
	{
		ManufacturerId:    1,
		ProductCategoryId: 1,
		Name:              "Intel Core i7-12700K",
		Price:             399.99,
		Quantity:          50,
		ArticleNumber:     1001,
		Description:       "Процессор Intel 12-го поколения с 12 ядрами и 20 потоками",
		ImageUrl:          "https://example.com/intel-i7.jpg",
	},
	{

		ManufacturerId:    2,
		ProductCategoryId: 2,
		Name:              "NVIDIA GeForce RTX 4070",
		Price:             599.99,
		Quantity:          30,
		ArticleNumber:     1002,
		Description:       "Видеокарта с поддержкой Ray Tracing и DLSS 3.0",
		ImageUrl:          "https://example.com/rtx4070.jpg",
	},
	{ManufacturerId: 1, ProductCategoryId: 1, Name: "Processor", Price: 250.00, Quantity: 50, ArticleNumber: 1001, Description: "Intel i9 Processor", ImageUrl: "https://example.com/processor.jpg"},
	{ManufacturerId: 1, ProductCategoryId: 2, Name: "Motherboard", Price: 150.00, Quantity: 30, ArticleNumber: 1002, Description: "ASUS Z590 Motherboard", ImageUrl: "https://example.com/motherboard.jpg"},
	{ManufacturerId: 2, ProductCategoryId: 3, Name: "Graphics Card", Price: 500.00, Quantity: 20, ArticleNumber: 1003, Description: "NVIDIA RTX 3080", ImageUrl: "https://example.com/graphics_card.jpg"},
	{ManufacturerId: 3, ProductCategoryId: 4, Name: "RAM", Price: 80.00, Quantity: 100, ArticleNumber: 1004, Description: "Corsair Vengeance 16GB DDR4", ImageUrl: "https://example.com/ram.jpg"},
	{ManufacturerId: 4, ProductCategoryId: 5, Name: "SSD", Price: 120.00, Quantity: 40, ArticleNumber: 1005, Description: "Samsung 970 EVO 1TB SSD", ImageUrl: "https://example.com/ssd.jpg"},
	{ManufacturerId: 5, ProductCategoryId: 6, Name: "Power Supply", Price: 90.00, Quantity: 60, ArticleNumber: 1006, Description: "EVGA 750W Gold Power Supply", ImageUrl: "https://example.com/power_supply.jpg"},
	{ManufacturerId: 6, ProductCategoryId: 7, Name: "CPU Cooler", Price: 50.00, Quantity: 80, ArticleNumber: 1007, Description: "Cooler Master Hyper 212", ImageUrl: "https://example.com/cpu_cooler.jpg"},
	{ManufacturerId: 7, ProductCategoryId: 8, Name: "Case", Price: 100.00, Quantity: 90, ArticleNumber: 1008, Description: "NZXT H510 Mid Tower Case", ImageUrl: "https://example.com/case.jpg"},
	{ManufacturerId: 8, ProductCategoryId: 9, Name: "Monitor", Price: 200.00, Quantity: 40, ArticleNumber: 1009, Description: "LG 27-inch 144Hz Monitor", ImageUrl: "https://example.com/monitor.jpg"},
	{ManufacturerId: 9, ProductCategoryId: 10, Name: "Keyboard", Price: 60.00, Quantity: 150, ArticleNumber: 1010, Description: "Logitech Mechanical Keyboard", ImageUrl: "https://example.com/keyboard.jpg"},
}

var testUsers = []models.User{
	{
		Username: "Admin",
		Password: "1234",
		Role:     "Admin",
	},
	{
		Username: "user",
		Password: "1234",
	},
}

// Пример корзин
var carts = []models.Cart{
	{UserID: 1}, // Корзина для пользователя john_doe
	{UserID: 2}, // Корзина для пользователя jane_doe
}
