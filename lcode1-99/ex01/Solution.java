//package ex01;

/**
 * 求给定二叉树的最小深度。
 * 最小深度是指树的根结点到最近叶子结点的最短路径上结点的数量。
 */
public class Solution {
    public int run(TreeNode root) {
        if (root == null) {
            return 0;
        }
        return findLeaf(root, 1);
    }

    private int findLeaf(TreeNode head, int level) {
        if (head.left == null && head.right == null) {
            return level;
        } else if (head.left == null && head.right != null) {
            return findLeaf(head.right, level + 1);
        } else if (head.left != null && head.right == null) {
            return findLeaf(head.left, level + 1);
        } else {
            int left = findLeaf(head.left, level + 1);
            int right = findLeaf(head.right, level + 1);
            return left > right ? right : left;
        }

    }
}


