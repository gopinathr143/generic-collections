package main

import (
	"fmt"
	"github.com/gopinathr143/generic-collections/collections"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	people1 := &Person{"Alice", 25}
	people2 := &Person{"Bob", 30}
	people3 := &Person{"Charlie", 35}
	// create a new GenericCollection from the people slice
	gc := collections.NewGenericCollection(people1, people2, people3)

	// count the number of items in the collection
	fmt.Println("Count:", gc.Count())

	// add a new person to the collection
	gc.Add(Person{"David", 40})

	// remove an item from the collection
	gc.Remove(people1)

	// remove an item at a specific index from the collection
	gc.RemoveAt(1)

	// remove multiple items from the collection
	gc.RemoveRange(0, 2)

	// check if the collection contains a person
	fmt.Println("Contains Bob:", gc.Contains(people2))

	// get the index of a person in the collection
	fmt.Println("Index of Charlie:", gc.IndexOf(people3))

	gc.Add(Person{"David", 40})
	gc.Add(Person{"Peter", 24})
	gc.Add(Person{"Paul", 36})

	// find the first person with an age greater than 30
	fmt.Println("First person with age > 30:", gc.Find(func(item interface{}) bool {
		return item.(Person).Age > 30
	}))

	// find all people with an age greater than 30
	fmt.Println("All people with age > 30:", gc.FindMany(func(item interface{}) bool {
		return item.(Person).Age > 30
	}))

	// select the names of all people in the collection
	names := gc.Select(func(item interface{}) interface{} {
		return item.(Person).Name
	})

	fmt.Println("Names:", names)

	// filter the people in the collection by age
	filtered := gc.Where(func(item interface{}) bool {
		return item.(Person).Age > 30
	})

	fmt.Println("Filtered:", filtered)

	// group the people in the collection by age
	groups := gc.GroupBy(func(item interface{}) interface{} {
		return item.(Person).Age
	})

	fmt.Println("Groups:", groups)

	// calculate the sum of ages using the Sum method
	sum := gc.Sum(func(item interface{}) float64 {
		return float64(item.(Person).Age)
	})

	fmt.Println("Sum of ages (using Sum method):", sum)

	// find the maximum age using the Max method
	max := gc.Max(func(item interface{}) float64 {
		return float64(item.(Person).Age)
	})

	fmt.Println("Maximum age (using Max method):", max)

	minAge := gc.Min(func(item interface{}) float64 {
		return float64(item.(Person).Age)
	})
	fmt.Println("Youngest person is", minAge, "years old.")

	// example of using Sort method to sort people by age
	gc.Sort(func(i, j interface{}) bool {
		return i.(Person).Age < j.(Person).Age
	})
	fmt.Println("People sorted by age (ascending):", gc.GetAll())

	// example of using Reverse method to reverse the order of people
	gc.Reverse()
	fmt.Println("People reversed:", gc.GetAll())

	// example of using FindByIndex method to get the person at index 2
	p := gc.FindByIndex(2)
	if p == nil {
		fmt.Println("Person not found.")
	} else {
		fmt.Println("Person at index 2:", p)
	}

	// example of using ReverseSort method to sort people by age in descending order
	gc.ReverseSort(func(i, j int) bool {
		return gc.GetAll()[i].(Person).Age < gc.GetAll()[j].(Person).Age
	})
	fmt.Println("People sorted by age (descending):", gc.GetAll())
}
