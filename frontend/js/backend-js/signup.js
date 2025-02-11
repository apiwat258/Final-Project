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
        const response = await fetch('http://127.0.0.1:8080/api/register', { // 🔹 ส่งไปที่ Backend
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password }),
        });

        const data = await response.json();
        if (response.ok) {
            localStorage.setItem("user_id", data.user_id);  // 🔹 บันทึก User ID
            localStorage.setItem('user_email', email); // 🔹 เก็บอีเมลไว้ใช้ในหน้าถัดไป
            window.location.href = 'Role.html';
        } else {
            alert(data.error || 'Registration failed!');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred. Please try again later.');
    }
});
