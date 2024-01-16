# [Length of Longest Substring go solution](https://leetcode.com/problems/longest-substring-without-repeating-characters)
  
## Try 1: Trying a simple map implementation

Creating a map and filling it with keys, and string index, by iterating in the string.

That way, whenever I try to add a key, value where key already exists, I can return the current index, which will be the biggest.

The problem is that this assumes that the longest substring in the string is always the first one. When faced with the test "pwwkew", my function was returning 2 ("pw"), while the longest substring was "wke".

## Try 2: Trying a list of longest streak, and returning the max

Adding to that logic, I just stored the longest substring found by Method 1 in a list, and kept iterating, reseting the map in the process.

When the string is over, I return the max int from the list

This method passed a number of tests, but had a fault. It didn't work if the longest string was actually in the middle of a long string.

The test "dvdf" didn't pass as the longest string was "vdf", which was in the middle of "dv", and "df".

## Try 3: Storing length of all substring in a map

So next, I figured I should go back to a map implementation, this time storing the length (value) of every string starting with character c (key)

Whenever a character in the string iteration already had a key in the map, I reset the value of this key to 1.

If the character is not in the map, it is added, and all values in the map are incremented by 1.

At the end, I return the max value of the map.

## Try 4: Back to Method1

All these got a little too overcomplicated. I figured the first method was the right one, it just needed some extension.

The approach of validating a substring when we find a duplicate character is probably the right one. Now I need to find a way to keep iterating, while storing the current longest string, and not lose the length of the next substring, even if 
they are in the finished substring (example dvdf)

So, like I wrote in comment, the solution was that when we find a char that is in the map, all substring before the character starting index are over.

By storing the max length at every iteration in the string s, we can ignore every char with a starting index < index of the duplicate char without having to
reiterate over the map (like in method 2 and 3). We just need to store the new starting point, and keep iterating over the string until the end.

When it's over, we return the last max_length that we stored, which is the longest substring.