import React, {Component} from "react";
import "./Orders.scss";
import Message from "../Message";

class Orders extends Component {
    render() {
        const orders = this.props.orders.map(msg =>
            <Message message={msg.data} />
        );

        return (
            <div className="Orders">
                <h2>Orders</h2>
                {orders}
            </div>
        )
    }
}

export default Orders;