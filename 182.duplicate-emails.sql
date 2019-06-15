--
-- @lc app=leetcode id=182 lang=mysql
--
-- [182] Duplicate Emails
--
# 228 ms 19.77 %
# Write your MySQL query statement below
SELECT Email 
from  Person 
GROUP BY Email
HAVING count(Email) > 1
;
