--
-- @lc app=leetcode id=183 lang=mysql
--
-- [183] Customers Who Never Order
--
-- 334 ms 5 %
# Write your MySQL query statement below
-- 左联结
SELECT Customers.Name AS Customers
FROM Customers LEFT JOIN Orders
ON Customers.Id = Orders.CustomerId
WHERE Orders.CustomerId IS NULL; #注意此处怎么判空的

-- 子查询
-- select customers.name as 'Customers'
-- from customers
-- where customers.id not in
-- (
--     select customerid from orders
-- );
