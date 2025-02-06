const API_BASE_URL = "http://localhost:8080/api/v1";

// ฟังก์ชันทดสอบว่า API Backend ทำงานได้ไหม
async function testAPI() {
    try {
        const response = await fetch(`${API_BASE_URL}/`);
        const data = await response.json();
        console.log("API Response:", data);
    } catch (error) {
        console.error("Error fetching API:", error);
    }
}

// ฟังก์ชันโหลดข้อมูลผู้ใช้ (ตัวอย่าง)
async function fetchUsers() {
    try {
        const response = await fetch(`${API_BASE_URL}/users`);
        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching users:", error);
    }
}

// เรียกฟังก์ชันทดสอบ API
testAPI();
