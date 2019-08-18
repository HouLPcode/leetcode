/*
 * @lc app=leetcode id=690 lang=cpp
 *
 * [690] Employee Importance
 */
/*
// Employee info
class Employee {
public:
    // It's the unique ID of each node.
    // unique id of this employee
    int id;
    // the importance value of this employee
    int importance;
    // the id of direct subordinates
    vector<int> subordinates;
};
*/
// 24 ms, faster than 97.34%
// BFS 思想 + map缓存 id->Employee 映射关系
class Solution {
public:
    int getImportance(vector<Employee*> employees, int id) {
        int rtn = 0;
        queue<int> q{{id}};
        unordered_map<int, Employee*> m;
        for (auto e : employees) m[e->id] = e;
        while (!q.empty()) {
            auto t = q.front(); q.pop();
            rtn += m[t]->importance;
            for (int num : m[t]->subordinates) {
                q.push(num);
            }
        }
        return rtn;
    }
};
