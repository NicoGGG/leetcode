import unittest
import twosum

class TestTwoSum(unittest.TestCase):

    def test_example(self):
        i1 = [2, 7, 11, 15]
        i2 = 9
        e1 = [0, 1]
        r1 = twosum.Solution.twoSum(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_2(self):
        i1 = [3, 3]
        i2 = 6
        e1 = [0, 1]
        r1 = twosum.Solution.twoSum(self, i1, i2)
        self.assertEqual(e1, r1)


if __name__ == '__main__':
    unittest.main()