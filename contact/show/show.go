package show

func Do() (string, error) {
	return plainTextContactDescription(), nil
}

func plainTextContactDescription() string {
	return `Hi! I am Marc Bouvier, aspiring Software Craftsman.  
I am always learning to improve my craft and trying new ways to add value to products.  
  
To reach me:  
- e-mail:     mailto:m.bouvier.dev@gmail.com  
- phone:      (+33) 6 66 15 95 38  
- LinkedIn:   https://www.linkedin.com/in/profileofmarcbouvier/  
- Web Site:   https://baldir.fr/`
}
