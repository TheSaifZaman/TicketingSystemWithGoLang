module main

go 1.18

replace (
	variablesAndFunctions v0.0.0 => ./VariablesAndFunctions
	pointers v0.0.0 => ./Pointers
	typesAndStructs v0.0.0 => ./TypesAndStructs
	receiversStructWithFunctions v0.0.0 => ./ReceiversStructWithFunctions
	mapsAndSlices v0.0.0 => ./MapsAndSlices
	conditionals v0.0.0 => ./Conditionals
)

require (
	variablesAndFunctions v0.0.0
	pointers v0.0.0
	typesAndStructs v0.0.0
	receiversStructWithFunctions v0.0.0
	mapsAndSlices v0.0.0
	conditionals v0.0.0
)