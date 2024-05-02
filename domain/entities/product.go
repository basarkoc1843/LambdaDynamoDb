package entities

type Product struct {
	ProductUUID string `dynamodbav:"ProductUUID"`
	Name        string `dynamodbav:"Name"`
	Description string `dynamodbav:"description"`
	Brand       string `dynamodbav:"brand"`
	Price       string `dynamodbav:"price"`
	Category    string `dynamodbav:"category"`
	DateAdded   string `dynamodbav:"dateAdded"`
}

type Personal struct {
	Name    string
	Surname string
}
