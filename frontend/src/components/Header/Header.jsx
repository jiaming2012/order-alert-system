import React from "react";
import "./Header.scss";
import logo from "./logo.jpg";

const Header = () => (
    <div className="header">
        <img src={logo} alt="" id="my-logo" />
        <h2>Order Messenger</h2>
    </div>
);

export default Header;