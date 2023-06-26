class Animal(object):
    @property
    def kind(self):
        #为了和方法名做区分，在属性前面加一个下划线
        return self._kind
    @kind.setter
    def kind(self, value):
        if not isinstance(value, str):
            raise ValueError('kind must be a str!')
        self._kind = value