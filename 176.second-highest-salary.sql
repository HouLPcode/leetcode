--
-- @lc app=leetcode id=176 lang=mysql
--
-- [176] Second Highest Salary
--
# Write your MySQL query statement below
# 通过LIMIT OFFSET 取出来第二个数
# 注意排序后的Salary可能有重复元素，必须去重
# 没有第二数时候需要返回null
SELECT (
    SELECT DISTINCT Salary 
    FROM Employee
    ORDER BY Salary DESC 
    LIMIT 1 OFFSET 1
    ) AS SecondHighestSalary; 
