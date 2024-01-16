import unittest

from main import invertTree, buildTree, TreeNode, insertNode


class TestInvert(unittest.TestCase):
    def test_invert(self):
        treenode1: TreeNode = buildTree([4, 2, 7, 1, 3, 6, 9])
        inverted_treenode1: TreeNode = invertTree(treenode1)

        self.assertEqual(inverted_treenode1.val, 4)
        self.assertTrue(inverted_treenode1.left is not None)
        self.assertEqual(inverted_treenode1.left.val, 7)
        self.assertTrue(inverted_treenode1.right is not None)
        self.assertEqual(inverted_treenode1.right.val, 2)
        self.assertTrue(inverted_treenode1.left.left is not None)
        self.assertEqual(inverted_treenode1.left.left.val, 9)
        self.assertTrue(inverted_treenode1.left.right is not None)
        self.assertEqual(inverted_treenode1.left.right.val, 6)
        self.assertTrue(inverted_treenode1.right.left is not None)
        self.assertEqual(inverted_treenode1.right.left.val, 3)
        self.assertTrue(inverted_treenode1.right.right is not None)
        self.assertEqual(inverted_treenode1.right.right.val, 1)
