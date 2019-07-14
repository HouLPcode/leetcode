--
-- @lc app=leetcode id=596 lang=mysql
--
-- [596] Classes More Than 5 Students
--
# 219 ms   36.5 %
# Write your MySQL query statement below
select class
from courses
group by class
having count(distinct student) > 4; # 注意此处的distinct，数据中有重复名字的人只计算一次？？？

