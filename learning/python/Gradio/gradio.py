from gradio import Interface

def greet(name):
    return "Hello " + name + "!"

demo = Interface(fn=greet, inputs="text", outputs="text")
demo.launch()