--
-- @lc app=leetcode id=182 lang=mysql
--
-- [182] Duplicate Emails
--
# 246ms 6.95%
# Write your MySQL query statement below
SELECT Email 
from (
    select Email, count(Email) as num 
    from Person 
    group by Email
) AS statistic #此处必须有as命名
where num > 1;
