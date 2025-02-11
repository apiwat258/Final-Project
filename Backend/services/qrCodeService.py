import qrcode
from pyzbar.pyzbar import decode
from PIL import Image

def generate_qr_code(data: str, file_path: str):
    """
    Generate a QR Code from the provided data and save it as an image file.
    :param data: The data to encode into the QR code.
    :param file_path: The path to save the generated QR code image.
    """
    qr = qrcode.QRCode(
        version=1,
        error_correction=qrcode.constants.ERROR_CORRECT_L,
        box_size=10,
        border=4,
    )
    qr.add_data(data)
    qr.make(fit=True)

    img = qr.make_image(fill="black", back_color="white")
    img.save(file_path)
    return file_path

def read_qr_code(file_path: str) -> str:
    """
    Read and decode the QR Code from an image file.
    :param file_path: The path to the QR code image.
    :return: The decoded data from the QR code.
    """
    img = Image.open(file_path)
    decoded_objects = decode(img)
    
    if decoded_objects:
        return decoded_objects[0].data.decode("utf-8")
    return "No QR Code found"
