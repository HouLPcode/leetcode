/*
 * @lc app=leetcode id=559 lang=cpp
 *
 * [559] Maximum Depth of N-ary Tree
 */
/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
// 128 ms, faster than 97.69%
// BFS
class Solution {
public:
    int maxDepth(Node* root) {
        if (!root) 
            return 0;
        
        queue<Node*>queue;
        queue.push(root);
        int max_depth = 0;
        while (!queue.empty()) {
            max_depth++;			
            for (int size = queue.size(); size; size--) {
                Node* curr = queue.front(); 
                queue.pop();
                for (Node* it : curr->children)
                    queue.push(it);
            }
        }
        return max_depth;
    }
};

