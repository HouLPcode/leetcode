--
-- @lc app=leetcode id=627 lang=mysql
--
-- [627] Swap Salary
--
# 157 ms 49.57 %
# IF 函数的应用 IF(expr,v1,v2)
# Write your MySQL query statement below
UPDATE salary
SET sex = IF(sex = 'm', 'f', 'm')

