# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## Level
**Medium**


## Question

Given a string `s`, return the longest _palindromic substring_ in s.

Example 1:

```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

Example 2:

```
Input: s = "cbbd"
Output: "bb"
```

Constraints:

* 1 <= s.length <= 1000
* s consist of only digits and English letters.


## Answer
### Solution

The palindrome string is obtained by taking each character of the master string as the central axis, and then the longest one is obtained by comparison.

### Runtime
0 ms

### Memory Usage
2.33 MB

## Reference
### Solution

Dynamic programming: `dp[i,j] ` indicates whether the string `s[i:j]` is a palindrome string.
```
    dp[i,i] = true
    dp[i,i+1] = s[i] == s[i+1]
    dp[i,j] = ((j-i < 3) || dp[i+1,j-1]) && s[i] == s[j]
```

### Runtime
52 ms

### Memory Usage
6.73 MB