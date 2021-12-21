package request

import "goodjobs/business/complexes"

type ComplexRequest struct {
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (Complex *ComplexRequest) ToDomain() *complexes.Domain {
	return &complexes.Domain{
		Name:    Complex.Name,
		Img: 	 Complex.Img,	
	}
}