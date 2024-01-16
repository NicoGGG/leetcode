import unittest
import medianTwoSortedArrays


class TestMedianTwoSortedArray(unittest.TestCase):
    def test_example1(self):
        i1 = [1, 3]
        i2 = [2]
        e1 = 2
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays1(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_example2(self):
        i1 = [1, 2]
        i2 = [3, 4]
        e1 = 2.5
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays1(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_trivial_example1(self):
        i1 = [1]
        i2 = []
        e1 = 1
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays1(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_trivial_example1(self):
        i1 = [0, 0]
        i2 = [0, 0]
        e1 = 0
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays1(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_optimal_example1(self):
        i1 = [1, 3]
        i2 = [2]
        e1 = 2
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays2(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_optimal_example2(self):
        i1 = [1, 2]
        i2 = [3, 4]
        e1 = 2.5
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays2(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_optimal_trivial_example1(self):
        i1 = [1]
        i2 = []
        e1 = 1
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays2(self, i1, i2)
        self.assertEqual(e1, r1)

    def test_optimal_trivial_example2(self):
        i1 = [0, 0]
        i2 = [0, 0]
        e1 = 0
        r1 = medianTwoSortedArrays.Solution.findMedianSortedArrays2(self, i1, i2)
        self.assertEqual(e1, r1)


if __name__ == "__main__":
    unittest.main()
