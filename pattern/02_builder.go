package pattern

// Паттерн Builder
/*
Плюсы:
* Позволяет создавать продукты пошагово.
* Позволяет использовать один и тот же код для создания различных продуктов.
*Изолирует сложный код сборки продукта от его основной бизнес-логики.

Минусы:
* Усложняет код программы из-за введения дополнительных классов.
* Клиент будет привязан к конкретным классам строителей,
  так как в интерфейсе директора может не быть метода получения результата.
*/

// Product - основной класс
type Product struct {
	// Эти два поля могут иметь различные значение
	name string
	cost int8
	// Это поле имеет только 2 значения: пустая строка для локально
	//расположенного продукта и some address для продукта на складе
	delivery string
}

// IProductBuilder - интерфейс продуцирующих структур
type IProductBuilder interface {
	GetProduct() Product
	SetName(name string)
	SetCost(cost int8)
	SetDelivery()
}

// StorageProductBuilder - билдер для продуктов в хранилище
type StorageProductBuilder struct {
	product Product
}

func (p *StorageProductBuilder) GetProduct() Product {
	return p.product
}

func (p *StorageProductBuilder) SetName(name string) {
	p.product.name = name
}

func (p *StorageProductBuilder) SetCost(cost int8) {
	p.product.cost = cost
}

func (p *StorageProductBuilder) SetDelivery() {
	p.product.delivery = "some address"
}

// LocalProductBuilder - билдер для продуктов на локальном складе магазина
type LocalProductBuilder struct {
	product Product
}

func (p *LocalProductBuilder) GetProduct() Product {
	return p.product
}

func (p *LocalProductBuilder) SetName(name string) {
	p.product.name = name
}

func (p *LocalProductBuilder) SetCost(cost int8) {
	p.product.cost = cost
}

func (p *LocalProductBuilder) SetDelivery() {
	p.product.delivery = ""
}

// ProductDirector управляет сборкой объекта продукта в зависимости от переданного билдера
type ProductDirector struct {
	builder IProductBuilder
}

func (p *ProductDirector) SetBuilder(builder IProductBuilder) {
	p.builder = builder
}

func (p *ProductDirector) BuildProduct(name string, cost int8) Product {
	p.builder.SetCost(cost)
	p.builder.SetName(name)
	p.builder.SetDelivery()
	return p.builder.GetProduct()
}
