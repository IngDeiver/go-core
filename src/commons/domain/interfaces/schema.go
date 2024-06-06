// Define una interfaz para los modelos que pueden proporcionar el nombre de la colección
package baseSchemaDomain

type CollectionNamer interface {
	CollectionName() string
}