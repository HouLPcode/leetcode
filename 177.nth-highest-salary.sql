--
-- @lc app=leetcode id=177 lang=mysql
--
-- [177] Nth Highest Salary
--
# 178 ms 88.64 %
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE M INT;
  SET M=N-1;
  RETURN (
      # Write your MySQL query statement below.  
      SELECT ifnull(
        (SELECT Salary 
        FROM Employee
        GROUP BY Salary
        ORDER BY Salary DESC 
        LIMIT 1 offset M),null) AS 'getNthHighestSalary(2)'
  );
END

