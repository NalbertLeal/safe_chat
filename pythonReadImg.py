from PIL import Image
from sys import exit

if __name__ == '__main__':
    image_in_path = "imgIn.bmp"

    img = Image.open(image_in_path)
    img = img.convert('RGB')
    img.getpixel((20, 300))
    imgarray = img.load()

    message_size = 0
    for w in range(0, 3):
        message_size += imgarray[0, w][0]
        message_size += imgarray[0, w][1]
        message_size += imgarray[0, w][2]

    print(message_size, '\n')

    red_sum = 0
    counter = 0
    message_counter = 0
    message = ''
    for h in range(0, img.size[0]):
        for w in range(0, img.size[1]):
            if message_counter == message_size:
                break
            elif counter == 8:
                counter = -1

                message += chr( imgarray[h, w][2] - (red_sum % 255) )
                message_counter += 1

                red_sum = 0
            else:
                red_sum += imgarray[h, w][2]
            counter += 1

    file = open("messageOut.txt", "w")
    file.write(message)
    file.close()
