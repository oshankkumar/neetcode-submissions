/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     constructor(val = 0, left = null, right = null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */


type Node = TreeNode | null;

class Solution {
  /**
   * @param {Node} root
   * @return {number}
   */

  height(node: Node): number {
    if (node === null) {
      return 0;
    }
    const leftHeight = this.height(node.left);
    const rightHeight = this.height(node.right);
    return Math.max(leftHeight, rightHeight) + 1;
  }

  diameterOfBinaryTree(root: Node): number {
    if (root === null) {
      return 0;
    }
    const leftHeight = this.height(root.left);
    const rightHeight = this.height(root.right);
    const diameterThroughRoot = leftHeight + rightHeight;
    const leftDiameter = this.diameterOfBinaryTree(root.left);
    const rightDiameter = this.diameterOfBinaryTree(root.right);
    return Math.max(diameterThroughRoot, leftDiameter, rightDiameter);
  }
}
