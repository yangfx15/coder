from enum import Enum

class LabelWidget(Enum):
    yes_no = "yes_no"
    flag = "flag"
    likert = "likert"
    category_flag = "category_flag"
    domain_flag = "domain_flag"

class Potato(Enum):
    def __new__(cls, widget: LabelWidget):
        obj.widget = widget
        return obj
    
    spam = "spam", LabelWidget.flag
    
    @classmethod
    def make(cls, num):
        potatos = []
        for i in range(num):
            potatos.append(cls.__new__(cls))
        return potatos
        
all_potatos = Potato.make(5)

print(all_potatos)