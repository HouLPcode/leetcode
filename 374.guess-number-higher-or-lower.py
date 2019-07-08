#
# @lc app=leetcode id=374 lang=python
#
# [374] Guess Number Higher or Lower
#
# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num):
# 16 ms 80.07 %
class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        s = 1
        e = n
        while s <= e:
            mid = (e-s)/2 + s
            res = guess(mid)
            if res == 0 :
                return mid
            elif res == 1 : 
                s = mid + 1
            elif res == -1 :
                e = mid - 1

