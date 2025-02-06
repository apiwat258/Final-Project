// เปิด-ปิด QR Reader
const qrIcon = document.getElementById('qr-icon');
const qrReaderContainer = document.getElementById('qr-reader-container');
const closeReader = document.getElementById('close-reader');
const qrResult = document.getElementById('qr-result');
let html5QrCode;

qrIcon.addEventListener('click', () => {
    qrReaderContainer.classList.remove('hidden');
    startQrReader();
});

closeReader.addEventListener('click', () => {
    qrReaderContainer.classList.add('hidden');
    stopQrReader();
});

// เริ่มต้น QR Reader
function startQrReader() {
    if (!html5QrCode) {
        html5QrCode = new Html5Qrcode("qr-reader");
    }
    html5QrCode.start(
        { facingMode: "environment" }, // ใช้กล้องหลัง
        {
            fps: 10, // frame per second
            qrbox: { width: 250, height: 250 } // ขนาดกรอบ
        },
        (decodedText) => {
            qrResult.textContent = `Result: ${decodedText}`;
            stopQrReader(); // หยุดเมื่อสแกนสำเร็จ
        },
        (errorMessage) => {
            console.log(`QR Code no match: ${errorMessage}`);
        }
    ).catch((err) => {
        console.error(`Unable to start scanning: ${err}`);
    });
}

// หยุด QR Reader
function stopQrReader() {
    if (html5QrCode) {
        html5QrCode.stop().then(() => {
            html5QrCode.clear();
        }).catch((err) => {
            console.error(`Unable to stop scanning: ${err}`);
        });
    }
}
