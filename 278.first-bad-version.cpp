/*
 * @lc app=leetcode id=278 lang=cpp
 *
 * [278] First Bad Version
 */
//  0 ms
// 第二种模板，直接跟右邻居比较
// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int s = 1, e = n;
        for ( ; s < e; ) {
            int mid = (e - s) / 2 + s;
            if (isBadVersion(mid) == false) { // mid左侧全是false
                s = mid + 1;
            } else {
                e = mid;
            }
        }
        return s; // mid是最后一个false，s是第一个true
    }
};

