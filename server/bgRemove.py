from rembg import remove
from PIL import Image
import io

def removeRequest(image: bytes) -> bytes :
    print("imageBg removing")
    opened = Image.open(io.BytesIO(image))
    removed = remove(opened)

    imgBytes = savePng(removed)
    
    return imgBytes.getvalue()

def saveJpeg(img: Image) -> io.BytesIO:
    rgb_image = Image.new("RGB", img.size, (255, 255, 255))
    rgb_image.paste(img, mask=img.split()[3])  # 3 is the alpha channel

    imgBytes = io.BytesIO()
    rgb_image.save(imgBytes, format="JPEG", quality=95, optimize=True, progressive=True)

    imgBytes.seek(0)
    return imgBytes

def savePng(img: Image) -> io.BytesIO:
    imgBytes = io.BytesIO()
    img.save(imgBytes, format="PNG")

    imgBytes.seek(0)
    return imgBytes