package models

import (
	"database/sql"

	"github.com/Aryan-Gupta4460/Product/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Search(keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) SearchPrices(min, max float64) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where price  >= ? and price <= ?", min, max)
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Create(product *entities.Product) (err error) {
	result, err := productModel.Db.Exec("insert into product(name,price,quantity) values(?,?,?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}

}

func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update product set name = ?,price=?,quantity =? where id =? ", product.Name, product.Price, product.Quantity, product.Id)
	if err != nil {
		return 0, err
	} else {

		return result.RowsAffected()
	}

}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("delete from product  where id =? ", id)
	if err != nil {
		return 0, err
	} else {

		return result.RowsAffected()
	}

}

//func respondWithError(w http.ResponseWriter, code int, msg string) {
//	respondWithJson(w, code, map[string]string{"error": msg})
//}

//func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
//	response, _ := json.Marshal(payload)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(code)
//	w.Write(response)
//}
