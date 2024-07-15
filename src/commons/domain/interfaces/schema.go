package baseSchemaDomain

// Defines an interface to the models that can provide the name of the collection
//  Use this if the collection name is hard to infer
type CollectionNamer interface {
	CollectionName() string
}