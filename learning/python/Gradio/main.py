import gradio as gr
import cv2
import numpy as np
import os
import argparse
from PIL import Image
import numpy as np

import torch
from torchvision.transforms.functional import to_tensor, to_pil_image

from model import Generator
from PIL import Image, ImageOps, ImageFilter

# 转换成漫画风格
def toCarttonStyle(img_rgb):
    # 属性设置
    num_down = 4  # 缩减像素采样的数目
    num_bilateral = 15  # 定义双边滤波的数目

    # 用高斯金字塔降低取样
    img_color = img_rgb
    for _ in range(num_down):
        img_color = cv2.pyrDown(img_color)

    # 重复使用小的双边滤波代替一个大的滤波
    for _ in range(num_bilateral):
        img_color = cv2.bilateralFilter(img_color, d=9, sigmaColor=9, sigmaSpace=7)

    # 升采样图片到原始大小
    for _ in range(num_down):
        img_color = cv2.pyrUp(img_color)

    # 转换为灰度并且使其产生中等的模糊
    img_gray = cv2.cvtColor(img_rgb, cv2.COLOR_RGB2GRAY)
    img_blur = cv2.medianBlur(img_gray, 7)

    # 检测到边缘并且增强其效果
    img_edge = cv2.adaptiveThreshold(img_blur, 255,
                                     cv2.ADAPTIVE_THRESH_MEAN_C,
                                     cv2.THRESH_BINARY,
                                     blockSize=9,
                                     C=2)
	
	# 算法处理后，照片的尺寸可能会不统一
    # 把照片的尺寸统一化
    height=img_rgb.shape[0]
    width = img_rgb.shape[1]
    img_color=cv2.resize(img_color,(width,height))
    
    # 转换回彩色图像
    img_edge = cv2.cvtColor(img_edge, cv2.COLOR_GRAY2RGB)
    img_cartoon = cv2.bitwise_and(img_color, img_edge)
    
    return img_cartoon

# 透明度转换，素描转换的一部分
def dodge(a, b, alpha):
    # alpha为图片透明度
    return min(int(a * 255 / (256 - b * alpha)), 255)


# 图片转换为素描
def toSketchStyle(img):
    blur = 25
    alpha = 1.0

    # 将文件转成灰色
    img1 = img.convert('L')

    img2 = img1.copy()

    img2 = ImageOps.invert(img2)

    # 模糊度
    for i in range(blur):
        img2 = img2.filter(ImageFilter.BLUR)

    width, height = img1.size
    for x in range(width):
        for y in range(height):
            a = img1.getpixel((x, y))
            b = img2.getpixel((x, y))
            img1.putpixel((x, y), dodge(a, b, alpha))

    return img1

def to_black(image):
    output = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    return output

def load_image(image_path, x32=False):
    img = Image.open(image_path).convert("RGB")

    if x32:
        def to_32s(x):
            return 256 if x < 256 else x - x % 32
        w, h = img.size
        img = img.resize((to_32s(w), to_32s(h)))

    return img

def carton(image):
    device = "cpu"
    
    net = Generator()
    net.load_state_dict(torch.load(args.checkpoint, map_location="cpu"))
    net.to(device).eval()
    print(f"model loaded: {args.checkpoint}")
    
    os.makedirs(args.output_dir, exist_ok=True)

    if os.path.splitext(image_name)[-1].lower() not in [".jpg", ".png", ".bmp", ".tiff"]:
        continue
        
    image = load_image(os.path.join(args.input_dir, image_name), args.x32)

    with torch.no_grad():
        image = to_tensor(image).unsqueeze(0) * 2 - 1
        out = net(image.to(device), args.upsample_align).cpu()
        out = out.squeeze(0).clip(-1, 1) * 0.5 + 0.5
        out = to_pil_image(out)

    strs = str.split(args.checkpoint, "/")

    image_name = strs[-1] + "_" + image_name
    out.save(os.path.join(args.output_dir, image_name))
    
    image = load_image(os.path.join(args.input_dir, image_name), args.x32)
    
    return cartoon

interface = gr.Interface(fn=carton, inputs="image", outputs="image")
interface.launch()