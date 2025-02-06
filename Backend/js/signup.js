document.getElementById('signup-form').addEventListener('submit', async function (event) {
    event.preventDefault();

    const email = document.getElementById('signup-email').value;
    const password = document.getElementById('signup-password').value;
    const confirmPassword = document.getElementById('signup-confirm-password').value;

    if (password !== confirmPassword) {
        alert('Passwords do not match!');
        return;
    }

    try {
        const response = await fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, password }),
        });

        const data = await response.json();

        if (response.ok) {
            // เก็บ user_id ใน localStorage เพื่อใช้ในหน้าถัดไป
            localStorage.setItem('user_id', data.user_id);
            // เปลี่ยนไปยังหน้าเลือกบทบาท
            window.location.href = 'Role.html';
        } else {
            // แสดงข้อผิดพลาดจากเซิร์ฟเวอร์
            alert(data.error || 'Registration failed!');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred. Please try again later.');
    }
});
