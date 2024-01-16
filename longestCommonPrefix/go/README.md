# [Longest Common Prefix go solution](https://leetcode.com/problems/longest-common-prefix)
    
## Try 1

For my approach, I keep using a map to help me compare the content of array without having to store to much temporary data.

First, I need to determine the least long string in the input array to set the limit of my for loop properly.
Since we know from the question that strings can't be more than 200 characters long, I init l to 200, and keep lowering it until I checked le length of all strings

After that, I start a for loop starting at 0, representing the index of each string I want to compare.

Then I iterate over each string, and put every i character in a map, with i as the key.

Then, if the map is not exactly 1 long, it means strings diverged at this point, and I return the prefix string.

If map length is 1, I append the character of index i (of any string, since they are all the same at this point) to r

## Try 2

First approach worked, and I validated the challenge. However, memory consumption wasn't optimal.

I realised I can save on some things.

First, I don't need to find the lowest string length. I can arbitrarily take the first one. I will need to check the length of the word of
every string in the array if it is lower than i however.

All this doesn't really do much as the problem isn't really that complex, but it's still nice to have something cleaner.