# _*_ coding: utf-8 _*_

import unittest

from hello import Student

class TestStudent(unittest.TestCase):
    
    def test_80_to_100(self):
        s1, s2 = Student('Bart', 80), Student('Lisa', 100)
        self.assertEqual(s1.get_grade(), 'A')
        self.assertEqual(s2.get_grade(), 'A')
        
    def test_60_to_80(self):
        s1, s2 = Student('Bart', 60), Student('Lisa', 79)
        self.assertEqual(s1.get_grade(), 'B')
        self.assertEqual(s2.get_grade(), 'B')
        
    def test_0_to_60(self):
        s1, s2 = Student('Bart', 0), Student('Lisa', 59)
        self.assertEqual(s1.get_grade(), 'C')
        self.assertEqual(s2.get_grade(), 'C')
        
#if __name__ == '__main__':
#   unittest.main()