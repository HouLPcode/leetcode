--
-- @lc app=leetcode id=620 lang=mysql
--
-- [620] Not Boring Movies
--
# 140 ms 44.45 %
# Write your MySQL query statement below
select *
from cinema
where mod(id,2)=1 
and description != 'boring'
order by rating desc;

