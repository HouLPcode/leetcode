--
-- @lc app=leetcode id=181 lang=mysql
--
-- [181] Employees Earning More Than Their Managers
--
# 299 ms 72.81 %
# Write your MySQL query statement below
# 获取Employee表的笛卡尔积，然后筛选
SELECT e1.Name AS Employee
from Employee AS e1 
JOIN Employee AS e2
ON e1.ManagerID = e2.Id AND e1.Salary > e2.Salary; 

