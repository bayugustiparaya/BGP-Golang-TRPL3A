package common

//Struct API
// Order struct (Model) ...

type Message  struct {
	Code      int  	 `json:"code"`
	Remark    string `json:"remark"`
	OrderID   string `json:"orderID"`	
	Orders   *Orders `json:"orders,omitempty"`
	Result   *Result  `json:"result,omitempty"`
}

type Orders struct {
	OrderID      string        `json:"orderID"`
	CustomerID   string        `json:"customerID"`
	EmployeeID   string        `json:"employeeID"`
	OrderDate    string        `json:"orderDate"`
	OrdersDet []OrdersDetail   `json:"ordersDetail"`
	
}

type OrdersDetail struct {
	OrderID      string  `json:"orderID"`
	ProductID  	 string  `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	UnitPrice    float64 `json:"UnitPrice"`
	Quantity     int     `json:"Quantity"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

// add strcut customer n product (menambahkan struct customer dan product)
type Customer struct {
	CustomerID   string `json:"customerID"`
	CompanyName  string `json:"companyName"`
	ContactName  string `json:"contactName"`
	ContactTitle string `json:"contactTitle"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Region			 string `json:"region"`
	PostalCode   string `json:"postalCode"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Fax					 string `json:"fax"`
}

type Product struct {
	ProductID				int			`json:"productID"`
	ProductName			string	`json:"productName"`
	SupplierID			int			`json:"supplierID"`
	CategoryID			int			`json:"categoryID"`
	QuantityPerUnit	string	`json:"quantityPerUnit"`
	UnitPrice				float32	`json:"unitPrice"`
	UnitInStock			int			`json:"unitInStock"`
	UnitOnOrder			int			`json:"unitOnOrder"`
	ReorderLevel		int			`json:"reorderLevel"`
	Discontinued		int8		`json:"discontinued"`
	Description			string	`json:"description"`
}

//End Struct API
