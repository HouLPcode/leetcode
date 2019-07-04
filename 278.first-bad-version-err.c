/*
 * @lc app=leetcode id=278 lang=c
 *
 * [278] First Bad Version
 */
// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

// 二分法
int firstBadVersion(int n) {
    int s, e;
    s = 1; // 题目要求从1开始计数
    e = n;
    for ( ; s <= e; ) {
        int mid = (e - s) / 2 + s;
        if (isBadVersion(mid) == false ){
            if (mid == e) {
                return -1; //全是false
            }else if (isBadVersion(mid+1) == true){
                return mid + 1;
            } else if (isBadVersion(mid+1) == false){
                s = mid + 1;
            }
        } else if (isBadVersion(mid) == true ){
            if (mid == 1) {
                return 1;
            }
            e = mid - 1;
        }
    }
    return -1; //全是false
}

