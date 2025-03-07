package models

type Livro struct{

	ID 	   			 uint 	  `gorm:"primaryKey"`
	Nome 			 string   `json:"nome"`
	Edicao 			 uint 	  `json:"edicao"`
	Ano 		     uint 	  `json:"ano"`
	Autor  		 	 string   `json:"autor"` 

}



