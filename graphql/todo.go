package graphql

import (
	"github.com/graphql-go/graphql"
	. "github.com/jerminehu/ait/model"
)

// define custom GraphQL ObjectType `todoType` for our Golang struct `Todo`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Description:"this type is TodoType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Description:"The data id flied",
		},
		"text": &graphql.Field{
			Type: graphql.String,
			Description:"The data text flied",
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
			Description:"The data done flied",
		},
	},
})

//WrapTodoQuerys for
func WrapTodoQuerys(schema *GraphqlSchema) {
	todoFiled(schema)
	lastTodoFiled(schema)
	todoListFiled(schema)
	createTodoFiled(schema)
	updateTodoFiled(schema)
}

// to set todoFiled for rootQuery
func todoFiled(schema *GraphqlSchema)   {

	schema.RootQuery.AddFieldConfig("todo",
		&graphql.Field{
			Type:        todoType,
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery, isOK := params.Args["id"].(string)
				if isOK {
					// Search for el with id
					for _, todo := range TodoList {
						if todo.ID == idQuery {
							return todo, nil
						}
					}
				}

				return Todo{}, nil
			},
		})
}

// to set lastTodoFiled for rootQuery
func lastTodoFiled(schema *GraphqlSchema)  {

	schema.RootQuery.AddFieldConfig("lastTodo", &graphql.Field{
		Type:        todoType,
		Description: "Last todo added",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return TodoList[len(TodoList)-1], nil
		},
	})

}


// to set todoListFiled for rootQuery
func todoListFiled(schema *GraphqlSchema)   {

	schema.RootQuery.AddFieldConfig(
		"todoList", &graphql.Field{
		Type:        graphql.NewList(todoType),
		Description: "List of todos <br> `curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}`",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return TodoList, nil
		},
	},)
}


// to set createTodoFiled for RootMutation
func createTodoFiled(schema *GraphqlSchema)   {

	schema.RootMutation.AddFieldConfig("createTodo", &graphql.Field{
		Type:        todoType, // the return type for this field
		Description: "Create new todo",
		Args: graphql.FieldConfigArgument{
			"text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {

			// marshall and cast the argument value
			text, _ := params.Args["text"].(string)

			// figure out new id
			newID := RandStringRunes(8)

			// perform mutation operation here
			// for e.g. create a Todo and save to DB.
			newTodo := Todo{
				ID:   newID,
				Text: text,
				Done: false,
			}

			TodoList = append(TodoList, newTodo)

			// return the new Todo object that we supposedly save to DB
			// Note here that
			// - we are returning a `Todo` struct instance here
			// - we previously specified the return Type to be `todoType`
			// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
			return newTodo, nil
		},
	},)
}

// to set updateTodoFiled for RootMutation
func updateTodoFiled(schema *GraphqlSchema)   {

	schema.RootMutation.AddFieldConfig("updateTodo", &graphql.Field{
		Type:        todoType, // the return type for this field
		Description: "Update existing todo, mark it done or not done",
		Args: graphql.FieldConfigArgument{
			"done": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			done, _ := params.Args["done"].(bool)
			id, _ := params.Args["id"].(string)
			affectedTodo := Todo{}

			// Search list for todo with id and change the done variable
			for i := 0; i < len(TodoList); i++ {
				if TodoList[i].ID == id {
					TodoList[i].Done = done
					// Assign updated todo so we can return it
					affectedTodo = TodoList[i]
					break
				}
			}
			// Return affected todo
			return affectedTodo, nil
		},
	},)
}



