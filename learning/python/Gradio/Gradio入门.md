### 快速搭建AI算法可视化部署演示

- **Gradio**

  Gradio的优势在于易用性，代码结构相比Streamlit简单，只需简单定义输入和输出接口即可快速构建简单的交互页面，更轻松部署模型。适合场景相对简单，想要快速部署应用的开发者。便于分享：gradio可以在启动应用时设置share=True参数创建外部分享链接，可以直接在微信中分享给用户使用。

- **Streamlit**

  Streamlit的优势在于可扩展性，相比Gradio复杂，完全熟练使用需要一定时间。可以使用Python编写完整的包含前后端的交互式应用。适合场景相对复杂，想要构建丰富多样交互页面的开发者。



gradio 安装：

> pip install gradio
>
> *#为了更快安装，可以使用清华镜像源*
> pip install -i https://pypi.tuna.tsinghua.edu.cn/simple gradio



