package request

import "goodjobs/business/complexes"

type ComplexRequest struct {
	Name string `json:"name"`
	Address	  string `json:"address"`
	Img  string `json:"img"`
}

func (Complex *ComplexRequest) ToDomain() *complexes.Domain {
	return &complexes.Domain{
		Name:    Complex.Name,
		Address: Complex.Address,
		Img: 	 Complex.Img,	
	}
}