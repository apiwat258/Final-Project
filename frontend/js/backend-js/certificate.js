document.getElementById("upload_certification").addEventListener("change", async function (event) {
    const file = event.target.files[0];
    if (!file) {
        alert("Please select a file.");
        return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
        console.log("📌 Uploading file to IPFS...");
        const response = await fetch("http://127.0.0.1:8080/api/v1/uploadCertificate", {
            method: "POST",
            body: formData
        });

        const result = await response.json();
        console.log("✅ IPFS Upload Result:", result);

        if (!response.ok || !result.cid) {
            alert("❌ Failed to upload file to IPFS");
            return;
        }

        const certificationCID = result.cid;
        console.log("✅ Certification CID:", certificationCID);

        // ✅ เก็บ `CID` ไว้ใน localStorage ใช้ร่วมกันทุกบทบาท
        localStorage.setItem("certification_cid", certificationCID);
        alert("File uploaded successfully! CID: " + certificationCID);

    } catch (error) {
        console.error("❌ Error uploading file:", error);
        alert("An error occurred while uploading.");
    }
});
