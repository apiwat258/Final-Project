const RawMilk = artifacts.require("RawMilk");
const Product = artifacts.require("Product");
const ProductLot = artifacts.require("ProductLot");
const ShippingEvents = artifacts.require("ShippingEvents");

module.exports = async function (deployer, network, accounts) {
  console.log("Deploying contracts using account:", accounts[0]); // ✅ Log ตรวจสอบ
  await deployer.deploy(RawMilk, accounts[0]); // ใช้ Account แรกเป็น Owner
  await deployer.deploy(Product);
  await deployer.deploy(ProductLot);
  await deployer.deploy(ShippingEvents);

  // ✅ บันทึกที่อยู่สัญญาลงไฟล์ JSON
  const fs = require("fs");
  const deployedContracts = {
    RawMilk: RawMilk.address,
    Product: Product.address,
    ProductLot: ProductLot.address,
    ShippingEvents: ShippingEvents.address,
  };
  fs.writeFileSync("deployedContracts.json", JSON.stringify(deployedContracts, null, 2));

  console.log("✅ Contracts deployed and addresses saved to deployedContracts.json");
};
