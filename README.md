# Ginger - for go generate a micro service from data schema
Add following line in top of your data schema file 
// go:generate gigner $GOFILE
and also add following line above each data type you want to build api
// @ginger
then, type "go generate"

# Example data schema file
You can generate following data type from JSON-to-Go online tool with a sample of your data object in Json format.




