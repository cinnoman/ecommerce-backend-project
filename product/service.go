package product

import "ecommerce/domain"

type service struct {
	prdctRepo ProductRepo
}

func NewService(prdctRepo ProductRepo) Service {
	return &service{
		prdctRepo: prdctRepo,
	}
}

func (svc *service) Create(newProduct domain.Product) (*domain.Product, error) {
	//return svc.prdctRepo.Create(newProduct)
	prdct, err := svc.prdctRepo.Create(newProduct)
	if err != nil {
		return nil, err
	}

	if prdct == nil {
		return nil, nil
	}

	return prdct, nil
}

func (svc *service) Get(pId int) (*domain.Product, error) {
	prdct, err := svc.prdctRepo.Get(pId)
	if err != nil {
		return nil, err
	}

	if prdct == nil {
		return nil, nil
	}

	return prdct, nil
}


func (svc *service) List() ([]*domain.Product, error) {
	prdct, err := svc.prdctRepo.List()
	if err != nil {
		return nil, err
	}

	if prdct == nil {
		return nil, nil
	}

	return prdct, nil
}

func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	prdct, err := svc.prdctRepo.Update(product)
	if err != nil {
		return nil, err
	}

	if prdct == nil {
		return nil, nil
	}

	return prdct, nil
}

func (svc *service) Delete(pId int) error {
	err := svc.Delete(pId)
	if err != nil {
		return err
	}

	return nil
}
