document.addEventListener("DOMContentLoaded", function () {
    const emailInput = document.getElementById("email");
    if (emailInput) {
        const email = localStorage.getItem("user_email");
        if (email) {
            emailInput.value = email;
        } else {
            alert("User email not found. Please register first.");
            window.location.href = "SignUp.html";
        }
    }

    document.getElementById("factory-form").addEventListener("submit", async function (event) {
        event.preventDefault();

        const userID = localStorage.getItem("user_id");
        if (!userID) {
            alert("User ID not found. Please register first.");
            return;
        }

        const certificationCID = localStorage.getItem("certification_cid");

        const farmerData = {
            userid: userID,
            company_name: document.getElementById("company_name").value,
            firstname: document.getElementById("firstname").value,
            lastname: document.getElementById("lastname").value,
            email: document.getElementById("email").value,
            address: document.getElementById("address").value,
            phone: document.getElementById("phone").value,
            city: document.getElementById("city").value
        };

        try {
            const response = await fetch("http://127.0.0.1:8080/api/v1/factory", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(farmerData)
            });

            const textResponse = await response.text();
            console.log("Server Response:", textResponse);

            if (!response.ok) {
                alert("Error: " + textResponse);
                return;
            }

            const data = JSON.parse(textResponse);
            alert("Farmer information saved successfully!");

            // ✅ บันทึก Certification CID ถ้ามี
            if (certificationCID) {
                console.log("📌 Sending certification data to backend...");
                const certResponse = await fetch("http://127.0.0.1:8080/api/v1/createCertification", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({
                        farmerid: data.farmer_id,
                        certificationtype: "Organic ACT",
                        certificationcid: certificationCID,
                        issued_date: "2025-02-07"
                    })
                });

                const certData = await certResponse.json();
                console.log("✅ Certification Response:", certData);

                if (certResponse.ok) {
                    alert("Certification saved successfully!");
                } else {
                    alert("Failed to save certification: " + certData.error);
                }
            }

            window.location.href = "index.html";

        } catch (error) {
            console.error("Error:", error);
            alert("An error occurred while saving data.");
        }
    });
});
