//vamos a construir los apis a partir de las necesidades que tenemos para trabajar con los sistemas
//a estos api tambien se los conoce como puntos finales

package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// *******************************************************************PERSONA*************************************************************************
// necesitamos pasarla a formato json para que pueda ser usada por el front
type Persona struct {
	Id              int          `json:"id"`
	CedulaRuc       string       `json:"cedula"`
	FechaNacimiento time.Time    `json:"fechaNacimiento"`
	Apellidos       string       `json:"apellidos"`
	Nombres         string       `json:"nombres"`
	Telefono        string       `json:"telefono"`
	Email           string       `json:"email"`
	Direcciones     []*Direccion `json:"direcciones"`
}

// se agregara una persona con la fecha de nacimiento con el formato correcto, para eso se inicializara correctamente la variable
var fechaNacimiento1, _ = time.Parse(time.RFC3339, "1990-01-01T00:00:00Z")
var fechaNacimiento2, _ = time.Parse(time.RFC3339, "1985-05-12T00:00:00Z")
var fechaNacimiento3, _ = time.Parse(time.RFC3339, "1993-07-23T00:00:00Z")
var fechaNacimiento4, _ = time.Parse(time.RFC3339, "2000-11-30T00:00:00Z")

// vamos a crear una lista de personas para crear en el api
var personas = []Persona{
	{Id: 1, CedulaRuc: "0150041580", FechaNacimiento: fechaNacimiento1, Apellidos: "PACHECO HERRERA", Nombres: "GALILEA KATHERINE",
		Telefono: "0996972323", Email: "galypachecoh@gmail.com",
		Direcciones: []*Direccion{
			{
				Id:           1,
				Pais:         "Ecuador",
				Provincia:    "Pichincha",
				CiudadCanton: "Quito",
				Calle1:       "Av. Primera",
				Calle2:       "N/A",
				NumCasa:      "123",
				CodPostal:    "170501",
				Referencia:   "Frente al parque",
				PersonaId:    1,
			},
			{
				Id:           2,
				Pais:         "Ecuador",
				Provincia:    "Guayas",
				CiudadCanton: "Guayaquil",
				Calle1:       "Calle Principal",
				Calle2:       "N/A",
				NumCasa:      "456",
				CodPostal:    "090507",
				Referencia:   "Esquina del mercado",
				PersonaId:    1,
			},
		},
	},
	{Id: 1, CedulaRuc: "0123456789", FechaNacimiento: fechaNacimiento2, Apellidos: "PACHECO", Nombres: "SILVIO",
		Telefono: "0504060402", Email: "silvio@gmail.com", Direcciones: []*Direccion{
			{
				Id:           3,
				Pais:         "Ecuador",
				Provincia:    "Guayas",
				CiudadCanton: "Guayaquil",
				Calle1:       "Calle Principal",
				Calle2:       "N/A",
				NumCasa:      "456",
				CodPostal:    "090507",
				Referencia:   "Esquina del mercado",
				PersonaId:    2,
			},
		},
	},
	{Id: 1, CedulaRuc: "02468101214", FechaNacimiento: fechaNacimiento3, Apellidos: "HERRERA", Nombres: "KATERINE",
		Telefono: "0408090607", Email: "katy@gmail.com", Direcciones: []*Direccion{
			{
				Id:           4,
				Pais:         "Ecuador",
				Provincia:    "Azuay",
				CiudadCanton: "Cuenca",
				Calle1:       "Av. Principal",
				Calle2:       "N/A",
				NumCasa:      "789",
				CodPostal:    "010203",
				Referencia:   "Frente al parque",
				PersonaId:    3,
			},
			{
				Id:           5,
				Pais:         "Ecuador",
				Provincia:    "Manabí",
				CiudadCanton: "Manta",
				Calle1:       "Calle Secundaria",
				Calle2:       "N/A",
				NumCasa:      "1011",
				CodPostal:    "040506",
				Referencia:   "Al lado del mercado",
				PersonaId:    3,
			},
		},
	},
	{Id: 1, CedulaRuc: "0105090807", FechaNacimiento: fechaNacimiento4, Apellidos: "PACHECO", Nombres: "DAVID",
		Telefono: "0504080903", Email: "david@gmail.com", Direcciones: []*Direccion{
			{
				Id:           6,
				Pais:         "Ecuador",
				Provincia:    "Manabí",
				CiudadCanton: "Manta",
				Calle1:       "Calle Secundaria",
				Calle2:       "N/A",
				NumCasa:      "1011",
				CodPostal:    "040506",
				Referencia:   "Al lado del mercado",
				PersonaId:    4,
			},
		},
	},
}

// realizar metodos que permitan presentar la informacion de la api
// crear el metodo para controlar
func getPersonas(c *gin.Context) {
	//serializar los datos, tranformarlos a json
	c.IndentedJSON(http.StatusOK, personas)
}

// crear un metodo que utilice el protocolo post para guardar una persona
func postPersona(c *gin.Context) {
	var newPersona Persona
	if err := c.BindJSON(&newPersona); err != nil {
		//convertir el json en datos para almacenar
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personas = append(personas, newPersona)
	c.IndentedJSON(http.StatusCreated, newPersona)
}

// crear metodo par acceder a un obejto a traves del ID
func getPersonabyID(c *gin.Context) {
	//capturar primero el ID que viene en la ruta establecida en el api
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	for _, a := range personas {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "La persona no está en la lista"})
}

// Modificar una persona por ID
func putPersona(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedPersona Persona
	if err := c.BindJSON(&updatedPersona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range personas {
		if a.Id == intID {
			personas[i] = updatedPersona
			c.IndentedJSON(http.StatusOK, updatedPersona)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "La persona no está en la lista"})
}

// Eliminar una persona por ID
func deletePersona(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range personas {
		if a.Id == intID {
			personas = append(personas[:i], personas[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Persona eliminada"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "La persona no está en la lista"})
}

// *******************************************************************DIRECCION*************************************************************************
type Direccion struct {
	Id           int    `json: "id"`
	Pais         string `json: "pais"`
	Provincia    string `json: "provincia"`
	CiudadCanton string `json: "ciudadCanton"`
	Calle1       string `json: "calle1"`
	Calle2       string `json:"calle2"`
	NumCasa      string `json: "numCasa"`
	CodPostal    string `json: "codPostal"`
	Referencia   string `json: "referencia"`
	PersonaId    int    `json: "persona"`
}

var direcciones = []Direccion{
	{
		Id:           1,
		Pais:         "Ecuador",
		Provincia:    "Pichincha",
		CiudadCanton: "Quito",
		Calle1:       "Av. Primera",
		Calle2:       "N/A",
		NumCasa:      "123",
		CodPostal:    "170501",
		Referencia:   "Frente al parque",
		PersonaId:    1,
	},
	{
		Id:           2,
		Pais:         "Ecuador",
		Provincia:    "Guayas",
		CiudadCanton: "Guayaquil",
		Calle1:       "Calle Principal",
		Calle2:       "N/A",
		NumCasa:      "456",
		CodPostal:    "090507",
		Referencia:   "Esquina del mercado",
		PersonaId:    1,
	},
	{
		Id:           3,
		Pais:         "Ecuador",
		Provincia:    "Guayas",
		CiudadCanton: "Guayaquil",
		Calle1:       "Calle Principal",
		Calle2:       "N/A",
		NumCasa:      "456",
		CodPostal:    "090507",
		Referencia:   "Esquina del mercado",
		PersonaId:    2,
	},
	{
		Id:           4,
		Pais:         "Ecuador",
		Provincia:    "Azuay",
		CiudadCanton: "Cuenca",
		Calle1:       "Av. Principal",
		Calle2:       "N/A",
		NumCasa:      "789",
		CodPostal:    "010203",
		Referencia:   "Frente al parque",
		PersonaId:    3,
	},
	{
		Id:           5,
		Pais:         "Ecuador",
		Provincia:    "Manabí",
		CiudadCanton: "Manta",
		Calle1:       "Calle Secundaria",
		Calle2:       "N/A",
		NumCasa:      "1011",
		CodPostal:    "040506",
		Referencia:   "Al lado del mercado",
		PersonaId:    3,
	},
	{
		Id:           6,
		Pais:         "Ecuador",
		Provincia:    "Manabí",
		CiudadCanton: "Manta",
		Calle1:       "Calle Secundaria",
		Calle2:       "N/A",
		NumCasa:      "1011",
		CodPostal:    "040506",
		Referencia:   "Al lado del mercado",
		PersonaId:    4,
	},
}

func getDirecciones(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, direcciones)
}

func postDirecciones(c *gin.Context) {
	var newDireccion Direccion
	if err := c.BindJSON(&newDireccion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	direcciones = append(direcciones, newDireccion)
	c.IndentedJSON(http.StatusCreated, newDireccion)
}

func getDireccionbyID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	for _, a := range direcciones {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "La dirección no está en la lista"})
}

func putDireccion(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedDireccion Direccion
	if err := c.BindJSON(&updatedDireccion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range direcciones {
		if a.Id == intID {
			direcciones[i] = updatedDireccion
			c.IndentedJSON(http.StatusOK, updatedDireccion)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "La dirección no está en la lista"})
}

func deleteDireccion(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range direcciones {
		if a.Id == intID {
			direcciones = append(direcciones[:i], direcciones[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Dirección eliminada"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "La dirección no está en la lista"})
}

// *******************************************************************USUARIO*************************************************************************
type Usuario struct {
	Persona       `json: "persona"`
	Name          string      `json: "name"`
	Password      string      `json: "password"`
	Rol           string      `json: "rol"`
	Image         string      `json: "image"`
	Productos     []*Producto `json: "productos"`
	PedidosCompra []*Pedido   `json: "pedidosCompra"`
	PedidosVenta  []*Pedido   `json: "pedidosVenta"`
}

var usuarios = []Usuario{
	{
		Persona: Persona{
			Id:              1,
			CedulaRuc:       "0150041580",
			FechaNacimiento: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
			Apellidos:       "PACHECO HERRERA",
			Nombres:         "GALILEA KATHERINE",
			Telefono:        "0996972323",
			Email:           "galypachecoh@gmail.com",
		},
		Name:      "usuario1",
		Password:  "password123",
		Rol:       "admin",
		Image:     "image1.jpg",
		Productos: []*Producto{
			// Lista de productos del usuario
		},
		PedidosCompra: []*Pedido{
			// Lista de pedidos de compra del usuario
		},
		PedidosVenta: []*Pedido{
			// Lista de pedidos de venta del usuario
		},
	},
}

func getUsuarios(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, usuarios)
}

func postUsuarios(c *gin.Context) {
	var newUsuario Usuario
	if err := c.BindJSON(&newUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usuarios = append(usuarios, newUsuario)
	c.IndentedJSON(http.StatusCreated, newUsuario)
}

func getUsuariobyID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	for _, a := range usuarios {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El usuario no está en la lista"})
}

func putUsuario(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedUsuario Usuario
	if err := c.BindJSON(&updatedUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range usuarios {
		if a.Id == intID {
			usuarios[i] = updatedUsuario
			c.IndentedJSON(http.StatusOK, updatedUsuario)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El usuario no está en la lista"})
}

func deleteUsuario(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range usuarios {
		if a.Id == intID {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El usuario no está en la lista"})
}

// *******************************************************************PRODUCTO*************************************************************************
type Producto struct {
	Id          int     `json: "id"`
	Descripcion string  `json: "descripcion"`
	Stock       int     `json: "stock"`
	Precio      float64 `json: "precio"`
	Descuento   float64 `json: "descuento"`
	Image       string  `json: "image"`
	Detalle     string  `json: "detalle"`
	Categoria   string  `json: "categoria"`
	UsuarioId   int     `json: "usuario"`
}

var productos = []Producto{
	{
		Id:          1,
		Descripcion: "Laptop HP 15 Pulgadas",
		Stock:       10,
		Precio:      799.99,
		Descuento:   0,
		Image:       "laptop_hp.jpg",
		Detalle:     "Laptop HP con procesador Intel Core i5 y 8GB de RAM",
		Categoria:   "Electrónica",
		UsuarioId:   1,
	},
	{
		Id:          2,
		Descripcion: "Smartphone Samsung Galaxy S20",
		Stock:       5,
		Precio:      999.99,
		Descuento:   0,
		Image:       "samsung_s20.jpg",
		Detalle:     "Smartphone Samsung con cámara de alta resolución y 128GB de almacenamiento",
		Categoria:   "Electrónica",
		UsuarioId:   2,
	},
	{
		Id:          3,
		Descripcion: "Camisa Polo Ralph Lauren",
		Stock:       20,
		Precio:      59.99,
		Descuento:   0,
		Image:       "camisa_polo.jpg",
		Detalle:     "Camisa Polo Ralph Lauren de algodón, ideal para uso casual",
		Categoria:   "Moda",
		UsuarioId:   1,
	},
}

func getProductos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, productos)
}

func postProductos(c *gin.Context) {
	var newProducto Producto
	if err := c.BindJSON(&newProducto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productos = append(productos, newProducto)
	c.IndentedJSON(http.StatusCreated, newProducto)
}

func getProductobyID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	for _, a := range productos {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El producto no está en la lista"})
}

func putProducto(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedProducto Producto
	if err := c.BindJSON(&updatedProducto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range productos {
		if a.Id == intID {
			productos[i] = updatedProducto
			c.IndentedJSON(http.StatusOK, updatedProducto)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El producto no está en la lista"})
}

func deleteProducto(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range productos {
		if a.Id == intID {
			productos = append(productos[:i], productos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El producto no está en la lista"})
}

// *******************************************************************PEDIDO*************************************************************************
type Pedido struct {
	Id        int       `json: "id"`
	Fecha     time.Time `json: "fecha"`
	Subtotal  float64   `json: "subtotal"`
	Impuesto  float64   `json: "impuesto"`
	Envio     float64   `json: "envio"`
	Total     float64   `json: "total"`
	UsuarioId int       `json: "usuario"`
}

var pedidos = []Pedido{
	{
		Id:        1,
		Fecha:     time.Now(),
		Subtotal:  100.0,
		Impuesto:  12.0,
		Envio:     5.0,
		Total:     117.0,
		UsuarioId: 1,
	},
	{
		Id:        2,
		Fecha:     time.Now(),
		Subtotal:  200.0,
		Impuesto:  24.0,
		Envio:     8.0,
		Total:     232.0,
		UsuarioId: 2,
	},
}

func getPedidos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pedidos)
}

func postPedidos(c *gin.Context) {
	var newPedido Pedido
	if err := c.BindJSON(&newPedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pedidos = append(pedidos, newPedido)
	c.IndentedJSON(http.StatusCreated, newPedido)
}

func getPedidobyID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	for _, a := range pedidos {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El pedido no está en la lista"})
}

func putPedido(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedPedido Pedido
	if err := c.BindJSON(&updatedPedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range pedidos {
		if a.Id == intID {
			pedidos[i] = updatedPedido
			c.IndentedJSON(http.StatusOK, updatedPedido)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El pedido no está en la lista"})
}

func deletePedido(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range pedidos {
		if a.Id == intID {
			pedidos = append(pedidos[:i], pedidos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Pedido eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El pedido no está en la lista"})
}

// *************************************************************DETALLE PEDIDO*************************************************************************
type DetallePedido struct {
	Id         int `json: "id"`
	Cantidad   int `json: "cantidad"`
	PedidoId   int `json: "pedidoId"`
	ProductoId int `json: "productoId"`
}

var detallePedidos = []DetallePedido{
	{
		Id:         1,
		Cantidad:   2,
		PedidoId:   1,
		ProductoId: 1,
	},
	{
		Id:         2,
		Cantidad:   1,
		PedidoId:   2,
		ProductoId: 2,
	},
}

func getDetalle(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, detallePedidos)
}

func postDetalle(c *gin.Context) {
	var newDetalle DetallePedido
	if err := c.BindJSON(&newDetalle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	detallePedidos = append(detallePedidos, newDetalle)
	c.IndentedJSON(http.StatusCreated, newDetalle)
}

func getDetallebyID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	for _, a := range detallePedidos {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El detalle de pedido no está en la lista"})
}

func putDetalle(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedDetalle DetallePedido
	if err := c.BindJSON(&updatedDetalle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range detallePedidos {
		if a.Id == intID {
			detallePedidos[i] = updatedDetalle
			c.IndentedJSON(http.StatusOK, updatedDetalle)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El detalle de pedido no está en la lista"})
}

func deleteDetalle(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range detallePedidos {
		if a.Id == intID {
			detallePedidos = append(detallePedidos[:i], detallePedidos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Detalle de pedido eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El detalle de pedido no está en la lista"})
}

// ******************************************************************CARRITO*************************************************************************
type Carrito struct {
	Id        int     `json: "id"`
	Cantidad  int     `json: "cantidad"`
	Subtotal  float64 `json: "subtotal"`
	Envio     float64 `json: "envio"`
	Impuesto  float64 `json: "impuesto"`
	Total     float64 `json: "total"`
	UsuarioId int     `json: "usuario"`
}

var carritos = []Carrito{
	{
		Id:        1,
		Cantidad:  2,
		Subtotal:  200.50,
		Envio:     10.0,
		Impuesto:  25.5,
		Total:     236.0,
		UsuarioId: 3,
	},
	{
		Id:        2,
		Cantidad:  1,
		Subtotal:  150.25,
		Envio:     8.0,
		Impuesto:  20.0,
		Total:     178.25,
		UsuarioId: 4,
	},
}

func getCarritos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, carritos)
}

func postCarritos(c *gin.Context) {
	var newCarrito Carrito
	if err := c.BindJSON(&newCarrito); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	carritos = append(carritos, newCarrito)
	c.IndentedJSON(http.StatusCreated, newCarrito)
}

func getCarritobyID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}
	for _, a := range carritos {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El carrito no está en la lista"})
}

func putCarrito(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var updatedCarrito Carrito
	if err := c.BindJSON(&updatedCarrito); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range carritos {
		if a.Id == intID {
			carritos[i] = updatedCarrito
			c.IndentedJSON(http.StatusOK, updatedCarrito)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El carrito no está en la lista"})
}

func deleteCarrito(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, a := range carritos {
		if a.Id == intID {
			carritos = append(carritos[:i], carritos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Carrito eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "El carrito no está en la lista"})
}

func main() {
	//una instancia del modulo adicional que se denomina gin
	router := gin.Default()

	//vamos a crear una ruta hacia la raiz "/"
	//en otros lenguajes se conoce como rutas
	router.GET("/", func(c *gin.Context) {
		//cuando el cliente haga la petcion get localhost:8080/, devolvera una peticion http
		c.JSON(200, gin.H{
			"message": "Bienvenidos a la primera pagina del sistema",
		})
	})

	//vamos a crear puntos de salida para los procesos de listar, modificar, eliminar o ver individual
	//GET, PUT, POST, DELETE
	router.GET("/personas", getPersonas)
	//crear punto de acceso para almacenar una persona
	router.POST("/personas", postPersona)
	//crear la ruta para una consulta individual de una persona
	router.GET("/persona/:id", getPersonabyID)
	//modificar los datos de una persona por ID
	router.PUT("/persona/:id", putPersona)
	//eliminar una persona por ID
	router.DELETE("/persona/:id", deletePersona)

	//end points de direccion
	router.GET("/direcciones", getDirecciones)
	router.POST("/direcciones", postDirecciones)
	router.GET("/direcciones/:id", getDireccionbyID)
	router.PUT("/direcciones/:id", putDireccion)
	router.DELETE("/direcciones/:id", deleteDireccion)

	//end points de usuario
	router.GET("/usuarios", getUsuarios)
	router.POST("/usuarios", postUsuarios)
	router.GET("/usuarios/:id", getUsuariobyID)
	router.PUT("/usuarios/:id", putUsuario)
	router.DELETE("/usuarios/:id", deleteUsuario)

	//end points de producto
	router.GET("/productos", getProductos)
	router.POST("/productos", postProductos)
	router.GET("/productos/:id", getProductobyID)
	router.PUT("/productos/:id", putProducto)
	router.DELETE("/productos/:id", deleteProducto)

	//end points de pedido
	router.GET("/pedidos", getPedidos)
	router.POST("/pedidos", postPedidos)
	router.GET("/pedidos/:id", getPedidobyID)
	router.PUT("/pedidos/:id", putPedido)
	router.DELETE("/pedidos/:id", deletePedido)

	//end points de detalle
	router.GET("/detalle", getDetalle)
	router.POST("/detalle", postDetalle)
	router.GET("/detalle/:id", getDetallebyID)
	router.PUT("/detalle/:id", putDetalle)
	router.DELETE("/detalle/:id", deleteDetalle)

	//end points de carrito
	router.GET("/carritos", getCarritos)
	router.POST("/carritos", postCarritos)
	router.GET("/carritos/:id", getCarritobyID)
	router.PUT("/carritos/:id", putCarrito)
	router.DELETE("/carritos/:id", deleteCarrito)

	//podemos crear y lanzar las rutas de las apis para que sean consumidas
	//es en encargado de crear un servidor web para poder verlo desde el navegador
	router.Run("localhost:8080")
}
