class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        return self.isSymmetricHelper(root, root)

    def isSymmetricHelper(self, p: TreeNode, q: TreeNode) -> bool:
        if not p and not q:
            return True
        if not p or not q:
            return False
        
        return p.val == q.val and \
            self.isSymmetricHelper(p.left, q.right) and \
            self.isSymmetricHelper(p.right, q.left)
