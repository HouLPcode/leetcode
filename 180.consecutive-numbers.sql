--
-- @lc app=leetcode id=180 lang=mysql
--
-- [180] Consecutive Numbers
--
# 331 ms 83.19 % 
# Write your MySQL query statement below
select 
    distinct l1.Num as ConsecutiveNums # 一定注意去重
from 
    Logs l1, Logs l2, Logs l3
where l1.Id = l2.Id - 1 #通过Id差1保证三条记录连续
and   l2.Id = l3.Id - 1
and   l1.Num = l2.Num
and   l2.Num = l3.Num 
