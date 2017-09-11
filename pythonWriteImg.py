from PIL import Image
from sys import exit

def read_file(path):
    with open(path, 'r') as f:
        return f.read()

if __name__ == '__main__':
    image_in_path = "imgIn.bmp"
    image_out_path = "imgOut.bmp"

    message = read_file("message.txt")

    img = Image.open(image_in_path)
    img = img.convert('RGB')
    img.getpixel((20, 300))
    imgarray = img.load()

    if len(message) > (img.size[0] * img.size[1])-3:
        print(">> Text too long to this image.")
        exit(1)

    message_size = len(message)
    for w in range(0, 3):
        lt = []
        for i in range(0, 3):
            if (message_size - 255) > 0:
                lt.append(255)
                message_size -= 255
            else:
                lt.append(message_size)
                message_size = 0
                if i == 0:
                    lt.append(0)
                    lt.append(0)
                elif i == 1:
                    lt.append(0)
                else:
                    break
        tp = (lt[0], lt[1], lt[2])
        imgarray[0, w] = tp

    red_sum = 0
    counter = 0
    message_counter = 0
    for h in range(0, img.size[0]):
        for w in range(0, img.size[1]):
            if h == 0 and w < 3:
                w = 3
            if counter == 8 and message_counter < len(message):
                counter = -1

                blue = imgarray[h, w][0]
                green = imgarray[h, w][1]
                red = (red_sum % 255) + ord(message[message_counter])
                message_counter += 1

                imgarray[h, w] = (blue, green, red)

                red_sum = 0
            else:
                red_sum += imgarray[h, w][2]
            counter += 1

    img.save(image_out_path)
