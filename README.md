Generic Collection
This package provides a generic collection type in Go that can store any type of value. It provides methods to add, remove, and find elements in the collection. The package also includes some useful methods to perform common operations on collections such as filtering, grouping, and selecting elements.

Creating a Generic Collection
To create a new Generic Collection, you can call the NewGenericCollection function and pass in the initial values as a variadic argument:

go
Copy code
collection := NewGenericCollection(1, 2, 3, 4, 5)
This will create a new GenericCollection with the values 1, 2, 3, 4, and 5.

Adding Elements to a Collection
To add elements to a GenericCollection, you can use the Add or AddRange methods:

go
Copy code
collection := NewGenericCollection()
collection.Add("hello")
collection.AddRange(1, 2, 3)
This will add the values "hello", 1, 2, and 3 to the collection.

Removing Elements from a Collection
To remove an element from a GenericCollection, you can use the Remove method and pass in the element to be removed:

go
Copy code
collection := NewGenericCollection(1, 2, 3, 4, 5)
collection.Remove(3)
This will remove the value 3 from the collection.

To remove an element at a specific index, you can use the RemoveAt method and pass in the index:

go
Copy code
collection := NewGenericCollection(1, 2, 3, 4, 5)
collection.RemoveAt(3)
This will remove the element at index 3 (which is the fourth element in the collection).

To remove a range of elements from the collection, you can use the RemoveRange method and pass in the index and the number of elements to remove:

go
Copy code
collection := NewGenericCollection(1, 2, 3, 4, 5)
collection.RemoveRange(1, 2)
This will remove two elements starting from index 1, so the collection will contain the values 1, 4, and 5.

Checking if an Element Exists in a Collection
To check if a value exists in a GenericCollection, you can use the Contains method and pass in the element to check:

go
Copy code
collection := NewGenericCollection(1, 2, 3, 4, 5)
if collection.Contains(3) {
    fmt.Println("The collection contains the value 3")
}
This will print The collection contains the value 3.

Finding Elements in a Collection
To find an element in a GenericCollection that satisfies a certain condition, you can use the Find method and pass in a predicate function:

go
Copy code
collection := NewGenericCollection("apple", "banana", "cherry")
result := collection.Find(func(item interface{}) bool {
    return len(item.(string)) > 5
})
fmt.Println(result)
This will print cherry, which is the first value in the collection that has a length greater than 5.

To find all elements in a collection that satisfy a certain condition, you can use the FindMany method and pass in a predicate function:

go
Copy code
collection := NewGenericCollection(1, 2, 3, 4, 5)
result := collection.FindMany(func(item interface{}) bool {
    return item.(int) > 2
})
fmt.Println(result)
This will print `[3 4 5]
