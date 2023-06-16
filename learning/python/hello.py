class Message(object):
    def __init__(self, name):
        self.name = name
        
m = Message('ssss')
d = {'m1':m}

value = d['m1']
value.name = 'dfad'

if d.get('m1'):
    print('d m1 exist')

if d.get('m2'):
    print('d m2 exist')

for k,v in d.items():
    print(k,v.name)
