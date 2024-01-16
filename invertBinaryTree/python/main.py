from typing import Optional


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def invertTree(root: Optional[TreeNode]) -> Optional[TreeNode]:
    if root:
        tmp = root.left
        root.left = root.right
        root.right = tmp
        invertTree(root.left)
        invertTree(root.right)

    return root


def insertNode(root: TreeNode, node: TreeNode):
    if root.val > node.val:
        if root.left:
            insertNode(root.left, node)
        else:
            root.left = node
    else:
        if root.right:
            insertNode(root.right, node)
        else:
            root.right = node


def buildTree(arr: list) -> TreeNode:
    tree = TreeNode(arr[0])
    for i in range(1, len(arr)):
        insertNode(tree, TreeNode(arr[i]))

    return tree
