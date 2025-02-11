// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Product is Ownable {
    constructor(address initialOwner) Ownable(initialOwner) {}

    struct ProductInfo {
        string productID;
        string factoryID;
        string productName;
        string nutritionalInfoCID;
        uint256 createdOn;
    }

    mapping(string => ProductInfo) public products;
    event ProductCreated(
        string productID,
        string factoryID,
        string productName,
        string nutritionalInfoCID,
        uint256 createdOn
    );

    function addProduct(
        string memory _productID,
        string memory _factoryID,
        string memory _productName,
        string memory _nutritionalInfoCID
    ) public onlyOwner {
        require(products[_productID].createdOn == 0, "Product already exists");

        products[_productID] = ProductInfo({
            productID: _productID,
            factoryID: _factoryID,
            productName: _productName,
            nutritionalInfoCID: _nutritionalInfoCID,
            createdOn: block.timestamp
        });

        emit ProductCreated(_productID, _factoryID, _productName, _nutritionalInfoCID, block.timestamp);
    }

    function getProduct(string memory _productID) public view returns (
        string memory, string memory, string memory, string memory, uint256
    ) {
        require(products[_productID].createdOn != 0, "Product does not exist");
        ProductInfo memory product = products[_productID];
        return (product.productID, product.factoryID, product.productName, product.nutritionalInfoCID, product.createdOn);
    }
}