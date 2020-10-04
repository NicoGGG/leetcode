# [Reverse integer go solution](https://leetcode.com/problems/reverse-integer/)
    
## Try 1: Using modulo

We can reverse the digit of an integer by using the modulo (%) 10 operation to extract the last digit of the current number.
We can then pop it from the current number by doing a division by 10.

Special edge case we have to consider is when the reversed number is > maxInt32. 