package types

//Мин-е Денежные еденицы(cent, diram)
type Money int64

//категория в которой сделана покупка (авто, магазин, ресторан)
type Category string

//Информация о платеже
type Payment struct{
	ID int
	Amount Money
	Category Category
}



