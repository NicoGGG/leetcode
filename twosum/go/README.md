# [Twosum go solution](https://leetcode.com/problems/two-sum/)

Trying an implementation of the one-pass hashmap.

On every iteration in the input list, save the value as a key in a map, and save the list index as the value

That way, if the target - current_value is present in the dictionnary, you can return list(dict[target - current_value], current_index). In go, you can check if a key exists in a map with the second return of a map call (v, ok := mymap["mykey"]). Ok is a boolean, representing if wether or not the key was found in map. If the key is not present in the map, ok == false, and v == 0.

Don't forget to check if map[target - current_value] is equal to current_index, in which case you must not return the solution.

Also, add the new entry to the dict **after** checking if you can return a solution, or else you risk overwriting an entry that was part of the solution (see unit test 2)